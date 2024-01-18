package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"
)

const DEBUG = false

func show_assignments(costs, stars [][]int, N, M int) int {
	total_cost := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if stars[i][j] == 1 {
				total_cost += costs[i][j]
				fmt.Printf("Agent %d -> Task %d : cost %d\n", i, j, costs[i][j])
			}
		}
	}
	fmt.Printf("\nTotal cost %d\n", total_cost)
	return total_cost
}

func debug_array(step int, costs, stars, primes [][]int, covered_agents, covered_tasks []int, assignments, N, M int) {
	// debug cost array
	fmt.Printf("STEP %d\n\n", step)
	fmt.Printf(" A \\ T ")
	for j := 0; j < M; j++ {
		fmt.Printf("|   %-2d  ", j)
	}
	fmt.Printf("\n")
	for i := 0; i < N; i++ {
		fmt.Printf("\n%-2d:\t", i)
		for j := 0; j < M; j++ {
			if stars[i][j] == 1 {
				fmt.Printf("   %-2d*  ", costs[i][j])
			} else if primes[i][j] == 1 {
				fmt.Printf("   %-2d'  ", costs[i][j])
			} else {
				fmt.Printf("   %-2d   ", costs[i][j])
			}
		}
		if covered_agents[i] == 1 {
			fmt.Printf("  x")
		} else {
			fmt.Printf("   ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\t")
	for j := 0; j < M; j++ {
		if covered_tasks[j] == 1 {
			fmt.Printf("   x    ")
		} else {
			fmt.Printf("        ")
		}
	}
	fmt.Printf("\n\n")
}

func clean(arr *[][]int, N, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			(*arr)[i][j] = 0
		}
	}
}

func clean_all(stars, primes *[][]int, covered_agents, covered_tasks *[]int, N, M int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			(*stars)[i][j] = 0
			(*primes)[i][j] = 0
			(*covered_tasks)[j] = 0
		}
		(*covered_agents)[i] = 0
	}
}

func task_in(arr [][]int, j, N int) int {
	for i := 0; i < N; i++ {
		if arr[i][j] == 1 {
			return i
		}
	}
	return -1
}

func agent_in(arr [][]int, i, N int) int {
	for j := 0; j < N; j++ {
		if arr[i][j] == 1 {
			return j
		}
	}
	return -1
}

func find_assignments(costs, stars *[][]int, N, M int) int {
	clean(stars, N, M)

	assignments := 0

	covered_agents := make([]int, N)
	covered_tasks := make([]int, M)
	for i := 0; i < N; i++ {
		covered_agents[i] = 0
	}
	for j := 0; j < M; j++ {
		covered_tasks[j] = 0
	}

	// repeat until every zero is covered
	for {
		min_nb_of_zeros := 10000000
		best_row := -1
		best_column := -1
		// find row or column with non null minimum number of zeros
		for i := 0; i < N; i++ {
			nb_of_zeros := 0
			for j := 0; j < M; j++ {
				if covered_agents[i] == 0 && covered_tasks[j] == 0 && (*costs)[i][j] == 0 {
					nb_of_zeros++
				}
			}
			if nb_of_zeros > 0 && nb_of_zeros < min_nb_of_zeros {
				min_nb_of_zeros = nb_of_zeros
				best_row = i
			}
		}
		for j := 0; j < M; j++ {
			nb_of_zeros := 0
			for i := 0; i < N; i++ {
				if covered_agents[i] == 0 && covered_tasks[j] == 0 && (*costs)[i][j] == 0 {
					nb_of_zeros++
				}
			}
			if nb_of_zeros > 0 && nb_of_zeros < min_nb_of_zeros {
				min_nb_of_zeros = nb_of_zeros
				best_column = j
			}
		}
		// assign first non covered zero from row or column with non null minimum number of zeros
		if best_column != -1 {
			for i := 0; i < N; i++ {
				if covered_agents[i] == 0 && covered_tasks[best_column] == 0 && (*costs)[i][best_column] == 0 {
					(*stars)[i][best_column] = 1
					covered_agents[i] = 1
					covered_tasks[best_column] = 1
					assignments++
					break
				}
			}
		} else if best_row != -1 {
			for j := 0; j < M; j++ {
				if covered_agents[best_row] == 0 && covered_tasks[j] == 0 && (*costs)[best_row][j] == 0 {
					(*stars)[best_row][j] = 1
					covered_agents[best_row] = 1
					covered_tasks[j] = 1
					assignments++
					break
				}
			}
		} else {
			break
		}
	}

	return assignments
}

