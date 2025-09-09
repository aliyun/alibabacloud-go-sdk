// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v string) *AttachHostAccountsToUserGroupRequest
	GetHosts() *string
	SetInstanceId(v string) *AttachHostAccountsToUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachHostAccountsToUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AttachHostAccountsToUserGroupRequest
	GetUserGroupId() *string
}

type AttachHostAccountsToUserGroupRequest struct {
	// The IDs of the host and host account that you want to authorize the user group to manage. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the user group is authorized to manage only the specified hosts. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the ID of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3",...]}, {"HostId":"3","HostAccountIds":["4","5","6",...]}, {"HostId":"4","HostAccountIds":["9","8","7",...]} ... ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host in which you want to authorize the user group to manage the specified hosts and host accounts.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to authorize the user group to manage the specified hosts and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
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

func (s AttachHostAccountsToUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupRequest) GetHosts() *string {
	return s.Hosts
}

func (s *AttachHostAccountsToUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachHostAccountsToUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachHostAccountsToUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachHostAccountsToUserGroupRequest) SetHosts(v string) *AttachHostAccountsToUserGroupRequest {
	s.Hosts = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetInstanceId(v string) *AttachHostAccountsToUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetRegionId(v string) *AttachHostAccountsToUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetUserGroupId(v string) *AttachHostAccountsToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
