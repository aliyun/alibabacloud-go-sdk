// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *BindSmartAccessGatewayRequest
	GetCcnId() *string
	SetOwnerAccount(v string) *BindSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *BindSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *BindSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *BindSmartAccessGatewayRequest
	GetSmartAGId() *string
	SetSmartAGUid(v int64) *BindSmartAccessGatewayRequest
	GetSmartAGUid() *int64
}

type BindSmartAccessGatewayRequest struct {
	// The ID of the CCN instance with which you want to associate the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-isdjvvkexkrpk*****
	CcnId        *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
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
	// sag-m7ez44zpayma*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 1250123456123456
	SmartAGUid *int64 `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
}

func (s BindSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *BindSmartAccessGatewayRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *BindSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *BindSmartAccessGatewayRequest) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *BindSmartAccessGatewayRequest) SetCcnId(v string) *BindSmartAccessGatewayRequest {
	s.CcnId = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetOwnerAccount(v string) *BindSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetOwnerId(v int64) *BindSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetRegionId(v string) *BindSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *BindSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *BindSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetSmartAGId(v string) *BindSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) SetSmartAGUid(v int64) *BindSmartAccessGatewayRequest {
	s.SmartAGUid = &v
	return s
}

func (s *BindSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
