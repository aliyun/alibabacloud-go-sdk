// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDeviceOtaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTenantDeviceOtaInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantDeviceOtaInfoRequest
	GetPageSize() *int32
	SetTaskId(v int32) *ListTenantDeviceOtaInfoRequest
	GetTaskId() *int32
}

type ListTenantDeviceOtaInfoRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTenantDeviceOtaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantDeviceOtaInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantDeviceOtaInfoRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ListTenantDeviceOtaInfoRequest) SetPageNumber(v int32) *ListTenantDeviceOtaInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTenantDeviceOtaInfoRequest) SetPageSize(v int32) *ListTenantDeviceOtaInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantDeviceOtaInfoRequest) SetTaskId(v int32) *ListTenantDeviceOtaInfoRequest {
	s.TaskId = &v
	return s
}

func (s *ListTenantDeviceOtaInfoRequest) Validate() error {
	return dara.Validate(s)
}
