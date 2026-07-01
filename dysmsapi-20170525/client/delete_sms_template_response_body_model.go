// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSmsTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *DeleteSmsTemplateResponseBody
	GetTemplateCode() *string
}

type DeleteSmsTemplateResponseBody struct {
	// 请求状态码。
	//
	// - 返回OK代表请求成功。
	//
	// - 其他错误码，请参见[API错误码](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// CCA2BCFF-2BA7-427C-90EE-AC6994748607
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 已删除的模板Code。
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DeleteSmsTemplateResponseBody) SetCode(v string) *DeleteSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetMessage(v string) *DeleteSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetRequestId(v string) *DeleteSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetTemplateCode(v string) *DeleteSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
