// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v string) *AttachHostAccountsToUserRequest
	GetHosts() *string
	SetInstanceId(v string) *AttachHostAccountsToUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachHostAccountsToUserRequest
	GetRegionId() *string
	SetUserId(v string) *AttachHostAccountsToUserRequest
	GetUserId() *string
}

type AttachHostAccountsToUserRequest struct {
	// The IDs of the hosts and host accounts that you want to authorize the user to manage. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the user is authorized to manage only the specified hosts. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host and the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the ID of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3"]}, {"HostId":"3","HostAccountIds":["4","5","6"]}, {"HostId":"4","HostAccountIds":["9","8","7"]}  ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host for which you want to authorize the user to manage the hosts and host accounts.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user to manage the hosts and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to authorize to manage the hosts and host accounts.
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

func (s AttachHostAccountsToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserRequest) GetHosts() *string {
	return s.Hosts
}

func (s *AttachHostAccountsToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachHostAccountsToUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachHostAccountsToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AttachHostAccountsToUserRequest) SetHosts(v string) *AttachHostAccountsToUserRequest {
	s.Hosts = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetInstanceId(v string) *AttachHostAccountsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetRegionId(v string) *AttachHostAccountsToUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetUserId(v string) *AttachHostAccountsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) Validate() error {
	return dara.Validate(s)
}
