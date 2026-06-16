// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CheckAppVerifyCodeRequest
	GetBizId() *string
	SetCode(v string) *CheckAppVerifyCodeRequest
	GetCode() *string
	SetTarget(v string) *CheckAppVerifyCodeRequest
	GetTarget() *string
	SetType(v string) *CheckAppVerifyCodeRequest
	GetType() *string
}

type CheckAppVerifyCodeRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The verification code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The phone number or email address.
	//
	// example:
	//
	// docker.io
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The recipient type: phone or email.
	//
	// example:
	//
	// question
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckAppVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAppVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *CheckAppVerifyCodeRequest) GetBizId() *string {
	return s.BizId
}

func (s *CheckAppVerifyCodeRequest) GetCode() *string {
	return s.Code
}

func (s *CheckAppVerifyCodeRequest) GetTarget() *string {
	return s.Target
}

func (s *CheckAppVerifyCodeRequest) GetType() *string {
	return s.Type
}

func (s *CheckAppVerifyCodeRequest) SetBizId(v string) *CheckAppVerifyCodeRequest {
	s.BizId = &v
	return s
}

func (s *CheckAppVerifyCodeRequest) SetCode(v string) *CheckAppVerifyCodeRequest {
	s.Code = &v
	return s
}

func (s *CheckAppVerifyCodeRequest) SetTarget(v string) *CheckAppVerifyCodeRequest {
	s.Target = &v
	return s
}

func (s *CheckAppVerifyCodeRequest) SetType(v string) *CheckAppVerifyCodeRequest {
	s.Type = &v
	return s
}

func (s *CheckAppVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
