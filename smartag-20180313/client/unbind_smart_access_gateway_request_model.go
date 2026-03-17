// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *UnbindSmartAccessGatewayRequest
	GetCcnId() *string
	SetOwnerAccount(v string) *UnbindSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnbindSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnbindSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UnbindSmartAccessGatewayRequest
	GetSmartAGId() *string
	SetSmartAGUid(v int64) *UnbindSmartAccessGatewayRequest
	GetSmartAGUid() *int64
}

type UnbindSmartAccessGatewayRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-kygbldwikz********
	CcnId        *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
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
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0ovhf732a********
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 1688401595963306
	SmartAGUid *int64 `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
}

func (s UnbindSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *UnbindSmartAccessGatewayRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *UnbindSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UnbindSmartAccessGatewayRequest) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *UnbindSmartAccessGatewayRequest) SetCcnId(v string) *UnbindSmartAccessGatewayRequest {
	s.CcnId = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetOwnerAccount(v string) *UnbindSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetOwnerId(v int64) *UnbindSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetRegionId(v string) *UnbindSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *UnbindSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *UnbindSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetSmartAGId(v string) *UnbindSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) SetSmartAGUid(v int64) *UnbindSmartAccessGatewayRequest {
	s.SmartAGUid = &v
	return s
}

func (s *UnbindSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
