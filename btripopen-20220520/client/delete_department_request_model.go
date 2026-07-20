// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutDeptId(v string) *DeleteDepartmentRequest
	GetOutDeptId() *string
}

type DeleteDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dept123
	OutDeptId *string `json:"out_dept_id,omitempty" xml:"out_dept_id,omitempty"`
}

func (s DeleteDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *DeleteDepartmentRequest) SetOutDeptId(v string) *DeleteDepartmentRequest {
	s.OutDeptId = &v
	return s
}

func (s *DeleteDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
