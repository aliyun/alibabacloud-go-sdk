// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySignInRecordPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySignInRecordPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*QuerySignInRecordPopResponseBodyData) *QuerySignInRecordPopResponseBody
	GetData() []*QuerySignInRecordPopResponseBodyData
	SetErrCode(v string) *QuerySignInRecordPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QuerySignInRecordPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QuerySignInRecordPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QuerySignInRecordPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySignInRecordPopResponseBody
	GetSuccess() *bool
}

type QuerySignInRecordPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data []*QuerySignInRecordPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s QuerySignInRecordPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySignInRecordPopResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySignInRecordPopResponseBody) GetData() []*QuerySignInRecordPopResponseBodyData {
	return s.Data
}

func (s *QuerySignInRecordPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QuerySignInRecordPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QuerySignInRecordPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySignInRecordPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySignInRecordPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySignInRecordPopResponseBody) SetAccessDeniedDetail(v string) *QuerySignInRecordPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetData(v []*QuerySignInRecordPopResponseBodyData) *QuerySignInRecordPopResponseBody {
	s.Data = v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetErrCode(v string) *QuerySignInRecordPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetErrMessage(v string) *QuerySignInRecordPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetHttpStatusCode(v int32) *QuerySignInRecordPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetRequestId(v string) *QuerySignInRecordPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) SetSuccess(v bool) *QuerySignInRecordPopResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySignInRecordPopResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySignInRecordPopResponseBodyData struct {
	// example:
	//
	// boarding
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// nfcid
	//
	// example:
	//
	// cshdsaodhoashd
	Rfid *string `json:"Rfid,omitempty" xml:"Rfid,omitempty"`
	// sessionId
	//
	// example:
	//
	// 2001
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QuerySignInRecordPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySignInRecordPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponseBodyData) GetEvent() *string {
	return s.Event
}

func (s *QuerySignInRecordPopResponseBodyData) GetRfid() *string {
	return s.Rfid
}

func (s *QuerySignInRecordPopResponseBodyData) GetSessionId() *int64 {
	return s.SessionId
}

func (s *QuerySignInRecordPopResponseBodyData) GetTime() *string {
	return s.Time
}

func (s *QuerySignInRecordPopResponseBodyData) SetEvent(v string) *QuerySignInRecordPopResponseBodyData {
	s.Event = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetRfid(v string) *QuerySignInRecordPopResponseBodyData {
	s.Rfid = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetSessionId(v int64) *QuerySignInRecordPopResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) SetTime(v string) *QuerySignInRecordPopResponseBodyData {
	s.Time = &v
	return s
}

func (s *QuerySignInRecordPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