func get_assignments(stars [][]int, N, M int) int {
	assignments := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if stars[i][j] == 1 {
				assignments++
			}
		}
	}
	return assignments
}

type HCell struct {
	x int
	y int
}

func teardown(search_start time.Time, step int, original_costs, stars, primes [][]int, covered_agents, covered_tasks []int, assignments, N, M int) int {
	elapsed := time.Since(search_start).Microseconds()
	total_cost := 0

	if N <= M && assignments == N {
		debug_array(step, original_costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
		total_cost = show_assignments(original_costs, stars, N, M)
	} else if M < N && assignments == M {
		debug_array(step, original_costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
		total_cost = show_assignments(original_costs, stars, N, M)
	}
	fmt.Printf("\nElapsed time: %d us\n", elapsed)
	return total_cost
}

func hungarian_method(costs [][]int, objective string) (int, error) {

	search_start := time.Now()

	N := len(costs)
	M := len(costs[0])

	step := 0

	// keep track of original values
	original_costs := make([][]int, N)
	for i := 0; i < N; i++ {
		original_costs[i] = make([]int, M)
		copy(original_costs[i], costs[i])
	}

	if objective == "maximise" {
		max_cost := 0
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if costs[i][j] > max_cost {
					max_cost = costs[i][j]
				}
			}
		}
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				costs[i][j] = max_cost - costs[i][j]
			}
		}
	}

	// utils arrays
	stars := make([][]int, N)
	primes := make([][]int, N)
	for i := 0; i < N; i++ {
		stars[i] = make([]int, M)
		primes[i] = make([]int, M)
		for j := 0; j < M; j++ {
			stars[i][j] = 0
			primes[i][j] = 0
		}
	}
	covered_agents := make([]int, N)
	for i := 0; i < N; i++ {
		covered_agents[i] = 0
	}
	covered_tasks := make([]int, M)
	for j := 0; j < M; j++ {
		covered_tasks[j] = 0
	}

	// debug cost array
	if DEBUG {
		debug_array(step, costs, stars, primes, covered_agents, covered_tasks, 0, N, M)
	}

	// STEP 1
	// reduce every row by its minimum
	if N <= M {
		for i := 0; i < N; i++ {
			min_cost := 10000000
			for j := 0; j < M; j++ {
				if costs[i][j] < min_cost {
					min_cost = costs[i][j]
				}
			}
			for j := 0; j < M; j++ {
				costs[i][j] -= min_cost
			}
		}
	}

	step++
	// find perfect assignments
	assignments := find_assignments(&costs, &stars, N, M)
	// debug cost array
	if DEBUG {
		debug_array(step, costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
	}
	if (N <= M && assignments == N) || (M < N && assignments == M) {
		return teardown(search_start, step, original_costs, stars, primes, covered_agents, covered_tasks, assignments, N, M), nil
	}
	clean(&stars, N, M)

	// STEP 2
	// reduce every column by its minimum
	if N >= M {
		for j := 0; j < M; j++ {
			min_cost := 10000000
			for i := 0; i < N; i++ {
				if costs[i][j] < min_cost {
					min_cost = costs[i][j]
				}
			}
			for i := 0; i < N; i++ {
				costs[i][j] -= min_cost
			}
		}
	}

	step++
	// find perfect assignments
	assignments = find_assignments(&costs, &stars, N, M)
	// debug cost array
	if DEBUG {
		debug_array(2, costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
	}

	if (N <= M && assignments != N) || (M < N && assignments != M) {
		// STEP 3
		// cover every assigned column
		for {
			repeat_previous_step := false
			for j := 0; j < M; j++ {
				if task_in(stars, j, N) != -1 {
					covered_tasks[j] = 1
				}
			}
			// while there is uncovered zeroes, for each one, prime it, if row contains assigned zero, uncover its column and cover the current row
			count := 0
			for count < N*M {
				count = 0
				for i := 0; i < N; i++ {
					for j := 0; j < M; j++ {
						if covered_tasks[j] == 0 && covered_agents[i] == 0 && costs[i][j] == 0 && primes[i][j] == 0 && stars[i][j] == 0 {
							primes[i][j] = 1
							task := agent_in(stars, i, N)
							if task != -1 {
								covered_tasks[task] = 0
								covered_agents[i] = 1
							} else {
								// if row does not contains assigned zero
								curr_node := HCell{x: j, y: i}
								path := []HCell{curr_node}
								for {
									// check if column has assigned agent, add to path
									agent := task_in(stars, curr_node.x, N)
									if agent != -1 {
										curr_node = HCell{x: curr_node.x, y: agent}
										path = append(path, curr_node)
										// check if row has primed zero, add to path
										task = agent_in(primes, agent, N)
										if task != -1 {
											curr_node = HCell{x: task, y: agent}
											path = append(path, curr_node)
										} else {
											err := fmt.Sprintf("ERROR: Prime not found in row: %d\n", agent)
											return 0, errors.New(err)
										}
									} else {
										break
									}
								}
								// reverse stars and primes on found path
								for _, p := range path {
									if primes[p.y][p.x] == 1 {
										primes[p.y][p.x] = 0
										stars[p.y][p.x] = 1
									} else if stars[p.y][p.x] == 1 {
										primes[p.y][p.x] = 1
										stars[p.y][p.x] = 0
									}
								}
								// clean all coverings and primes
								for i := 0; i < N; i++ {
									covered_agents[i] = 0
									covered_tasks[i] = 0
								}
								clean(&primes, N, M)
								repeat_previous_step = true
								break
							}
						} else {
							count++
						}
					}
					if repeat_previous_step {
						break
					}
				}
				if repeat_previous_step {
					break
				}
			}

			if repeat_previous_step {
				continue
			}

			step++
			assignments = get_assignments(stars, N, M)
			// debug cost array
			if DEBUG {
				debug_array(step, costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
			}

			if (N <= M && assignments == N) || (M < N && assignments == M) {
				break
			}

			// STEP 4
			// find minimum uncovered value
			min_uncovered := 10000000
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					if covered_agents[i] == 0 && covered_tasks[j] == 0 {
						if costs[i][j] < min_uncovered {
							min_uncovered = costs[i][j]
						}
					}
				}
			}
			// add minimum uncovered value to coverings interserctions and substract it from uncovered values
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					if covered_agents[i] == 1 && covered_tasks[j] == 1 {
						costs[i][j] += min_uncovered
					} else if covered_agents[i] == 0 && covered_tasks[j] == 0 {
						costs[i][j] -= min_uncovered
					}
				}
			}

			step++
			clean_all(&stars, &primes, &covered_agents, &covered_tasks, N, M)
			// search assignments
			assignments = find_assignments(&costs, &stars, N, M)
			if DEBUG {
				debug_array(step, costs, stars, primes, covered_agents, covered_tasks, assignments, N, M)
			}
			if (N <= M && assignments == N) || (M < N && assignments == M) {
				break
			} else {
				step -= 2
			}
		}
	}

	return teardown(search_start, step, original_costs, stars, primes, covered_agents, covered_tasks, assignments, N, M), nil
}

