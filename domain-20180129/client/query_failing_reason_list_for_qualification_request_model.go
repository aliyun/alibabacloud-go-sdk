// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailingReasonListForQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryFailingReasonListForQualificationRequest
	GetInstanceId() *string
	SetLang(v string) *QueryFailingReasonListForQualificationRequest
	GetLang() *string
	SetLimit(v int32) *QueryFailingReasonListForQualificationRequest
	GetLimit() *int32
	SetQualificationType(v string) *QueryFailingReasonListForQualificationRequest
	GetQualificationType() *string
	SetUserClientIp(v string) *QueryFailingReasonListForQualificationRequest
	GetUserClientIp() *string
}

type QueryFailingReasonListForQualificationRequest struct {
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
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
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

func (s QueryFailingReasonListForQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFailingReasonListForQualificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryFailingReasonListForQualificationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryFailingReasonListForQualificationRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryFailingReasonListForQualificationRequest) GetQualificationType() *string {
	return s.QualificationType
}

func (s *QueryFailingReasonListForQualificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryFailingReasonListForQualificationRequest) SetInstanceId(v string) *QueryFailingReasonListForQualificationRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetLang(v string) *QueryFailingReasonListForQualificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetLimit(v int32) *QueryFailingReasonListForQualificationRequest {
	s.Limit = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetQualificationType(v string) *QueryFailingReasonListForQualificationRequest {
	s.QualificationType = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetUserClientIp(v string) *QueryFailingReasonListForQualificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) Validate() error {
	return dara.Validate(s)
}
