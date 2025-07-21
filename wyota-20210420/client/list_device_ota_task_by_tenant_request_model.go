// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceOtaTaskByTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDeviceOtaTaskByTenantRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeviceOtaTaskByTenantRequest
	GetPageSize() *int32
}

type ListDeviceOtaTaskByTenantRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceOtaTaskByTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceOtaTaskByTenantRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceOtaTaskByTenantRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeviceOtaTaskByTenantRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceOtaTaskByTenantRequest) SetPageNumber(v int32) *ListDeviceOtaTaskByTenantRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantRequest) SetPageSize(v int32) *ListDeviceOtaTaskByTenantRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceOtaTaskByTenantRequest) Validate() error {
	return dara.Validate(s)
}
