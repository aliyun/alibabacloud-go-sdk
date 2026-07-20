// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CustomerNoteCreateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CustomerNoteCreateResponseBody
	GetCode() *string
	SetData(v int64) *CustomerNoteCreateResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CustomerNoteCreateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CustomerNoteCreateResponseBody
	GetMessage() *string
	SetMsg(v string) *CustomerNoteCreateResponseBody
	GetMsg() *string
	SetRequestId(v string) *CustomerNoteCreateResponseBody
	GetRequestId() *string
}

type CustomerNoteCreateResponseBody struct {
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
	// true
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// A9B725C7-3DBD-576B-AC91-F6F22AB99A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomerNoteCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerNoteCreateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CustomerNoteCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerNoteCreateResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CustomerNoteCreateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CustomerNoteCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerNoteCreateResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CustomerNoteCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerNoteCreateResponseBody) SetAccessDeniedDetail(v string) *CustomerNoteCreateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetCode(v string) *CustomerNoteCreateResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetData(v int64) *CustomerNoteCreateResponseBody {
	s.Data = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetHttpStatusCode(v int32) *CustomerNoteCreateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetMessage(v string) *CustomerNoteCreateResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetMsg(v string) *CustomerNoteCreateResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) SetRequestId(v string) *CustomerNoteCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerNoteCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
