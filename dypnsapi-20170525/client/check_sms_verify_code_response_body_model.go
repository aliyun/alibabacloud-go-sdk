// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSmsVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckSmsVerifyCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CheckSmsVerifyCodeResponseBody
	GetCode() *string
	SetMessage(v string) *CheckSmsVerifyCodeResponseBody
	GetMessage() *string
	SetModel(v *CheckSmsVerifyCodeResponseBodyModel) *CheckSmsVerifyCodeResponseBody
	GetModel() *CheckSmsVerifyCodeResponseBodyModel
	SetSuccess(v bool) *CheckSmsVerifyCodeResponseBody
	GetSuccess() *bool
}

type CheckSmsVerifyCodeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [Response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11174283.0.0.70c5616bkj38Wa).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *CheckSmsVerifyCodeResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSmsVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSmsVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckSmsVerifyCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckSmsVerifyCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckSmsVerifyCodeResponseBody) GetModel() *CheckSmsVerifyCodeResponseBodyModel {
	return s.Model
}

func (s *CheckSmsVerifyCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckSmsVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *CheckSmsVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetCode(v string) *CheckSmsVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetMessage(v string) *CheckSmsVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetModel(v *CheckSmsVerifyCodeResponseBodyModel) *CheckSmsVerifyCodeResponseBody {
	s.Model = v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetSuccess(v bool) *CheckSmsVerifyCodeResponseBody {
	s.Success = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckSmsVerifyCodeResponseBodyModel struct {
	// The external ID.
	//
	// example:
	//
	// 1212312
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- PASS: The verification is successful.
	//
	// 	- UNKNOWN: The verification failed.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CheckSmsVerifyCodeResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CheckSmsVerifyCodeResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponseBodyModel) GetOutId() *string {
	return s.OutId
}

func (s *CheckSmsVerifyCodeResponseBodyModel) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CheckSmsVerifyCodeResponseBodyModel) SetOutId(v string) *CheckSmsVerifyCodeResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBodyModel) SetVerifyResult(v string) *CheckSmsVerifyCodeResponseBodyModel {
	s.VerifyResult = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
