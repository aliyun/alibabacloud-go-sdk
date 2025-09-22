// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderSessionListPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryOrderSessionListPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*QueryOrderSessionListPopResponseBodyData) *QueryOrderSessionListPopResponseBody
	GetData() []*QueryOrderSessionListPopResponseBodyData
	SetErrCode(v string) *QueryOrderSessionListPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryOrderSessionListPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryOrderSessionListPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryOrderSessionListPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryOrderSessionListPopResponseBody
	GetSuccess() *bool
}

type QueryOrderSessionListPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QueryOrderSessionListPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QueryOrderSessionListPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderSessionListPopResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryOrderSessionListPopResponseBody) GetData() []*QueryOrderSessionListPopResponseBodyData {
	return s.Data
}

func (s *QueryOrderSessionListPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryOrderSessionListPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryOrderSessionListPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryOrderSessionListPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrderSessionListPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrderSessionListPopResponseBody) SetAccessDeniedDetail(v string) *QueryOrderSessionListPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetData(v []*QueryOrderSessionListPopResponseBodyData) *QueryOrderSessionListPopResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetErrCode(v string) *QueryOrderSessionListPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetErrMessage(v string) *QueryOrderSessionListPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetHttpStatusCode(v int32) *QueryOrderSessionListPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetRequestId(v string) *QueryOrderSessionListPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) SetSuccess(v bool) *QueryOrderSessionListPopResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryOrderSessionListPopResponseBodyData struct {
	// example:
	//
	// 1
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	SignInDate *string `json:"SignInDate,omitempty" xml:"SignInDate,omitempty"`
}

func (s QueryOrderSessionListPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderSessionListPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopResponseBodyData) GetSessionId() *int64 {
	return s.SessionId
}

func (s *QueryOrderSessionListPopResponseBodyData) GetSignInDate() *string {
	return s.SignInDate
}

func (s *QueryOrderSessionListPopResponseBodyData) SetSessionId(v int64) *QueryOrderSessionListPopResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBodyData) SetSignInDate(v string) *QueryOrderSessionListPopResponseBodyData {
	s.SignInDate = &v
	return s
}

func (s *QueryOrderSessionListPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
