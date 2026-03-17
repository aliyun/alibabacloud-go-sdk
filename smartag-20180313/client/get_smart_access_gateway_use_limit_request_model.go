// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAccessGatewayUseLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetSmartAccessGatewayUseLimitRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetSmartAccessGatewayUseLimitRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetSmartAccessGatewayUseLimitRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetSmartAccessGatewayUseLimitRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmartAccessGatewayUseLimitRequest
	GetResourceOwnerId() *int64
}

type GetSmartAccessGatewayUseLimitRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetSmartAccessGatewayUseLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAccessGatewayUseLimitRequest) GoString() string {
	return s.String()
}

func (s *GetSmartAccessGatewayUseLimitRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetSmartAccessGatewayUseLimitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmartAccessGatewayUseLimitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSmartAccessGatewayUseLimitRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmartAccessGatewayUseLimitRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmartAccessGatewayUseLimitRequest) SetOwnerAccount(v string) *GetSmartAccessGatewayUseLimitRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitRequest) SetOwnerId(v int64) *GetSmartAccessGatewayUseLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitRequest) SetRegionId(v string) *GetSmartAccessGatewayUseLimitRequest {
	s.RegionId = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitRequest) SetResourceOwnerAccount(v string) *GetSmartAccessGatewayUseLimitRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitRequest) SetResourceOwnerId(v int64) *GetSmartAccessGatewayUseLimitRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitRequest) Validate() error {
	return dara.Validate(s)
}
