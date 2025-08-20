// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListInstanceRequest
	GetInstanceName() *string
	SetInstanceStatus(v string) *ListInstanceRequest
	GetInstanceStatus() *string
	SetPageNo(v int32) *ListInstanceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListInstanceRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstanceRequest
	GetResourceGroupId() *string
}

type ListInstanceRequest struct {
	// Deprecated
	//
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- `PENDING`: The instance is being initialized.
	//
	// 	- `INIT_ERROR`: The initialization of the instance fails.
	//
	// 	- `STARTING`: The instance is being started.
	//
	// 	- `RUNNING`: The instance is running.
	//
	// 	- `STOPPING`: The instance is being stopped.
	//
	// 	- `STOPPED`: The instance is stopped.
	//
	// 	- `DELETING`: The instance is being deleted.
	//
	// 	- `DELETED`: The instance is deleted.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListInstanceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceRequest) SetInstanceName(v string) *ListInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceStatus(v string) *ListInstanceRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstanceRequest) SetPageNo(v int32) *ListInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int32) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetResourceGroupId(v string) *ListInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
