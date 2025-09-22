// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketOrCredentialsSignInPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *TicketOrCredentialsSignInPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v interface{}) *TicketOrCredentialsSignInPopResponseBody
	GetData() interface{}
	SetErrCode(v string) *TicketOrCredentialsSignInPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *TicketOrCredentialsSignInPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *TicketOrCredentialsSignInPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *TicketOrCredentialsSignInPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketOrCredentialsSignInPopResponseBody
	GetSuccess() *bool
}

type TicketOrCredentialsSignInPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s TicketOrCredentialsSignInPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketOrCredentialsSignInPopResponseBody) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetData() interface{} {
	return s.Data
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketOrCredentialsSignInPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetAccessDeniedDetail(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetData(v interface{}) *TicketOrCredentialsSignInPopResponseBody {
	s.Data = v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetErrCode(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetErrMessage(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetHttpStatusCode(v int32) *TicketOrCredentialsSignInPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetRequestId(v string) *TicketOrCredentialsSignInPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) SetSuccess(v bool) *TicketOrCredentialsSignInPopResponseBody {
	s.Success = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponseBody) Validate() error {
	return dara.Validate(s)
}
