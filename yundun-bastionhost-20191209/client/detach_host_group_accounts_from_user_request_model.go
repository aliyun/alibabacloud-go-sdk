// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v string) *DetachHostGroupAccountsFromUserRequest
	GetHostGroups() *string
	SetInstanceId(v string) *DetachHostGroupAccountsFromUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachHostGroupAccountsFromUserRequest
	GetRegionId() *string
	SetUserId(v string) *DetachHostGroupAccountsFromUserRequest
	GetUserId() *string
}

type DetachHostGroupAccountsFromUserRequest struct {
	// The ID of the host group and the name of the host account on which you want to revoke permissions from the user. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the permissions on the specified host groups and all host accounts in the host groups are revoked from the user. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the name of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostGroupId":"1"}, {"HostGroupId":"2","HostAccountNames":["root","111","abc"]}]
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user from which you want to revoke permissions on the specified host groups and host accounts.
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

func (s DetachHostGroupAccountsFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserRequest) GetHostGroups() *string {
	return s.HostGroups
}

func (s *DetachHostGroupAccountsFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostGroupAccountsFromUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostGroupAccountsFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DetachHostGroupAccountsFromUserRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserRequest {
	s.HostGroups = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetInstanceId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetRegionId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetUserId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) Validate() error {
	return dara.Validate(s)
}
