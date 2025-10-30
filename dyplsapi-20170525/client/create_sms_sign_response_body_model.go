// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSmsSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSmsSignResponseBody
	GetCode() *string
	SetData(v *CreateSmsSignResponseBodyData) *CreateSmsSignResponseBody
	GetData() *CreateSmsSignResponseBodyData
	SetMessage(v string) *CreateSmsSignResponseBody
	GetMessage() *string
	SetSuccess(v bool) *CreateSmsSignResponseBody
	GetSuccess() *bool
}

type CreateSmsSignResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateSmsSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsSignResponseBody) GetData() *CreateSmsSignResponseBodyData {
	return s.Data
}

func (s *CreateSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSmsSignResponseBody) SetAccessDeniedDetail(v string) *CreateSmsSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetCode(v string) *CreateSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetData(v *CreateSmsSignResponseBodyData) *CreateSmsSignResponseBody {
	s.Data = v
	return s
}

func (s *CreateSmsSignResponseBody) SetMessage(v string) *CreateSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetSuccess(v bool) *CreateSmsSignResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSmsSignResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSmsSignResponseBodyData struct {
	// 短信接收者号码签名串(加到短信内容中供解析真实被叫号码)
	//
	// example:
	//
	// sign23343466
	CalledNoSign *string `json:"CalledNoSign,omitempty" xml:"CalledNoSign,omitempty"`
}

func (s CreateSmsSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBodyData) GetCalledNoSign() *string {
	return s.CalledNoSign
}

func (s *CreateSmsSignResponseBodyData) SetCalledNoSign(v string) *CreateSmsSignResponseBodyData {
	s.CalledNoSign = &v
	return s
}

func (s *CreateSmsSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
