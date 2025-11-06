// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelQualificationVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CancelQualificationVerificationRequest
	GetInstanceId() *string
	SetLang(v string) *CancelQualificationVerificationRequest
	GetLang() *string
	SetQualificationType(v string) *CancelQualificationVerificationRequest
	GetQualificationType() *string
	SetUserClientIp(v string) *CancelQualificationVerificationRequest
	GetUserClientIp() *string
}

type CancelQualificationVerificationRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// knet
	QualificationType *string `json:"QualificationType,omitempty" xml:"QualificationType,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelQualificationVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelQualificationVerificationRequest) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelQualificationVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelQualificationVerificationRequest) GetQualificationType() *string {
	return s.QualificationType
}

func (s *CancelQualificationVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CancelQualificationVerificationRequest) SetInstanceId(v string) *CancelQualificationVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetLang(v string) *CancelQualificationVerificationRequest {
	s.Lang = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetQualificationType(v string) *CancelQualificationVerificationRequest {
	s.QualificationType = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetUserClientIp(v string) *CancelQualificationVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *CancelQualificationVerificationRequest) Validate() error {
	return dara.Validate(s)
}
