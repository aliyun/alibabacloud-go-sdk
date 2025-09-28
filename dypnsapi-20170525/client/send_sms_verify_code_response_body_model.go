// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendSmsVerifyCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendSmsVerifyCodeResponseBody
	GetCode() *string
	SetMessage(v string) *SendSmsVerifyCodeResponseBody
	GetMessage() *string
	SetModel(v *SendSmsVerifyCodeResponseBodyModel) *SendSmsVerifyCodeResponseBody
	GetModel() *SendSmsVerifyCodeResponseBodyModel
	SetRequestId(v string) *SendSmsVerifyCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendSmsVerifyCodeResponseBody
	GetSuccess() *bool
}

type SendSmsVerifyCodeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. If OK is returned, the request is successful. For more information, see [Response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11174283.0.0.70c5616bkj38Wa).
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
	Model     *SendSmsVerifyCodeResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendSmsVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSmsVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendSmsVerifyCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendSmsVerifyCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendSmsVerifyCodeResponseBody) GetModel() *SendSmsVerifyCodeResponseBodyModel {
	return s.Model
}

func (s *SendSmsVerifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSmsVerifyCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendSmsVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *SendSmsVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetCode(v string) *SendSmsVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetMessage(v string) *SendSmsVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetModel(v *SendSmsVerifyCodeResponseBodyModel) *SendSmsVerifyCodeResponseBody {
	s.Model = v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetRequestId(v string) *SendSmsVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetSuccess(v bool) *SendSmsVerifyCodeResponseBody {
	s.Success = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendSmsVerifyCodeResponseBodyModel struct {
	// The business ID.
	//
	// example:
	//
	// 112231421412414124123^4
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 1231231313
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// API-reqelekrqkllkkewrlwrjlsdfsdf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification code.
	//
	// example:
	//
	// 42324
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s SendSmsVerifyCodeResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s SendSmsVerifyCodeResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponseBodyModel) GetBizId() *string {
	return s.BizId
}

func (s *SendSmsVerifyCodeResponseBodyModel) GetOutId() *string {
	return s.OutId
}

func (s *SendSmsVerifyCodeResponseBodyModel) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSmsVerifyCodeResponseBodyModel) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetBizId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.BizId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetOutId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetRequestId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.RequestId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetVerifyCode(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.VerifyCode = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
