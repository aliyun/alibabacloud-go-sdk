// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwareForUserDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListSoftwareForUserDeviceRequest
	GetCurrentPage() *int64
	SetDeviceTag(v string) *ListSoftwareForUserDeviceRequest
	GetDeviceTag() *string
	SetPageSize(v int64) *ListSoftwareForUserDeviceRequest
	GetPageSize() *int64
}

type ListSoftwareForUserDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSoftwareForUserDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwareForUserDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSoftwareForUserDeviceRequest) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *ListSoftwareForUserDeviceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSoftwareForUserDeviceRequest) SetCurrentPage(v int64) *ListSoftwareForUserDeviceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSoftwareForUserDeviceRequest) SetDeviceTag(v string) *ListSoftwareForUserDeviceRequest {
	s.DeviceTag = &v
	return s
}

func (s *ListSoftwareForUserDeviceRequest) SetPageSize(v int64) *ListSoftwareForUserDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *ListSoftwareForUserDeviceRequest) Validate() error {
	return dara.Validate(s)
}
