// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v int64) *DeleteDepartmentRequest
	GetDepartmentId() *int64
	SetInstanceId(v string) *DeleteDepartmentRequest
	GetInstanceId() *string
}

type DeleteDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *DeleteDepartmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDepartmentRequest) SetDepartmentId(v int64) *DeleteDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *DeleteDepartmentRequest) SetInstanceId(v string) *DeleteDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
