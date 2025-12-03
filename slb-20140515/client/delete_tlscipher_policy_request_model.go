// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTLSCipherPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteTLSCipherPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTLSCipherPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTLSCipherPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTLSCipherPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTLSCipherPolicyRequest
	GetResourceOwnerId() *int64
	SetTLSCipherPolicyId(v string) *DeleteTLSCipherPolicyRequest
	GetTLSCipherPolicyId() *string
}

type DeleteTLSCipherPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the TLS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tls-bp1lp2076qx4ebridp******
	TLSCipherPolicyId *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s DeleteTLSCipherPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTLSCipherPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTLSCipherPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTLSCipherPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTLSCipherPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTLSCipherPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTLSCipherPolicyRequest) GetTLSCipherPolicyId() *string {
	return s.TLSCipherPolicyId
}

func (s *DeleteTLSCipherPolicyRequest) SetOwnerAccount(v string) *DeleteTLSCipherPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetOwnerId(v int64) *DeleteTLSCipherPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetRegionId(v string) *DeleteTLSCipherPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetResourceOwnerAccount(v string) *DeleteTLSCipherPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetResourceOwnerId(v int64) *DeleteTLSCipherPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetTLSCipherPolicyId(v string) *DeleteTLSCipherPolicyRequest {
	s.TLSCipherPolicyId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) Validate() error {
	return dara.Validate(s)
}
