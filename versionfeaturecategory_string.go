// Code generated by "stringer -type=VersionFeatureCategory"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VersionFeatureBoolean-0]
	_ = x[VersionFeatureDescription-1]
}

const _VersionFeatureCategory_name = "VersionFeatureBooleanVersionFeatureDescription"

var _VersionFeatureCategory_index = [...]uint8{0, 21, 46}

func (i VersionFeatureCategory) String() string {
	if i < 0 || i >= VersionFeatureCategory(len(_VersionFeatureCategory_index)-1) {
		return "VersionFeatureCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VersionFeatureCategory_name[_VersionFeatureCategory_index[i]:_VersionFeatureCategory_index[i+1]]
}
