// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartListShrink(v string) *DepartmentSaveShrinkRequest
	GetDepartListShrink() *string
}

type DepartmentSaveShrinkRequest struct {
	DepartListShrink *string `json:"depart_list,omitempty" xml:"depart_list,omitempty"`
}

func (s DepartmentSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *DepartmentSaveShrinkRequest) GetDepartListShrink() *string {
	return s.DepartListShrink
}

func (s *DepartmentSaveShrinkRequest) SetDepartListShrink(v string) *DepartmentSaveShrinkRequest {
	s.DepartListShrink = &v
	return s
}

func (s *DepartmentSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
