// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhoneMessageQrdlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeletePhoneMessageQrdlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeletePhoneMessageQrdlResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePhoneMessageQrdlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePhoneMessageQrdlResponseBody
	GetRequestId() *string
}

type DeletePhoneMessageQrdlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePhoneMessageQrdlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeletePhoneMessageQrdlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePhoneMessageQrdlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePhoneMessageQrdlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePhoneMessageQrdlResponseBody) SetAccessDeniedDetail(v string) *DeletePhoneMessageQrdlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) SetCode(v string) *DeletePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) SetMessage(v string) *DeletePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) SetRequestId(v string) *DeletePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponseBody) Validate() error {
	return dara.Validate(s)
}
