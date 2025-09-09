// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupName(v string) *ListHostGroupsForUserRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *ListHostGroupsForUserRequest
	GetInstanceId() *string
	SetMode(v string) *ListHostGroupsForUserRequest
	GetMode() *string
	SetPageNumber(v string) *ListHostGroupsForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostGroupsForUserRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostGroupsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostGroupsForUserRequest
	GetUserId() *string
}

type ListHostGroupsForUserRequest struct {
	// The name of the host group to query. The name can be up to 128 characters in length. Only exact match is supported.
	//
	// example:
	//
	// group
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host whose user you want to query.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The category of the host groups to query. Valid values:
	//
	// 	- **Authorized*	- (default): queries the host groups that the user is authorized to manage.
	//
	// 	- **Unauthorized**: queries the host groups that the user is not authorized to manage.
	//
	// example:
	//
	// Authorized
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// １
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host whose user you want to query.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostGroupsForUserRequest) GetMode() *string {
	return s.Mode
}

func (s *ListHostGroupsForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostGroupsForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostGroupsForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostGroupsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListHostGroupsForUserRequest) SetHostGroupName(v string) *ListHostGroupsForUserRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetInstanceId(v string) *ListHostGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetMode(v string) *ListHostGroupsForUserRequest {
	s.Mode = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetPageNumber(v string) *ListHostGroupsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetPageSize(v string) *ListHostGroupsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetRegionId(v string) *ListHostGroupsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetUserId(v string) *ListHostGroupsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
