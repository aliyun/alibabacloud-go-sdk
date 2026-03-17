// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappVerifyCodeRequest
	GetCustSpaceId() *string
	SetLocale(v string) *GetChatappVerifyCodeRequest
	GetLocale() *string
	SetMethod(v string) *GetChatappVerifyCodeRequest
	GetMethod() *string
	SetOwnerId(v int64) *GetChatappVerifyCodeRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetChatappVerifyCodeRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetChatappVerifyCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappVerifyCodeRequest
	GetResourceOwnerId() *int64
}

type GetChatappVerifyCodeRequest struct {
	// The space ID of the RAM user within the ISV account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 229393838*****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The language. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The verification method.
	//
	// Valid values:
	//
	// 	- Voice: sends the verification code via phone call.
	//
	// 	- sms: sends the verification code via SMS.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	Method  *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetChatappVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *GetChatappVerifyCodeRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappVerifyCodeRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetChatappVerifyCodeRequest) GetMethod() *string {
	return s.Method
}

func (s *GetChatappVerifyCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappVerifyCodeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappVerifyCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappVerifyCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappVerifyCodeRequest) SetCustSpaceId(v string) *GetChatappVerifyCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetLocale(v string) *GetChatappVerifyCodeRequest {
	s.Locale = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetMethod(v string) *GetChatappVerifyCodeRequest {
	s.Method = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetOwnerId(v int64) *GetChatappVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetPhoneNumber(v string) *GetChatappVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetResourceOwnerAccount(v string) *GetChatappVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) SetResourceOwnerId(v int64) *GetChatappVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
