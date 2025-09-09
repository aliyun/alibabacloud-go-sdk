// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v string) *DetachHostAccountsFromUserRequest
	GetHosts() *string
	SetInstanceId(v string) *DetachHostAccountsFromUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachHostAccountsFromUserRequest
	GetRegionId() *string
	SetUserId(v string) *DetachHostAccountsFromUserRequest
	GetUserId() *string
}

type DetachHostAccountsFromUserRequest struct {
	// The IDs of the hosts and host accounts on which you want to revoke permissions from the user. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the permissions on the specified hosts and all accounts of the hosts are revoked from the user. For more information about this parameter, see the Description of the Hosts parameter section of this topic.
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the host IDs and the [ListHostAccountsForUser](https://help.aliyun.com/document_detail/466581.html) operation to query the host account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3"]}, {"HostId":"3","HostAccountIds":["4","5","6"]}, {"HostId":"4","HostAccountIds":["9","8","7"]} ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host on which you want to revoke permissions on the specified hosts and host accounts from the user.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to revoke permissions on the specified hosts and host accounts from the user.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user from whom you want to revoke permissions on the specified hosts and host accounts.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostAccountsFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserRequest) GetHosts() *string {
	return s.Hosts
}

func (s *DetachHostAccountsFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostAccountsFromUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostAccountsFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DetachHostAccountsFromUserRequest) SetHosts(v string) *DetachHostAccountsFromUserRequest {
	s.Hosts = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetInstanceId(v string) *DetachHostAccountsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetRegionId(v string) *DetachHostAccountsFromUserRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetUserId(v string) *DetachHostAccountsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) Validate() error {
	return dara.Validate(s)
}
