// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetDomainName() []*string
	SetIdentityCredential(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetIdentityCredential() *string
	SetIdentityCredentialNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetIdentityCredentialNo() *string
	SetIdentityCredentialType(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetIdentityCredentialType() *string
	SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
	GetUserClientIp() *string
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest struct {
	// This parameter is required.
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// This parameter is required.
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// This parameter is required.
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// This parameter is required.
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp           *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetIdentityCredential() *string {
	return s.IdentityCredential
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetIdentityCredentialNo() *string {
	return s.IdentityCredentialNo
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetIdentityCredentialType() *string {
	return s.IdentityCredentialType
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetDomainName(v []*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredential(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredentialNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredentialType(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) Validate() error {
	return dara.Validate(s)
}
