// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v int64) *UpdateDepartmentRequest
	GetDepartmentId() *int64
	SetDepartmentName(v string) *UpdateDepartmentRequest
	GetDepartmentName() *string
	SetInstanceId(v string) *UpdateDepartmentRequest
	GetInstanceId() *string
}

type UpdateDepartmentRequest struct {
	// The department ID.
	//
	// You can invoke the [GetAllDepartment](https://help.aliyun.com/document_detail/2717975.html) API and view the **DepartmentId*	- field in the response to obtain the department ID.
	//
	// > This parameter does not support updates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// The department name.
	//
	// > This parameter supports updates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 部门A
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *UpdateDepartmentRequest) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *UpdateDepartmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDepartmentRequest) SetDepartmentId(v int64) *UpdateDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateDepartmentRequest) SetDepartmentName(v string) *UpdateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateDepartmentRequest) SetInstanceId(v string) *UpdateDepartmentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
