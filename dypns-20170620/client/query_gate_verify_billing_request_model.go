// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *QueryGateVerifyBillingRequest
	GetAuthenticationType() *int32
	SetMonth(v string) *QueryGateVerifyBillingRequest
	GetMonth() *string
	SetOwnerId(v int64) *QueryGateVerifyBillingRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryGateVerifyBillingRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryGateVerifyBillingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryGateVerifyBillingRequest
	GetResourceOwnerId() *int64
}

type QueryGateVerifyBillingRequest struct {
	// This parameter is required.
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// This parameter is required.
	Month                *string `json:"Month,omitempty" xml:"Month,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryGateVerifyBillingRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *QueryGateVerifyBillingRequest) GetMonth() *string {
	return s.Month
}

func (s *QueryGateVerifyBillingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryGateVerifyBillingRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryGateVerifyBillingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryGateVerifyBillingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryGateVerifyBillingRequest) SetAuthenticationType(v int32) *QueryGateVerifyBillingRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) SetMonth(v string) *QueryGateVerifyBillingRequest {
	s.Month = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) SetOwnerId(v int64) *QueryGateVerifyBillingRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) SetProdCode(v string) *QueryGateVerifyBillingRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyBillingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) SetResourceOwnerId(v int64) *QueryGateVerifyBillingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryGateVerifyBillingRequest) Validate() error {
	return dara.Validate(s)
}
