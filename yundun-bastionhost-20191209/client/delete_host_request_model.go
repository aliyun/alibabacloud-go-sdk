// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *DeleteHostRequest
	GetHostId() *string
	SetInstanceId(v string) *DeleteHostRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteHostRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteHostRequest
	GetRegionId() *string
}

type DeleteHostRequest struct {
	// The ID of the host to be deleted.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The instance ID of the bastion host to which the host to be deleted belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
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
	// The region ID of the bastion host instance to which the host to be deleted belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostRequest) GetHostId() *string {
	return s.HostId
}

func (s *DeleteHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHostRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostRequest) SetHostId(v string) *DeleteHostRequest {
	s.HostId = &v
	return s
}

func (s *DeleteHostRequest) SetInstanceId(v string) *DeleteHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostRequest) SetProjectId(v int64) *DeleteHostRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteHostRequest) SetRegionId(v string) *DeleteHostRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostRequest) Validate() error {
	return dara.Validate(s)
}
