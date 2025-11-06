// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForDomainRealNameVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryFailReasonForDomainRealNameVerificationRequest
	GetDomainName() *string
	SetLang(v string) *QueryFailReasonForDomainRealNameVerificationRequest
	GetLang() *string
	SetRealNameVerificationAction(v string) *QueryFailReasonForDomainRealNameVerificationRequest
	GetRealNameVerificationAction() *string
	SetUserClientIp(v string) *QueryFailReasonForDomainRealNameVerificationRequest
	GetUserClientIp() *string
}

type QueryFailReasonForDomainRealNameVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACTIVATE
	RealNameVerificationAction *string `json:"RealNameVerificationAction,omitempty" xml:"RealNameVerificationAction,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) GetRealNameVerificationAction() *string {
	return s.RealNameVerificationAction
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetDomainName(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetLang(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetRealNameVerificationAction(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.RealNameVerificationAction = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetUserClientIp(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) Validate() error {
	return dara.Validate(s)
}
