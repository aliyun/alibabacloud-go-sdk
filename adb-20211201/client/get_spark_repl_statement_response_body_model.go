// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkReplStatementResponseBodyData) *GetSparkReplStatementResponseBody
	GetData() *GetSparkReplStatementResponseBodyData
	SetRequestId(v string) *GetSparkReplStatementResponseBody
	GetRequestId() *string
}

type GetSparkReplStatementResponseBody struct {
	// The returned data.
	Data *GetSparkReplStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkReplStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplStatementResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkReplStatementResponseBody) GetData() *GetSparkReplStatementResponseBodyData {
	return s.Data
}

func (s *GetSparkReplStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkReplStatementResponseBody) SetData(v *GetSparkReplStatementResponseBodyData) *GetSparkReplStatementResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkReplStatementResponseBody) SetRequestId(v string) *GetSparkReplStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkReplStatementResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkReplStatementResponseBodyData struct {
	// The ID of the Alibaba Cloud account to which the instance belongs.
	//
	// example:
	//
	// 144740799645****
	AliyunUid *int64 `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The executed code.
	//
	// example:
	//
	// print(1+1)
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution status of the code. Valid values:
	//
	// - CANCELLED: The execution is canceled.
	//
	// - RUNNING: The execution is in progress.
	//
	// - SUCCEEDED: The execution is completed.
	//
	// - ERROR: The execution failed.
	//
	// example:
	//
	// RUNNING
	CodeState *string `json:"CodeState,omitempty" xml:"CodeState,omitempty"`
	// The type of the code. Valid values:
	//
	// - PYTHON
	//
	// - SCALA
	//
	// example:
	//
	// PYTHON
	CodeType *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// The array of column names.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The end time of the task. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1730968194000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// stackoverflow error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The output result, which is a string in JSON format.
	//
	// example:
	//
	// {"text/plain": 2}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The type of the output result.
	//
	// example:
	//
	// TEXT
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The start time of the task. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1730968194000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The unique ID of the code block in this Spark job.
	//
	// example:
	//
	// 123
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s GetSparkReplStatementResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkReplStatementResponseBodyData) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *GetSparkReplStatementResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetSparkReplStatementResponseBodyData) GetCodeState() *string {
	return s.CodeState
}

func (s *GetSparkReplStatementResponseBodyData) GetCodeType() *string {
	return s.CodeType
}

func (s *GetSparkReplStatementResponseBodyData) GetColumns() []*string {
	return s.Columns
}

func (s *GetSparkReplStatementResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSparkReplStatementResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetSparkReplStatementResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *GetSparkReplStatementResponseBodyData) GetOutputType() *string {
	return s.OutputType
}

func (s *GetSparkReplStatementResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSparkReplStatementResponseBodyData) GetStatementId() *int64 {
	return s.StatementId
}

func (s *GetSparkReplStatementResponseBodyData) SetAliyunUid(v int64) *GetSparkReplStatementResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetCode(v string) *GetSparkReplStatementResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetCodeState(v string) *GetSparkReplStatementResponseBodyData {
	s.CodeState = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetCodeType(v string) *GetSparkReplStatementResponseBodyData {
	s.CodeType = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetColumns(v []*string) *GetSparkReplStatementResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetEndTime(v int64) *GetSparkReplStatementResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetError(v string) *GetSparkReplStatementResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetOutput(v string) *GetSparkReplStatementResponseBodyData {
	s.Output = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetOutputType(v string) *GetSparkReplStatementResponseBodyData {
	s.OutputType = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetStartTime(v int64) *GetSparkReplStatementResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) SetStatementId(v int64) *GetSparkReplStatementResponseBodyData {
	s.StatementId = &v
	return s
}

func (s *GetSparkReplStatementResponseBodyData) Validate() error {
	return dara.Validate(s)
}
