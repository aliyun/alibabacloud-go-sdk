// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostIds(v string) *ModifyHostsProjectRequest
	GetHostIds() *string
	SetInstanceId(v string) *ModifyHostsProjectRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ModifyHostsProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyHostsProjectRequest
	GetRegionId() *string
	SetTargetProjectId(v int64) *ModifyHostsProjectRequest
	GetTargetProjectId() *int64
}

type ModifyHostsProjectRequest struct {
	// The IDs of the hosts that you want to move. This parameter is a JSON-formatted string. You can specify up to 100 host IDs. Separate multiple host IDs with commas (,).
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the project to which the hosts currently belong.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the target project to which you want to move the hosts.
	//
	// > You can call the ListProjectsForProjectAuthorizationPrincipal operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ModifyHostsProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsProjectRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *ModifyHostsProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostsProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ModifyHostsProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostsProjectRequest) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ModifyHostsProjectRequest) SetHostIds(v string) *ModifyHostsProjectRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsProjectRequest) SetInstanceId(v string) *ModifyHostsProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsProjectRequest) SetProjectId(v int64) *ModifyHostsProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyHostsProjectRequest) SetRegionId(v string) *ModifyHostsProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsProjectRequest) SetTargetProjectId(v int64) *ModifyHostsProjectRequest {
	s.TargetProjectId = &v
	return s
}

func (s *ModifyHostsProjectRequest) Validate() error {
	return dara.Validate(s)
}
