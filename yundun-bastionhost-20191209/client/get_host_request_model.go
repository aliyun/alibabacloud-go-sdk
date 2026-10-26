// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *GetHostRequest
	GetHostId() *string
	SetInstanceId(v string) *GetHostRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetHostRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetHostRequest
	GetRegionId() *string
}

type GetHostRequest struct {
	// The ID of the host to query. Only one host ID is supported.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The instance ID of the bastion host to which the host to query belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The region ID of the bastion host instance to which the host to query belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostRequest) GoString() string {
	return s.String()
}

func (s *GetHostRequest) GetHostId() *string {
	return s.HostId
}

func (s *GetHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostRequest) SetHostId(v string) *GetHostRequest {
	s.HostId = &v
	return s
}

func (s *GetHostRequest) SetInstanceId(v string) *GetHostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostRequest) SetProjectId(v int64) *GetHostRequest {
	s.ProjectId = &v
	return s
}

func (s *GetHostRequest) SetRegionId(v string) *GetHostRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostRequest) Validate() error {
	return dara.Validate(s)
}
