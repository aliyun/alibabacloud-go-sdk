// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneWithTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetPhoneWithTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetPhoneWithTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhoneWithTokenRequest
	GetResourceOwnerId() *int64
	SetSpToken(v string) *GetPhoneWithTokenRequest
	GetSpToken() *string
}

type GetPhoneWithTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token for phone number verification that is obtained by the JavaScript SDK. The validity period of the token is 10 minutes for China Telecom, 30 minutes for China Unicom, and 2 minutes for China Mobile. The token can be used only once.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	SpToken *string `json:"SpToken,omitempty" xml:"SpToken,omitempty"`
}

func (s GetPhoneWithTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneWithTokenRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhoneWithTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhoneWithTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhoneWithTokenRequest) GetSpToken() *string {
	return s.SpToken
}

func (s *GetPhoneWithTokenRequest) SetOwnerId(v int64) *GetPhoneWithTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *GetPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetResourceOwnerId(v int64) *GetPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetSpToken(v string) *GetPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

func (s *GetPhoneWithTokenRequest) Validate() error {
	return dara.Validate(s)
}
