// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingPublicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *QueryGateVerifyBillingPublicRequest
	GetAuthenticationType() *int32
	SetMonth(v string) *QueryGateVerifyBillingPublicRequest
	GetMonth() *string
	SetOwnerId(v int64) *QueryGateVerifyBillingPublicRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryGateVerifyBillingPublicRequest
	GetResourceOwnerAccount() *string
}

type QueryGateVerifyBillingPublicRequest struct {
	// The verification method. Valid values:
	//
	// 	- **0**: phone number verification
	//
	// 	- **1**: one-click logon
	//
	// 	- **2**: all
	//
	// 	- **3**: facial recognition
	//
	// 	- **4**: SMS verification
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The month in which the bill is generated. Specify this parameter in the YYYYMM format. Example: 202111.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202111
	Month                *string `json:"Month,omitempty" xml:"Month,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s QueryGateVerifyBillingPublicRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *QueryGateVerifyBillingPublicRequest) GetMonth() *string {
	return s.Month
}

func (s *QueryGateVerifyBillingPublicRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryGateVerifyBillingPublicRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryGateVerifyBillingPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyBillingPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetMonth(v string) *QueryGateVerifyBillingPublicRequest {
	s.Month = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetOwnerId(v int64) *QueryGateVerifyBillingPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyBillingPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) Validate() error {
	return dara.Validate(s)
}
