// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *UnlockSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnlockSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnlockSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnlockSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnlockSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UnlockSmartAccessGatewayRequest
	GetSmartAGId() *string
}

type UnlockSmartAccessGatewayRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-w9unmktmupcde*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UnlockSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *UnlockSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnlockSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnlockSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnlockSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnlockSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UnlockSmartAccessGatewayRequest) SetOwnerAccount(v string) *UnlockSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) SetOwnerId(v int64) *UnlockSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) SetRegionId(v string) *UnlockSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *UnlockSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *UnlockSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) SetSmartAGId(v string) *UnlockSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *UnlockSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
