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
	// The status code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.101345.0.0.74326ff2J5EZyt).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
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
