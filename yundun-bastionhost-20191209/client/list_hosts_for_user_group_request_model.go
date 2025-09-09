// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAddress(v string) *ListHostsForUserGroupRequest
	GetHostAddress() *string
	SetHostName(v string) *ListHostsForUserGroupRequest
	GetHostName() *string
	SetInstanceId(v string) *ListHostsForUserGroupRequest
	GetInstanceId() *string
	SetMode(v string) *ListHostsForUserGroupRequest
	GetMode() *string
	SetOSType(v string) *ListHostsForUserGroupRequest
	GetOSType() *string
	SetPageNumber(v string) *ListHostsForUserGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostsForUserGroupRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostsForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListHostsForUserGroupRequest
	GetUserGroupId() *string
}

type ListHostsForUserGroupRequest struct {
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
	// The ID of the bastion host on which you want to query the hosts that the user group is authorized or not authorized to manage.
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
	// 	- **Authorized**: queries the hosts that the user group is authorized to manage. This is the default value.
	//
	// 	- **Unauthorized**: queries the hosts that the user group is not authorized to manage.
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
	// The region ID of the bastion host on which you want to query the hosts that the user group is authorized or not authorized to manage.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group for which you want to query hosts.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostsForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *ListHostsForUserGroupRequest) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostsForUserGroupRequest) GetMode() *string {
	return s.Mode
}

func (s *ListHostsForUserGroupRequest) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsForUserGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostsForUserGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostsForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostsForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListHostsForUserGroupRequest) SetHostAddress(v string) *ListHostsForUserGroupRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetHostName(v string) *ListHostsForUserGroupRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetInstanceId(v string) *ListHostsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetMode(v string) *ListHostsForUserGroupRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetOSType(v string) *ListHostsForUserGroupRequest {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetPageNumber(v string) *ListHostsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetPageSize(v string) *ListHostsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetRegionId(v string) *ListHostsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetUserGroupId(v string) *ListHostsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
