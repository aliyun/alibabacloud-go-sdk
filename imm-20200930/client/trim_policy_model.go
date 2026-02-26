// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrimPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetDisableDeleteEmptyCell(v bool) *TrimPolicy
	GetDisableDeleteEmptyCell() *bool
	SetDisableDeleteRepeatedStyle(v bool) *TrimPolicy
	GetDisableDeleteRepeatedStyle() *bool
	SetDisableDeleteUnusedPicture(v bool) *TrimPolicy
	GetDisableDeleteUnusedPicture() *bool
	SetDisableDeleteUnusedShape(v bool) *TrimPolicy
	GetDisableDeleteUnusedShape() *bool
}

type TrimPolicy struct {
	// Specifies whether to prevent all empty cells from being deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisableDeleteEmptyCell *bool `json:"DisableDeleteEmptyCell,omitempty" xml:"DisableDeleteEmptyCell,omitempty"`
	// Specifies whether to prevent all duplicate styles from being deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisableDeleteRepeatedStyle *bool `json:"DisableDeleteRepeatedStyle,omitempty" xml:"DisableDeleteRepeatedStyle,omitempty"`
	// Specifies whether to prevent unused cell images from being deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisableDeleteUnusedPicture *bool `json:"DisableDeleteUnusedPicture,omitempty" xml:"DisableDeleteUnusedPicture,omitempty"`
	// Specifies whether to prevent unused shapes from being deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisableDeleteUnusedShape *bool `json:"DisableDeleteUnusedShape,omitempty" xml:"DisableDeleteUnusedShape,omitempty"`
}

func (s TrimPolicy) String() string {
	return dara.Prettify(s)
}

func (s TrimPolicy) GoString() string {
	return s.String()
}

func (s *TrimPolicy) GetDisableDeleteEmptyCell() *bool {
	return s.DisableDeleteEmptyCell
}

func (s *TrimPolicy) GetDisableDeleteRepeatedStyle() *bool {
	return s.DisableDeleteRepeatedStyle
}

func (s *TrimPolicy) GetDisableDeleteUnusedPicture() *bool {
	return s.DisableDeleteUnusedPicture
}

func (s *TrimPolicy) GetDisableDeleteUnusedShape() *bool {
	return s.DisableDeleteUnusedShape
}

func (s *TrimPolicy) SetDisableDeleteEmptyCell(v bool) *TrimPolicy {
	s.DisableDeleteEmptyCell = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteRepeatedStyle(v bool) *TrimPolicy {
	s.DisableDeleteRepeatedStyle = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteUnusedPicture(v bool) *TrimPolicy {
	s.DisableDeleteUnusedPicture = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteUnusedShape(v bool) *TrimPolicy {
	s.DisableDeleteUnusedShape = &v
	return s
}

func (s *TrimPolicy) Validate() error {
	return dara.Validate(s)
}
