// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthWeighting(v string) *ModifyInstanceNetworkOptionsRequest
	GetBandwidthWeighting() *string
	SetInstanceId(v string) *ModifyInstanceNetworkOptionsRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *ModifyInstanceNetworkOptionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceNetworkOptionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceNetworkOptionsRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceNetworkOptionsRequest struct {
	// example:
	//
	// Vpc-L1
	BandwidthWeighting *string `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceNetworkOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkOptionsRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkOptionsRequest) GetBandwidthWeighting() *string {
	return s.BandwidthWeighting
}

func (s *ModifyInstanceNetworkOptionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetworkOptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceNetworkOptionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceNetworkOptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceNetworkOptionsRequest) SetBandwidthWeighting(v string) *ModifyInstanceNetworkOptionsRequest {
	s.BandwidthWeighting = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsRequest) SetInstanceId(v string) *ModifyInstanceNetworkOptionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsRequest) SetOwnerId(v int64) *ModifyInstanceNetworkOptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsRequest) SetResourceOwnerAccount(v string) *ModifyInstanceNetworkOptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsRequest) SetResourceOwnerId(v int64) *ModifyInstanceNetworkOptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsRequest) Validate() error {
	return dara.Validate(s)
}
