// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetQualificationVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetQualificationVerificationRequest
	GetInstanceId() *string
	SetLang(v string) *ResetQualificationVerificationRequest
	GetLang() *string
	SetUserClientIp(v string) *ResetQualificationVerificationRequest
	GetUserClientIp() *string
}

type ResetQualificationVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// S20181*****85212
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ResetQualificationVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetQualificationVerificationRequest) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetQualificationVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetQualificationVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ResetQualificationVerificationRequest) SetInstanceId(v string) *ResetQualificationVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetQualificationVerificationRequest) SetLang(v string) *ResetQualificationVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ResetQualificationVerificationRequest) SetUserClientIp(v string) *ResetQualificationVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *ResetQualificationVerificationRequest) Validate() error {
	return dara.Validate(s)
}
