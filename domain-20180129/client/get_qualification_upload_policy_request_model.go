// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetQualificationUploadPolicyRequest
	GetLang() *string
	SetUserClientIp(v string) *GetQualificationUploadPolicyRequest
	GetUserClientIp() *string
}

type GetQualificationUploadPolicyRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s GetQualificationUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *GetQualificationUploadPolicyRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetQualificationUploadPolicyRequest) SetLang(v string) *GetQualificationUploadPolicyRequest {
	s.Lang = &v
	return s
}

func (s *GetQualificationUploadPolicyRequest) SetUserClientIp(v string) *GetQualificationUploadPolicyRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetQualificationUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
