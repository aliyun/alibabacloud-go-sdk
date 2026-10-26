// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKeyId(v string) *GetHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *GetHostShareKeyRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetHostShareKeyRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetHostShareKeyRequest
	GetRegionId() *string
}

type GetHostShareKeyRequest struct {
	// The ID of the host shared key to query.
	//
	// > You can call [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10427
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to obtain this parameter.
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
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *GetHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostShareKeyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostShareKeyRequest) SetHostShareKeyId(v string) *GetHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostShareKeyRequest) SetInstanceId(v string) *GetHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostShareKeyRequest) SetProjectId(v int64) *GetHostShareKeyRequest {
	s.ProjectId = &v
	return s
}

func (s *GetHostShareKeyRequest) SetRegionId(v string) *GetHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
