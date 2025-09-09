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
	SetRegionId(v string) *GetHostShareKeyRequest
	GetRegionId() *string
}

type GetHostShareKeyRequest struct {
	// The ID of the shared key whose information you want to query.
	//
	// >  You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10427
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetHostShareKeyRequest) SetRegionId(v string) *GetHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
