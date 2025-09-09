// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountIds(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetHostAccountIds() *string
	SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachHostAccountsFromHostShareKeyRequest
	GetRegionId() *string
}

type DetachHostAccountsFromHostShareKeyRequest struct {
	// The host account IDs.
	//
	// >  You can call the [ListHostAccountsForHostShareKey](https://help.aliyun.com/document_detail/462975.html) operation to query the host account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The shared key ID.
	//
	// >  You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
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

func (s DetachHostAccountsFromHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetHostAccountIds() *string {
	return s.HostAccountIds
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostAccountIds(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostAccountIds = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetInstanceId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetRegionId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
