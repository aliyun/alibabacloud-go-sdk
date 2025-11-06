// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQualificationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryQualificationDetailRequest
	GetInstanceId() *string
	SetLang(v string) *QueryQualificationDetailRequest
	GetLang() *string
	SetQualificationType(v string) *QueryQualificationDetailRequest
	GetQualificationType() *string
	SetUserClientIp(v string) *QueryQualificationDetailRequest
	GetUserClientIp() *string
}

type QueryQualificationDetailRequest struct {
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

func (s QueryQualificationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQualificationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryQualificationDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryQualificationDetailRequest) GetQualificationType() *string {
	return s.QualificationType
}

func (s *QueryQualificationDetailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryQualificationDetailRequest) SetInstanceId(v string) *QueryQualificationDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetLang(v string) *QueryQualificationDetailRequest {
	s.Lang = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetQualificationType(v string) *QueryQualificationDetailRequest {
	s.QualificationType = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetUserClientIp(v string) *QueryQualificationDetailRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryQualificationDetailRequest) Validate() error {
	return dara.Validate(s)
}
