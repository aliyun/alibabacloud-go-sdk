// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConversionDataIntlResponseBody
	GetCode() *string
	SetMessage(v string) *ConversionDataIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConversionDataIntlResponseBody
	GetRequestId() *string
}

type ConversionDataIntlResponseBody struct {
	// 状态码。取值：
	//
	// - OK：代表请求成功。
	//
	// - 其他错误码，请参见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
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
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConversionDataIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataIntlResponseBody) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConversionDataIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConversionDataIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConversionDataIntlResponseBody) SetCode(v string) *ConversionDataIntlResponseBody {
	s.Code = &v
	return s
}

func (s *ConversionDataIntlResponseBody) SetMessage(v string) *ConversionDataIntlResponseBody {
	s.Message = &v
	return s
}

func (s *ConversionDataIntlResponseBody) SetRequestId(v string) *ConversionDataIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConversionDataIntlResponseBody) Validate() error {
	return dara.Validate(s)
}
