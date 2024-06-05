package word_search

import "fmt"

func WordExists_2024_06_05(board [][]byte, word string) bool {
	num_rows := len(board)
	num_cols := len(board[0])
	seen := make(map[string]struct{})


	var dfs func(row, col, current_word_index int) bool
	dfs = func(row, col, current_word_index int) bool {
		seen_key := fmt.Sprintf("%d-%d", row, col)

		if current_word_index == len(word) {
			return true
		}
	
		if row < 0 || col < 0 || row >= num_rows || col >= num_cols {
			return false
		}

		_, found := seen[seen_key]

		if found {
			return false
		}

		if (board[row][col] != word[current_word_index]) {
			return false
		}

		// Add to seen
		seen[seen_key] = struct{}{}

        // Explore all possible directions: down, up, right, left
        up := dfs(row+1, col, current_word_index+1)
        down :=   dfs(row-1, col, current_word_index+1)
		left := dfs(row, col-1, current_word_index+1)
		right := dfs(row, col+1, current_word_index+1)

		// Remove from seen
		delete(seen, seen_key)

		return up || down || left || right
	}

	for row := 0; row < num_rows; row++ {
		for col := 0; col < num_cols; col++ {
			if dfs(row, col, 0) {
				return true
			}
		}
	}

	return false
}