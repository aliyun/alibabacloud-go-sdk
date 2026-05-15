// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentName(v string) *CreateDepartmentRequest
	GetDepartmentName() *string
	SetInstanceId(v string) *CreateDepartmentRequest
	GetInstanceId() *string
}

type CreateDepartmentRequest struct {
	// This parameter is required.
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *CreateDepartmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDepartmentRequest) SetDepartmentName(v string) *CreateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *CreateDepartmentRequest) SetInstanceId(v string) *CreateDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
