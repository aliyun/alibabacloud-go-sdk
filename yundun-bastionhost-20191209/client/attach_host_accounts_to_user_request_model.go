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
	// The IDs of the hosts and host accounts to be authorized for the user. You can specify up to 10 host IDs, and up to 9 host account IDs for each host. If you do not specify host account IDs, only the hosts are authorized for the user. For more information about the structure of this parameter, see the Hosts parameter structure description in the request parameters section.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query host IDs and call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query host account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3"]}, {"HostId":"3","HostAccountIds":["4","5","6"]}, {"HostId":"4","HostAccountIds":["9","8","7"]}  ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The instance ID of the Bastionhost instance to which the user to be granted authorization for hosts and host accounts belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance where the user to be authorized for hosts and host accounts resides.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to whom you want to authorize hosts and host accounts.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the value of this parameter.
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
