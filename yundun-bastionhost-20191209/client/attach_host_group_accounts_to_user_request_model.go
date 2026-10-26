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
	// The IDs of the host groups and the names of the host accounts to which you want to grant the user access. You can specify up to 10 host group IDs and up to 10 host account names for each host group. If you do not specify host account names, the user is granted access only to the host groups. For more information about the structure of this parameter, see the HostGroups parameter structure description in the request parameter list.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to obtain host group IDs and call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain host account names.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostGroupId":"1"}, {"HostGroupId":"2","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"3","HostAccountNames":["root","111","abc"]}, {"HostGroupId":"4","HostAccountNames":["root","111","abc"]} ]
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The instance ID of the Bastionhost instance to which the user to be granted authorization on host groups and host accounts belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the region ID of the Bastionhost instance where the user to be granted permissions on host groups and host accounts is located.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be granted permissions on host groups and host accounts.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to obtain this parameter.
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
