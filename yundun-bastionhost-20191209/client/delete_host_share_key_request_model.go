// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKeyId(v string) *DeleteHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *DeleteHostShareKeyRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteHostShareKeyRequest
	GetRegionId() *string
}

type DeleteHostShareKeyRequest struct {
	// The shared key ID.
	//
	// >  You must specify this parameter. You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// example:
	//
	// 11206
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

func (s DeleteHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *DeleteHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostShareKeyRequest) SetHostShareKeyId(v string) *DeleteHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *DeleteHostShareKeyRequest) SetInstanceId(v string) *DeleteHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostShareKeyRequest) SetRegionId(v string) *DeleteHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
