// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDeviceGroupsRequest
	GetCurrentPage() *int32
	SetDeviceGroupIds(v []*string) *ListDeviceGroupsRequest
	GetDeviceGroupIds() []*string
	SetName(v string) *ListDeviceGroupsRequest
	GetName() *string
	SetPageSize(v int32) *ListDeviceGroupsRequest
	GetPageSize() *int32
}

type ListDeviceGroupsRequest struct {
	// The page number of the current page displayed during paging in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IDs of the device groups. Duplicate IDs are not allowed.
	DeviceGroupIds []*string `json:"DeviceGroupIds,omitempty" xml:"DeviceGroupIds,omitempty" type:"Repeated"`
	// The name of the device tag. The name can be up to 128 characters in length and supports Chinese characters and uppercase and lowercase English letters. It can also contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Public service system user group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Settings for the number of entries per page during paging in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeviceGroupsRequest) GetDeviceGroupIds() []*string {
	return s.DeviceGroupIds
}

func (s *ListDeviceGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListDeviceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceGroupsRequest) SetCurrentPage(v int32) *ListDeviceGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetDeviceGroupIds(v []*string) *ListDeviceGroupsRequest {
	s.DeviceGroupIds = v
	return s
}

func (s *ListDeviceGroupsRequest) SetName(v string) *ListDeviceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListDeviceGroupsRequest) SetPageSize(v int32) *ListDeviceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
