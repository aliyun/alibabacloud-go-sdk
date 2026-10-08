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
	// The domain names to be verified in bulk.
	//
	// This parameter is required.
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The Base64-encoded content of the identity credential file.
	//
	// This parameter is required.
	IdentityCredential *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	// The ID number of the identity credential.
	//
	// This parameter is required.
	IdentityCredentialNo *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	// The type of the identity credential. Valid values: IDC, Passport, and OfficerAcademy.
	//
	// This parameter is required.
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	// The response language. Valid values: zh-CN and en-US. The default is en-US.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The client IP address.
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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
