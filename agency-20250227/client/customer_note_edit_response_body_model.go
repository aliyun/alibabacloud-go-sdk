// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteEditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CustomerNoteEditResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CustomerNoteEditResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CustomerNoteEditResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CustomerNoteEditResponseBody
	GetMessage() *string
	SetMsg(v string) *CustomerNoteEditResponseBody
	GetMsg() *string
	SetRequestId(v string) *CustomerNoteEditResponseBody
	GetRequestId() *string
}

type CustomerNoteEditResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 成功
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomerNoteEditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteEditResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerNoteEditResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CustomerNoteEditResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerNoteEditResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CustomerNoteEditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerNoteEditResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CustomerNoteEditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerNoteEditResponseBody) SetAccessDeniedDetail(v string) *CustomerNoteEditResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CustomerNoteEditResponseBody) SetCode(v string) *CustomerNoteEditResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerNoteEditResponseBody) SetHttpStatusCode(v int32) *CustomerNoteEditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CustomerNoteEditResponseBody) SetMessage(v string) *CustomerNoteEditResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerNoteEditResponseBody) SetMsg(v string) *CustomerNoteEditResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerNoteEditResponseBody) SetRequestId(v string) *CustomerNoteEditResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerNoteEditResponseBody) Validate() error {
	return dara.Validate(s)
}
