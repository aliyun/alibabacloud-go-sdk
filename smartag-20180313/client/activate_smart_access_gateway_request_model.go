// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ActivateSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ActivateSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ActivateSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ActivateSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ActivateSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ActivateSmartAccessGatewayRequest
	GetSmartAGId() *string
}

type ActivateSmartAccessGatewayRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
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
	// sag-ke3kq4evpi8p6******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ActivateSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *ActivateSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ActivateSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ActivateSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ActivateSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ActivateSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ActivateSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ActivateSmartAccessGatewayRequest) SetOwnerAccount(v string) *ActivateSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) SetOwnerId(v int64) *ActivateSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) SetRegionId(v string) *ActivateSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *ActivateSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *ActivateSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) SetSmartAGId(v string) *ActivateSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *ActivateSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