func main() {

	var objective = flag.String("objective", "minimise", "Algorithm objective: minimise (default) or maximise")
	flag.Parse()

	fmt.Println(*objective)

	// input costs array: Agents \ Tasks, must be squared
	/*costs := [][]int{
		{1, 3, 3, 6, 4, 99, 5, 9, 7},
		{2, 4, 4, 5, 7, 5, 6, 6, 8},
		{2, 4, 4, 5, 7, 5, 6, 6, 8},
		{3, 99, 7, 4, 99, 99, 99, 99, 99},
		{3, 99, 5, 10, 99, 99, 99, 99, 99},
		{4, 6, 6, 9, 9, 7, 8, 10, 10},
		{4, 6, 6, 9, 9, 7, 8, 10, 10},
		{5, 99, 7, 8, 99, 99, 99, 99, 99},
		{6, 99, 8, 7, 99, 99, 99, 99, 99},
	}*/

	/*costs := [][]int{
		{4, 6, 3, 8},
		{7, 5, 12, 6},
		{3, 6, 9, 2},
		{1000, 5, 7, 4},
	}*/

	costs := [][]int{
		{18, 11, 16, 20},
		{14, 19, 26, 18},
		{21, 23, 35, 29},
		{32, 27, 21, 17},
		{16, 15, 28, 25},
	}

	_, err := hungarian_method(costs, *objective)
	if err != nil {
		log.Fatal(err)
	}
}
