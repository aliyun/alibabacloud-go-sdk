// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCaptchaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyCaptchaResponseBody
	GetCode() *string
	SetMessage(v string) *VerifyCaptchaResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyCaptchaResponseBody
	GetRequestId() *string
	SetResult(v *VerifyCaptchaResponseBodyResult) *VerifyCaptchaResponseBody
	GetResult() *VerifyCaptchaResponseBodyResult
	SetSuccess(v bool) *VerifyCaptchaResponseBody
	GetSuccess() *bool
}

type VerifyCaptchaResponseBody struct {
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
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *VerifyCaptchaResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyCaptchaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyCaptchaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyCaptchaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCaptchaResponseBody) GetResult() *VerifyCaptchaResponseBodyResult {
	return s.Result
}

func (s *VerifyCaptchaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyCaptchaResponseBody) SetCode(v string) *VerifyCaptchaResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetMessage(v string) *VerifyCaptchaResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetRequestId(v string) *VerifyCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetResult(v *VerifyCaptchaResponseBodyResult) *VerifyCaptchaResponseBody {
	s.Result = v
	return s
}

func (s *VerifyCaptchaResponseBody) SetSuccess(v bool) *VerifyCaptchaResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyCaptchaResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyCaptchaResponseBodyResult struct {
	// example:
	//
	// true
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyCaptchaResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s VerifyCaptchaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponseBodyResult) GetVerifyResult() *bool {
	return s.VerifyResult
}

func (s *VerifyCaptchaResponseBodyResult) SetVerifyResult(v bool) *VerifyCaptchaResponseBodyResult {
	s.VerifyResult = &v
	return s
}

func (s *VerifyCaptchaResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
