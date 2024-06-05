package word_search

var TestCases = []struct {
    Board    [][]byte
    Word     string
    Expected bool
}{
    {Board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, Word: "ABCCED", Expected: true},
    {Board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, Word: "ZEEP", Expected: false},
}



