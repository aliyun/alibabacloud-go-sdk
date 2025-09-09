// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v string) *DetachHostGroupAccountsFromUserGroupRequest
	GetHostGroups() *string
	SetInstanceId(v string) *DetachHostGroupAccountsFromUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachHostGroupAccountsFromUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupRequest
	GetUserGroupId() *string
}

type DetachHostGroupAccountsFromUserGroupRequest struct {
	// The ID of the host group and the name of host account on which you want to revoke permissions from the user group. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the permissions on the specified host groups and all host accounts in the host groups are revoked from the user group. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// >  You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the name of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostGroupId":"1"}, {"HostGroupId":"2","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"3","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"4","HostAccountNames":["root","111","abc"]}]
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user group.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to revoke permissions on the specified host groups and host accounts.
	//
	// >  You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) GetHostGroups() *string {
	return s.HostGroups
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.HostGroups = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetInstanceId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetRegionId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
