package url

import (
	"syscall/js"
)

// Use Struct so we cn use define functions
type jsGlobal struct {
	Doc js.Value
	// DivGridMap map[string]DivGrid // Dictionary of ivGrids so multiple Grids can be handled in one htmlpage
}

/*
// Store DivGrid
func (j jsGlobal) GetGrid(DivID string) DivGrid {
	return (j.DivGridMap[DivID])
}

// Store DivGrid
func (j *jsGlobal) SetGrid(DivID string, value DivGrid) {
	j.DivGridMap[DivID] = value
}

// Exist the Grid
func (j jsGlobal) ExistGrid(DivID string) bool {
	if _, ok := j.DivGridMap[DivID]; ok {
		return true
	}
	return false
}
*/
