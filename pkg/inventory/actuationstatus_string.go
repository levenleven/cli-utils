// Code generated by "stringer -type=ActuationStatus -linecomment"; DO NOT EDIT.

package inventory

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ActuationPending-0]
	_ = x[ActuationSucceeded-1]
	_ = x[ActuationSkipped-2]
	_ = x[ActuationFailed-3]
}

const _ActuationStatus_name = "PendingSucceededSkippedFailed"

var _ActuationStatus_index = [...]uint8{0, 7, 16, 23, 29}

func (i ActuationStatus) String() string {
	if i < 0 || i >= ActuationStatus(len(_ActuationStatus_index)-1) {
		return "ActuationStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ActuationStatus_name[_ActuationStatus_index[i]:_ActuationStatus_index[i+1]]
}
