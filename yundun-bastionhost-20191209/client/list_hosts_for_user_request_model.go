// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAddress(v string) *ListHostsForUserRequest
	GetHostAddress() *string
	SetHostName(v string) *ListHostsForUserRequest
	GetHostName() *string
	SetInstanceId(v string) *ListHostsForUserRequest
	GetInstanceId() *string
	SetMode(v string) *ListHostsForUserRequest
	GetMode() *string
	SetOSType(v string) *ListHostsForUserRequest
	GetOSType() *string
	SetPageNumber(v string) *ListHostsForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostsForUserRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListHostsForUserRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostsForUserRequest
	GetUserId() *string
}

type ListHostsForUserRequest struct {
	// The address of the host to query. You can use a domain name or an IP address. Only exact match is supported.
	//
	// example:
	//
	// 192.168.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host to query. Only exact match is supported.
	//
	// example:
	//
	// abc
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the Bastionhost instance to which the user belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode of the host list to query. Valid values:
	//
	// - **Authorized**: Authorized hosts. This is the default value.
	//
	// - **Unauthorized**: Unauthorized hosts.
	//
	// example:
	//
	// Authorized
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The operating system type of the host to query. Valid values:
	//
	// - **Linux**
	//
	// - **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The page number of the current page in a paging query. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for a paging query.
	//
	// The maximum value of the PageSize parameter is 100. Default value: 20. If this parameter is left empty, 20 entries are returned by default.
	//
	// > Specify a value for the PageSize parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the user belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to query.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *ListHostsForUserRequest) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostsForUserRequest) GetMode() *string {
	return s.Mode
}

func (s *ListHostsForUserRequest) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostsForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostsForUserRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListHostsForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListHostsForUserRequest) SetHostAddress(v string) *ListHostsForUserRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserRequest) SetHostName(v string) *ListHostsForUserRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserRequest) SetInstanceId(v string) *ListHostsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserRequest) SetMode(v string) *ListHostsForUserRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserRequest) SetOSType(v string) *ListHostsForUserRequest {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserRequest) SetPageNumber(v string) *ListHostsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsForUserRequest) SetPageSize(v string) *ListHostsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsForUserRequest) SetProjectId(v int64) *ListHostsForUserRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHostsForUserRequest) SetRegionId(v string) *ListHostsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserRequest) SetUserId(v string) *ListHostsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostsForUserRequest) Validate() error {
	return dara.Validate(s)
}
