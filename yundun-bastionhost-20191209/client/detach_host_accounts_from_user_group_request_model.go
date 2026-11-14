// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v string) *DetachHostAccountsFromUserGroupRequest
	GetHosts() *string
	SetInstanceId(v string) *DetachHostAccountsFromUserGroupRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DetachHostAccountsFromUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *DetachHostAccountsFromUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *DetachHostAccountsFromUserGroupRequest
	GetUserGroupId() *string
}

type DetachHostAccountsFromUserGroupRequest struct {
	// Specifies the IDs of the hosts and host accounts whose authorization you want to remove from the user group.
	//
	// You can specify up to 10 host IDs, and up to 10 host account IDs for each host. You can leave the host account IDs empty. If you do not specify host account IDs, the host and all its authorized host accounts are removed from the user group. For more information about the structure of this parameter, see the Hosts parameter structure description in the request parameters section.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain host IDs and call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain host account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3",...]}, {"HostId":"3","HostAccountIds":["4","5","6"]}, {"HostId":"4","HostAccountIds":["9","8","7"]} ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// Specifies the instance ID of the Bastionhost instance that manages the user group for which you want to remove the authorization of hosts and host accounts.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies the region ID of the Bastionhost instance that manages the user group from which you want to remove authorized hosts and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the ID of the user group from which you want to remove authorized hosts and host accounts.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupRequest) GetHosts() *string {
	return s.Hosts
}

func (s *DetachHostAccountsFromUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostAccountsFromUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DetachHostAccountsFromUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostAccountsFromUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachHostAccountsFromUserGroupRequest) SetHosts(v string) *DetachHostAccountsFromUserGroupRequest {
	s.Hosts = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetInstanceId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetProjectId(v int64) *DetachHostAccountsFromUserGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetRegionId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetUserGroupId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
