// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConversionDataResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *ConversionDataResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *ConversionDataResponseBody
	GetResponseDescription() *string
}

type ConversionDataResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code. Returning OK means the request was successful. For other error codes, please refer to the [Error codes](https://help.aliyun.com/document_detail/180674.html) list.
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

func (s ConversionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataResponseBody) GoString() string {
	return s.String()
}

func (s *ConversionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConversionDataResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *ConversionDataResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *ConversionDataResponseBody) SetRequestId(v string) *ConversionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConversionDataResponseBody) SetResponseCode(v string) *ConversionDataResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *ConversionDataResponseBody) SetResponseDescription(v string) *ConversionDataResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *ConversionDataResponseBody) Validate() error {
	return dara.Validate(s)
}
