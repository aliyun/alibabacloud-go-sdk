// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShortUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddShortUrlResponseBody
	GetCode() *string
	SetData(v *AddShortUrlResponseBodyData) *AddShortUrlResponseBody
	GetData() *AddShortUrlResponseBodyData
	SetMessage(v string) *AddShortUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddShortUrlResponseBody
	GetRequestId() *string
}

type AddShortUrlResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the short URL.
	Data *AddShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddShortUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddShortUrlResponseBody) GetData() *AddShortUrlResponseBodyData {
	return s.Data
}

func (s *AddShortUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddShortUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddShortUrlResponseBody) SetCode(v string) *AddShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *AddShortUrlResponseBody) SetData(v *AddShortUrlResponseBodyData) *AddShortUrlResponseBody {
	s.Data = v
	return s
}

func (s *AddShortUrlResponseBody) SetMessage(v string) *AddShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *AddShortUrlResponseBody) SetRequestId(v string) *AddShortUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShortUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddShortUrlResponseBodyData struct {
	// The time when the short URL expires.
	//
	// > The value of **ExpireDate*	- is on the hour.
	//
	// example:
	//
	// 2021-09-19 00:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The short URL.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	// The source URL.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s AddShortUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddShortUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponseBodyData) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *AddShortUrlResponseBodyData) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *AddShortUrlResponseBodyData) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *AddShortUrlResponseBodyData) SetExpireDate(v string) *AddShortUrlResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *AddShortUrlResponseBodyData) SetShortUrl(v string) *AddShortUrlResponseBodyData {
	s.ShortUrl = &v
	return s
}

func (s *AddShortUrlResponseBodyData) SetSourceUrl(v string) *AddShortUrlResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *AddShortUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
