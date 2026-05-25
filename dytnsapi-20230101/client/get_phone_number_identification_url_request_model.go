// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetPhoneNumberIdentificationUrlRequest
	GetAuthCode() *string
	SetIp(v string) *GetPhoneNumberIdentificationUrlRequest
	GetIp() *string
	SetOutId(v string) *GetPhoneNumberIdentificationUrlRequest
	GetOutId() *string
	SetOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetPhoneNumberIdentificationUrlRequest
	GetPhoneNumber() *string
	SetRememberPhoneNumber(v bool) *GetPhoneNumberIdentificationUrlRequest
	GetRememberPhoneNumber() *bool
	SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest
	GetResourceOwnerId() *int64
}

type GetPhoneNumberIdentificationUrlRequest struct {
	// The authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// K***9i7CIe
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The IP address of the subscriber\\"s phone.
	//
	// example:
	//
	// 114.124.***.13
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The external ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 149b03d2-a749-4e6e-8f5b-34******5815
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the subscriber. The phone number is in the Mobile Station International Subscriber Directory Number (MSISDN) format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 628211****113
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Specifies whether to remember the phone number.
	//
	// example:
	//
	// true
	RememberPhoneNumber  *bool   `json:"RememberPhoneNumber,omitempty" xml:"RememberPhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetIp() *string {
	return s.Ip
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetRememberPhoneNumber() *bool {
	return s.RememberPhoneNumber
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhoneNumberIdentificationUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetAuthCode(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.AuthCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetIp(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.Ip = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetOutId(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.OutId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetPhoneNumber(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetRememberPhoneNumber(v bool) *GetPhoneNumberIdentificationUrlRequest {
	s.RememberPhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) Validate() error {
	return dara.Validate(s)
}
