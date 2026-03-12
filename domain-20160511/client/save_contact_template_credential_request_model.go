// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v int64) *SaveContactTemplateCredentialRequest
	GetContactTemplateId() *int64
	SetCredential(v string) *SaveContactTemplateCredentialRequest
	GetCredential() *string
	SetCredentialNo(v string) *SaveContactTemplateCredentialRequest
	GetCredentialNo() *string
	SetLang(v string) *SaveContactTemplateCredentialRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveContactTemplateCredentialRequest
	GetUserClientIp() *string
}

type SaveContactTemplateCredentialRequest struct {
	// This parameter is required.
	ContactTemplateId *int64 `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	// This parameter is required.
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// This parameter is required.
	CredentialNo *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveContactTemplateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateCredentialRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveContactTemplateCredentialRequest) GetCredential() *string {
	return s.Credential
}

func (s *SaveContactTemplateCredentialRequest) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *SaveContactTemplateCredentialRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveContactTemplateCredentialRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveContactTemplateCredentialRequest) SetContactTemplateId(v int64) *SaveContactTemplateCredentialRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveContactTemplateCredentialRequest) SetCredential(v string) *SaveContactTemplateCredentialRequest {
	s.Credential = &v
	return s
}

func (s *SaveContactTemplateCredentialRequest) SetCredentialNo(v string) *SaveContactTemplateCredentialRequest {
	s.CredentialNo = &v
	return s
}

func (s *SaveContactTemplateCredentialRequest) SetLang(v string) *SaveContactTemplateCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveContactTemplateCredentialRequest) SetUserClientIp(v string) *SaveContactTemplateCredentialRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveContactTemplateCredentialRequest) Validate() error {
	return dara.Validate(s)
}
