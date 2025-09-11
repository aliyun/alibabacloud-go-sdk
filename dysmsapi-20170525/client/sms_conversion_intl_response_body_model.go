// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SmsConversionIntlResponseBody
	GetCode() *string
	SetMessage(v string) *SmsConversionIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *SmsConversionIntlResponseBody
	GetRequestId() *string
}

type SmsConversionIntlResponseBody struct {
	// The response code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.101345.0.0.74326ff2J5EZyt).
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

func (s SmsConversionIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionIntlResponseBody) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *SmsConversionIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SmsConversionIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmsConversionIntlResponseBody) SetCode(v string) *SmsConversionIntlResponseBody {
	s.Code = &v
	return s
}

func (s *SmsConversionIntlResponseBody) SetMessage(v string) *SmsConversionIntlResponseBody {
	s.Message = &v
	return s
}

func (s *SmsConversionIntlResponseBody) SetRequestId(v string) *SmsConversionIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsConversionIntlResponseBody) Validate() error {
	return dara.Validate(s)
}
