// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitChatQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *SubmitChatQuestionResponseBody
	GetCost() *int64
	SetData(v *SubmitChatQuestionResponseBodyData) *SubmitChatQuestionResponseBody
	GetData() *SubmitChatQuestionResponseBodyData
	SetDataType(v string) *SubmitChatQuestionResponseBody
	GetDataType() *string
	SetErrCode(v string) *SubmitChatQuestionResponseBody
	GetErrCode() *string
	SetMessage(v string) *SubmitChatQuestionResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitChatQuestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitChatQuestionResponseBody
	GetSuccess() *bool
	SetTime(v string) *SubmitChatQuestionResponseBody
	GetTime() *string
}

type SubmitChatQuestionResponseBody struct {
	// example:
	//
	// null
	Cost *int64                              `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *SubmitChatQuestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 915AAAB9-4908-5224-9E53-9E9D7D0AA94B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s SubmitChatQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitChatQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *SubmitChatQuestionResponseBody) GetData() *SubmitChatQuestionResponseBodyData {
	return s.Data
}

func (s *SubmitChatQuestionResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *SubmitChatQuestionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SubmitChatQuestionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitChatQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitChatQuestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitChatQuestionResponseBody) GetTime() *string {
	return s.Time
}

func (s *SubmitChatQuestionResponseBody) SetCost(v int64) *SubmitChatQuestionResponseBody {
	s.Cost = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetData(v *SubmitChatQuestionResponseBodyData) *SubmitChatQuestionResponseBody {
	s.Data = v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetDataType(v string) *SubmitChatQuestionResponseBody {
	s.DataType = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetErrCode(v string) *SubmitChatQuestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetMessage(v string) *SubmitChatQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetRequestId(v string) *SubmitChatQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetSuccess(v bool) *SubmitChatQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetTime(v string) *SubmitChatQuestionResponseBody {
	s.Time = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitChatQuestionResponseBodyData struct {
	// example:
	//
	// 1869307330227937280
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s SubmitChatQuestionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitChatQuestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponseBodyData) GetBatchId() *string {
	return s.BatchId
}

func (s *SubmitChatQuestionResponseBodyData) SetBatchId(v string) *SubmitChatQuestionResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *SubmitChatQuestionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
