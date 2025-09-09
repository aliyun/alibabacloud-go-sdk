// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v string) *AttachHostGroupAccountsToUserGroupRequest
	GetHostGroups() *string
	SetInstanceId(v string) *AttachHostGroupAccountsToUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachHostGroupAccountsToUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupRequest
	GetUserGroupId() *string
}

type AttachHostGroupAccountsToUserGroupRequest struct {
	// The ID of the host group and the name of the host account that you want to authorize the user group to manage. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the user group is authorized to manage only the specified host groups. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the name of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostGroupId":"1"}, {"HostGroupId":"2","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"3","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"4","HostAccountNames":["root","111","abc"]}]
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to authorize the user group to manage the specified host groups and host accounts.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user group to manage the specified host groups and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to authorize to manage the specified host groups and host accounts.
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

func (s AttachHostGroupAccountsToUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupRequest) GetHostGroups() *string {
	return s.HostGroups
}

func (s *AttachHostGroupAccountsToUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachHostGroupAccountsToUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachHostGroupAccountsToUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.HostGroups = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetInstanceId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetRegionId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
