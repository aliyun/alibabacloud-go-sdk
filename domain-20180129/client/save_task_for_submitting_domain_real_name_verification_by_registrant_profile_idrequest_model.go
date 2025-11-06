// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
	GetDomainName() *string
	SetInstanceId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
	GetInstanceId() *string
	SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
	GetUserClientIp() *string
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetDomainName(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetInstanceId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetRegistrantProfileId(v int64) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) Validate() error {
	return dara.Validate(s)
}
