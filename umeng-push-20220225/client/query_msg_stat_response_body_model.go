// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMsgStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMsgStatResponseBody
	GetCode() *string
	SetData(v *QueryMsgStatResponseBodyData) *QueryMsgStatResponseBody
	GetData() *QueryMsgStatResponseBodyData
	SetHttpStatusCode(v int32) *QueryMsgStatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryMsgStatResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMsgStatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMsgStatResponseBody
	GetSuccess() *bool
}

type QueryMsgStatResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryMsgStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMsgStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMsgStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMsgStatResponseBody) GetData() *QueryMsgStatResponseBodyData {
	return s.Data
}

func (s *QueryMsgStatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryMsgStatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMsgStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMsgStatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMsgStatResponseBody) SetCode(v string) *QueryMsgStatResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetData(v *QueryMsgStatResponseBodyData) *QueryMsgStatResponseBody {
	s.Data = v
	return s
}

func (s *QueryMsgStatResponseBody) SetHttpStatusCode(v int32) *QueryMsgStatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetMessage(v string) *QueryMsgStatResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetRequestId(v string) *QueryMsgStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetSuccess(v bool) *QueryMsgStatResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMsgStatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMsgStatResponseBodyData struct {
	// example:
	//
	// 1
	Accept *int64 `json:"Accept,omitempty" xml:"Accept,omitempty"`
	// example:
	//
	// 1
	Arrive *int64 `json:"Arrive,omitempty" xml:"Arrive,omitempty"`
	// example:
	//
	// 0
	ClosePush *int64 `json:"ClosePush,omitempty" xml:"ClosePush,omitempty"`
	// example:
	//
	// 0
	Dismiss *int64 `json:"Dismiss,omitempty" xml:"Dismiss,omitempty"`
	// example:
	//
	// ufe29y2167046828041801
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// 1
	Open *int64 `json:"Open,omitempty" xml:"Open,omitempty"`
	// example:
	//
	// 1
	Sent *int64 `json:"Sent,omitempty" xml:"Sent,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryMsgStatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMsgStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponseBodyData) GetAccept() *int64 {
	return s.Accept
}

func (s *QueryMsgStatResponseBodyData) GetArrive() *int64 {
	return s.Arrive
}

func (s *QueryMsgStatResponseBodyData) GetClosePush() *int64 {
	return s.ClosePush
}

func (s *QueryMsgStatResponseBodyData) GetDismiss() *int64 {
	return s.Dismiss
}

func (s *QueryMsgStatResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMsgStatResponseBodyData) GetOpen() *int64 {
	return s.Open
}

func (s *QueryMsgStatResponseBodyData) GetSent() *int64 {
	return s.Sent
}

func (s *QueryMsgStatResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryMsgStatResponseBodyData) SetAccept(v int64) *QueryMsgStatResponseBodyData {
	s.Accept = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetArrive(v int64) *QueryMsgStatResponseBodyData {
	s.Arrive = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetClosePush(v int64) *QueryMsgStatResponseBodyData {
	s.ClosePush = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetDismiss(v int64) *QueryMsgStatResponseBodyData {
	s.Dismiss = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetMsgId(v string) *QueryMsgStatResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetOpen(v int64) *QueryMsgStatResponseBodyData {
	s.Open = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetSent(v int64) *QueryMsgStatResponseBodyData {
	s.Sent = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetStatus(v int32) *QueryMsgStatResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
