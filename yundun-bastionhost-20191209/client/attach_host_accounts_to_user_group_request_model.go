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
	SetProjectId(v int64) *AttachHostAccountsToUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *AttachHostAccountsToUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AttachHostAccountsToUserGroupRequest
	GetUserGroupId() *string
}

type AttachHostAccountsToUserGroupRequest struct {
	// The host IDs and host account IDs that you want to authorize the user group to manage. You can specify up to 10 host IDs, and each host supports up to 10 host account IDs. You do not need to specify host account IDs. If you do not specify host account IDs, only the hosts are authorized for the user group. For the specific structure of this parameter, see the Hosts parameter structure description under the request parameters list.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query host IDs and call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query host account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"HostId":"1"}, {"HostId":"2","HostAccountIds":["1","2","3",...]}, {"HostId":"3","HostAccountIds":["4","5","6",...]}, {"HostId":"4","HostAccountIds":["9","8","7",...]} ... ]
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The instance ID of the bastion host where the user group that you want to grant authorization resides.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
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
	// The region ID of the bastion host where the user group that you want to authorize resides.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to authorize to manage hosts and host accounts.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query this parameter.
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

func (s *AttachHostAccountsToUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *AttachHostAccountsToUserGroupRequest) SetProjectId(v int64) *AttachHostAccountsToUserGroupRequest {
	s.ProjectId = &v
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
