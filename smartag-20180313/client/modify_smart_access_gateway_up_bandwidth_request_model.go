// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayUpBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySmartAccessGatewayUpBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySmartAccessGatewayUpBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySmartAccessGatewayUpBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayUpBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySmartAccessGatewayUpBandwidthRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySmartAccessGatewayUpBandwidthRequest
	GetSmartAGId() *string
	SetUpBandwidth4G(v int32) *ModifySmartAccessGatewayUpBandwidthRequest
	GetUpBandwidth4G() *int32
	SetUpBandwidthWan(v int32) *ModifySmartAccessGatewayUpBandwidthRequest
	GetUpBandwidthWan() *int32
}

type ModifySmartAccessGatewayUpBandwidthRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-jsy******************
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The maximum upstream bandwidth of 4G network connections established by the SAG device. Unit: Mbit/s.
	//
	// example:
	//
	// 3
	UpBandwidth4G *int32 `json:"UpBandwidth4G,omitempty" xml:"UpBandwidth4G,omitempty"`
	// The maximum upstream bandwidth of network connections established on the WAN port of the SAG device. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	UpBandwidthWan *int32 `json:"UpBandwidthWan,omitempty" xml:"UpBandwidthWan,omitempty"`
}

func (s ModifySmartAccessGatewayUpBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayUpBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetUpBandwidth4G() *int32 {
	return s.UpBandwidth4G
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) GetUpBandwidthWan() *int32 {
	return s.UpBandwidthWan
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetOwnerAccount(v string) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetOwnerId(v int64) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetRegionId(v string) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetResourceOwnerId(v int64) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetSmartAGId(v string) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetUpBandwidth4G(v int32) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.UpBandwidth4G = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) SetUpBandwidthWan(v int32) *ModifySmartAccessGatewayUpBandwidthRequest {
	s.UpBandwidthWan = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
