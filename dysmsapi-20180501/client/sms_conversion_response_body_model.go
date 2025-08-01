// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SmsConversionResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *SmsConversionResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *SmsConversionResponseBody
	GetResponseDescription() *string
}

type SmsConversionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/180674.html).
	//
	// example:
	//
	// OK
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
}

func (s SmsConversionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionResponseBody) GoString() string {
	return s.String()
}

func (s *SmsConversionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmsConversionResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *SmsConversionResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *SmsConversionResponseBody) SetRequestId(v string) *SmsConversionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsConversionResponseBody) SetResponseCode(v string) *SmsConversionResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SmsConversionResponseBody) SetResponseDescription(v string) *SmsConversionResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SmsConversionResponseBody) Validate() error {
	return dara.Validate(s)
}
