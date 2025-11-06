// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForRegistrantProfileRealNameVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest
	GetLang() *string
	SetRegistrantProfileID(v int64) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest
	GetRegistrantProfileID() *int64
	SetUserClientIp(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest
	GetUserClientIp() *string
}

type QueryFailReasonForRegistrantProfileRealNameVerificationRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	RegistrantProfileID *int64 `json:"RegistrantProfileID,omitempty" xml:"RegistrantProfileID,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) GetRegistrantProfileID() *int64 {
	return s.RegistrantProfileID
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetLang(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileID(v int64) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileID = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) Validate() error {
	return dara.Validate(s)
}
