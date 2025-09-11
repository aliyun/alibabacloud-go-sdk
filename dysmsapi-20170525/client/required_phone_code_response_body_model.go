// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequiredPhoneCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RequiredPhoneCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RequiredPhoneCodeResponseBody
	GetCode() *string
	SetData(v string) *RequiredPhoneCodeResponseBody
	GetData() *string
	SetMessage(v string) *RequiredPhoneCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *RequiredPhoneCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RequiredPhoneCodeResponseBody
	GetSuccess() *bool
}

type RequiredPhoneCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequiredPhoneCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RequiredPhoneCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RequiredPhoneCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RequiredPhoneCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *RequiredPhoneCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *RequiredPhoneCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RequiredPhoneCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RequiredPhoneCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RequiredPhoneCodeResponseBody) SetAccessDeniedDetail(v string) *RequiredPhoneCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) SetCode(v string) *RequiredPhoneCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) SetData(v string) *RequiredPhoneCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) SetMessage(v string) *RequiredPhoneCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) SetRequestId(v string) *RequiredPhoneCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) SetSuccess(v bool) *RequiredPhoneCodeResponseBody {
	s.Success = &v
	return s
}

func (s *RequiredPhoneCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
