// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyIntelligentCaptchaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyIntelligentCaptchaResponseBody
	GetCode() *string
	SetMessage(v string) *VerifyIntelligentCaptchaResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyIntelligentCaptchaResponseBody
	GetRequestId() *string
	SetResult(v *VerifyIntelligentCaptchaResponseBodyResult) *VerifyIntelligentCaptchaResponseBody
	GetResult() *VerifyIntelligentCaptchaResponseBodyResult
	SetSuccess(v bool) *VerifyIntelligentCaptchaResponseBody
	GetSuccess() *bool
}

type VerifyIntelligentCaptchaResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 95784F***D39FDC5
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *VerifyIntelligentCaptchaResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyIntelligentCaptchaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyIntelligentCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyIntelligentCaptchaResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyIntelligentCaptchaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyIntelligentCaptchaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyIntelligentCaptchaResponseBody) GetResult() *VerifyIntelligentCaptchaResponseBodyResult {
	return s.Result
}

func (s *VerifyIntelligentCaptchaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyIntelligentCaptchaResponseBody) SetCode(v string) *VerifyIntelligentCaptchaResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBody) SetMessage(v string) *VerifyIntelligentCaptchaResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBody) SetRequestId(v string) *VerifyIntelligentCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBody) SetResult(v *VerifyIntelligentCaptchaResponseBodyResult) *VerifyIntelligentCaptchaResponseBody {
	s.Result = v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBody) SetSuccess(v bool) *VerifyIntelligentCaptchaResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyIntelligentCaptchaResponseBodyResult struct {
	CertifyId  *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// example:
	//
	// true
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyIntelligentCaptchaResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s VerifyIntelligentCaptchaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) GetCertifyId() *string {
	return s.CertifyId
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) GetVerifyResult() *bool {
	return s.VerifyResult
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) SetCertifyId(v string) *VerifyIntelligentCaptchaResponseBodyResult {
	s.CertifyId = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) SetVerifyCode(v string) *VerifyIntelligentCaptchaResponseBodyResult {
	s.VerifyCode = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) SetVerifyResult(v bool) *VerifyIntelligentCaptchaResponseBodyResult {
	s.VerifyResult = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
