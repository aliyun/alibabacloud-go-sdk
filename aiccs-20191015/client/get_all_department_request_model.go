// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAllDepartmentRequest
	GetInstanceId() *string
}

type GetAllDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAllDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAllDepartmentRequest) SetInstanceId(v string) *GetAllDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAllDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
