// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGuestRfidPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindGuestRfidPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *BindGuestRfidPopResponseBody
	GetData() *bool
	SetErrCode(v string) *BindGuestRfidPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *BindGuestRfidPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *BindGuestRfidPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *BindGuestRfidPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindGuestRfidPopResponseBody
	GetSuccess() *bool
}

type BindGuestRfidPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindGuestRfidPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindGuestRfidPopResponseBody) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindGuestRfidPopResponseBody) GetData() *bool {
	return s.Data
}

func (s *BindGuestRfidPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *BindGuestRfidPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *BindGuestRfidPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindGuestRfidPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindGuestRfidPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindGuestRfidPopResponseBody) SetAccessDeniedDetail(v string) *BindGuestRfidPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetData(v bool) *BindGuestRfidPopResponseBody {
	s.Data = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetErrCode(v string) *BindGuestRfidPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetErrMessage(v string) *BindGuestRfidPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetHttpStatusCode(v int32) *BindGuestRfidPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetRequestId(v string) *BindGuestRfidPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) SetSuccess(v bool) *BindGuestRfidPopResponseBody {
	s.Success = &v
	return s
}

func (s *BindGuestRfidPopResponseBody) Validate() error {
	return dara.Validate(s)
}
