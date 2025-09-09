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
	SetRegionId(v string) *ListHostsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostsForUserRequest
	GetUserId() *string
}

type ListHostsForUserRequest struct {
	// The endpoint of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	//
	// example:
	//
	// 192.168.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host that you want to query. Only exact match is supported.
	//
	// example:
	//
	// abc
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host on which you want to query the hosts that the user is authorized or not authorized to manage.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the category of the hosts that you want to query. Valid values:
	//
	// 	- **Authorized**: queries the hosts that the user is authorized to manage. This is the default value.
	//
	// 	- **Unauthorized**: queries the hosts that the user is not authorized to manage.
	//
	// example:
	//
	// Authorized
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The operating system of the host that you want to query. Valid values:
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The number of the page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query the hosts that the user is authorized or not authorized to manage.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
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
