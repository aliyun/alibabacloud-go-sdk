// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v string) *AttachHostGroupAccountsToUserRequest
	GetHostGroups() *string
	SetInstanceId(v string) *AttachHostGroupAccountsToUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachHostGroupAccountsToUserRequest
	GetRegionId() *string
	SetUserId(v string) *AttachHostGroupAccountsToUserRequest
	GetUserId() *string
}

type AttachHostGroupAccountsToUserRequest struct {
	// The ID of the host group and the name of the host account that you want to authorize the user to manage. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the user is authorized to manage only the specified host groups. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the name of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostGroupId":"1"}, {"HostGroupId":"2","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"3","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"4","HostAccountNames":["root","111","abc"]} ]
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to authorize the user to manage the host groups and host accounts.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user to manage the host groups and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to authorize to manage the host groups and host accounts.
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

func (s AttachHostGroupAccountsToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserRequest) GetHostGroups() *string {
	return s.HostGroups
}

func (s *AttachHostGroupAccountsToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachHostGroupAccountsToUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachHostGroupAccountsToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AttachHostGroupAccountsToUserRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserRequest {
	s.HostGroups = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetInstanceId(v string) *AttachHostGroupAccountsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetRegionId(v string) *AttachHostGroupAccountsToUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetUserId(v string) *AttachHostGroupAccountsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) Validate() error {
	return dara.Validate(s)
}
