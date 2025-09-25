// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateQualityCheckTaskResponseBody
	GetCost() *int64
	SetData(v *CreateQualityCheckTaskResponseBodyData) *CreateQualityCheckTaskResponseBody
	GetData() *CreateQualityCheckTaskResponseBodyData
	SetDataType(v string) *CreateQualityCheckTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateQualityCheckTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateQualityCheckTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQualityCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityCheckTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateQualityCheckTaskResponseBody
	GetTime() *string
}

type CreateQualityCheckTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                  `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *CreateQualityCheckTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
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

func (s CreateQualityCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateQualityCheckTaskResponseBody) GetData() *CreateQualityCheckTaskResponseBodyData {
	return s.Data
}

func (s *CreateQualityCheckTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateQualityCheckTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateQualityCheckTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQualityCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityCheckTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateQualityCheckTaskResponseBody) SetCost(v int64) *CreateQualityCheckTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetData(v *CreateQualityCheckTaskResponseBodyData) *CreateQualityCheckTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetDataType(v string) *CreateQualityCheckTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetErrCode(v string) *CreateQualityCheckTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetMessage(v string) *CreateQualityCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetRequestId(v string) *CreateQualityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetSuccess(v bool) *CreateQualityCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetTime(v string) *CreateQualityCheckTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateQualityCheckTaskResponseBodyData struct {
	// taskId
	//
	// example:
	//
	// 172373500521
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateQualityCheckTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateQualityCheckTaskResponseBodyData) SetTaskId(v string) *CreateQualityCheckTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
