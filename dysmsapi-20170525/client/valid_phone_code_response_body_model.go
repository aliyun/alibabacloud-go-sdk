// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidPhoneCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ValidPhoneCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ValidPhoneCodeResponseBody
	GetCode() *string
	SetData(v bool) *ValidPhoneCodeResponseBody
	GetData() *bool
	SetMessage(v string) *ValidPhoneCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValidPhoneCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidPhoneCodeResponseBody
	GetSuccess() *bool
}

type ValidPhoneCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ValidPhoneCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidPhoneCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ValidPhoneCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ValidPhoneCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValidPhoneCodeResponseBody) GetData() *bool {
	return s.Data
}

func (s *ValidPhoneCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidPhoneCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidPhoneCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidPhoneCodeResponseBody) SetAccessDeniedDetail(v string) *ValidPhoneCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) SetCode(v string) *ValidPhoneCodeResponseBody {
	s.Code = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) SetData(v bool) *ValidPhoneCodeResponseBody {
	s.Data = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) SetMessage(v string) *ValidPhoneCodeResponseBody {
	s.Message = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) SetRequestId(v string) *ValidPhoneCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) SetSuccess(v bool) *ValidPhoneCodeResponseBody {
	s.Success = &v
	return s
}

func (s *ValidPhoneCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
