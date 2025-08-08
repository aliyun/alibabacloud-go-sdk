// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPushAnalysisTaskDetailResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryPushAnalysisTaskDetailResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryPushAnalysisTaskDetailResponseBodyResultContent) *QueryPushAnalysisTaskDetailResponseBody
	GetResultContent() *QueryPushAnalysisTaskDetailResponseBodyResultContent
	SetResultMessage(v string) *QueryPushAnalysisTaskDetailResponseBody
	GetResultMessage() *string
}

type QueryPushAnalysisTaskDetailResponseBody struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                               `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryPushAnalysisTaskDetailResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                               `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryPushAnalysisTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushAnalysisTaskDetailResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryPushAnalysisTaskDetailResponseBody) GetResultContent() *QueryPushAnalysisTaskDetailResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryPushAnalysisTaskDetailResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryPushAnalysisTaskDetailResponseBody) SetRequestId(v string) *QueryPushAnalysisTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBody) SetResultCode(v string) *QueryPushAnalysisTaskDetailResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBody) SetResultContent(v *QueryPushAnalysisTaskDetailResponseBodyResultContent) *QueryPushAnalysisTaskDetailResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBody) SetResultMessage(v string) *QueryPushAnalysisTaskDetailResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPushAnalysisTaskDetailResponseBodyResultContent struct {
	Data *QueryPushAnalysisTaskDetailResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryPushAnalysisTaskDetailResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskDetailResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContent) GetData() *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	return s.Data
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContent) SetData(v *QueryPushAnalysisTaskDetailResponseBodyResultContentData) *QueryPushAnalysisTaskDetailResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type QueryPushAnalysisTaskDetailResponseBodyResultContentData struct {
	Duration       *string  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime        *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PushArrivalNum *float32 `json:"PushArrivalNum,omitempty" xml:"PushArrivalNum,omitempty"`
	PushNum        *float32 `json:"PushNum,omitempty" xml:"PushNum,omitempty"`
	PushSuccessNum *float32 `json:"PushSuccessNum,omitempty" xml:"PushSuccessNum,omitempty"`
	StartTime      *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId         *int64   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryPushAnalysisTaskDetailResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskDetailResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetDuration() *string {
	return s.Duration
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetPushArrivalNum() *float32 {
	return s.PushArrivalNum
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetPushNum() *float32 {
	return s.PushNum
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetPushSuccessNum() *float32 {
	return s.PushSuccessNum
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetDuration(v string) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.Duration = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetEndTime(v int64) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.EndTime = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetPushArrivalNum(v float32) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.PushArrivalNum = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetPushNum(v float32) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.PushNum = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetPushSuccessNum(v float32) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.PushSuccessNum = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetStartTime(v int64) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.StartTime = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) SetTaskId(v int64) *QueryPushAnalysisTaskDetailResponseBodyResultContentData {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
