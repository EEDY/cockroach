// Code generated by "stringer -type=opType"; DO NOT EDIT.

package schemachange

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[addColumn-0]
	_ = x[addConstraint-1]
	_ = x[createIndex-2]
	_ = x[createSequence-3]
	_ = x[createTable-4]
	_ = x[createTableAs-5]
	_ = x[createView-6]
	_ = x[dropColumn-7]
	_ = x[dropColumnDefault-8]
	_ = x[dropColumnNotNull-9]
	_ = x[dropColumnStored-10]
	_ = x[dropConstraint-11]
	_ = x[dropIndex-12]
	_ = x[dropSequence-13]
	_ = x[dropTable-14]
	_ = x[dropView-15]
	_ = x[renameColumn-16]
	_ = x[renameIndex-17]
	_ = x[renameSequence-18]
	_ = x[renameTable-19]
	_ = x[renameView-20]
	_ = x[setColumnDefault-21]
	_ = x[setColumnNotNull-22]
	_ = x[setColumnType-23]
}

const _opType_name = "addColumnaddConstraintcreateIndexcreateSequencecreateTablecreateTableAscreateViewdropColumndropColumnDefaultdropColumnNotNulldropColumnStoreddropConstraintdropIndexdropSequencedropTabledropViewrenameColumnrenameIndexrenameSequencerenameTablerenameViewsetColumnDefaultsetColumnNotNullsetColumnType"

var _opType_index = [...]uint16{0, 9, 22, 33, 47, 58, 71, 81, 91, 108, 125, 141, 155, 164, 176, 185, 193, 205, 216, 230, 241, 251, 267, 283, 296}

func (i opType) String() string {
	if i < 0 || i >= opType(len(_opType_index)-1) {
		return "opType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _opType_name[_opType_index[i]:_opType_index[i+1]]
}
