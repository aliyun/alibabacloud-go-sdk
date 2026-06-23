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
	// The address of the host. You can set this parameter to a domain name or an IP address. Exact match is supported.
	//
	// example:
	//
	// 192.168.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host. Exact match is supported.
	//
	// example:
	//
	// abc
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host instance that contains the user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query for authorized or unauthorized hosts. Valid values:
	//
	// - **Authorized*	- (default)
	//
	// - **Unauthorized**
	//
	// example:
	//
	// Authorized
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The operating system of the host. Valid values:
	//
	// - **Linux**
	//
	// - **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.<br>Maximum value: 100. Default value: 20.<br>
	//
	// > We recommend that you specify this parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the bastion host instance is located.
	//
	// > For more information about regions and their corresponding IDs, see [regions and availability zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group whose hosts you want to list.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to obtain the user group ID.
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
