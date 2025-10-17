// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkReplStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CancelSparkReplStatementResponseBodyData) *CancelSparkReplStatementResponseBody
	GetData() *CancelSparkReplStatementResponseBodyData
	SetRequestId(v string) *CancelSparkReplStatementResponseBody
	GetRequestId() *string
}

type CancelSparkReplStatementResponseBody struct {
	// The returned data.
	Data *CancelSparkReplStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelSparkReplStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkReplStatementResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSparkReplStatementResponseBody) GetData() *CancelSparkReplStatementResponseBodyData {
	return s.Data
}

func (s *CancelSparkReplStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSparkReplStatementResponseBody) SetData(v *CancelSparkReplStatementResponseBodyData) *CancelSparkReplStatementResponseBody {
	s.Data = v
	return s
}

func (s *CancelSparkReplStatementResponseBody) SetRequestId(v string) *CancelSparkReplStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSparkReplStatementResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelSparkReplStatementResponseBodyData struct {
	// The ID of the Alibaba Cloud account that owns the cluster.
	//
	// example:
	//
	// 190063530332****
	AliyunUid *int64 `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The code that is executed.
	//
	// example:
	//
	// print(1+1)
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The code execution status. Valid values:
	//
	// 	- CANCELLED
	//
	// 	- RUNNING
	//
	// 	- SUCCEEDED
	//
	// 	- ERROR
	//
	// example:
	//
	// RUNNING
	CodeState *string `json:"CodeState,omitempty" xml:"CodeState,omitempty"`
	// The code type. Valid values:
	//
	// 	- SCALA
	//
	// 	- PYTHON
	//
	// example:
	//
	// PYTHON
	CodeType *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// The column names.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The end time of the execution. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1730968056000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// StackOverflow Exception:
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The code execution result, which is a JSON string that conforms to Apache Livy.
	//
	// example:
	//
	// {"text/plain": "2"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The execution result type, which is in the JSON format. Valid values:
	//
	// 	- TEXT: the text content that conforms to Apache Livy.
	//
	// 	- TABLE: the table content that conforms to Apache Livy.
	//
	// example:
	//
	// TEXT
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The start time of the execution. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1730968056000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The unique ID of the code block in the Spark job.
	//
	// example:
	//
	// 123
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s CancelSparkReplStatementResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkReplStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelSparkReplStatementResponseBodyData) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *CancelSparkReplStatementResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *CancelSparkReplStatementResponseBodyData) GetCodeState() *string {
	return s.CodeState
}

func (s *CancelSparkReplStatementResponseBodyData) GetCodeType() *string {
	return s.CodeType
}

func (s *CancelSparkReplStatementResponseBodyData) GetColumns() []*string {
	return s.Columns
}

func (s *CancelSparkReplStatementResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CancelSparkReplStatementResponseBodyData) GetError() *string {
	return s.Error
}

func (s *CancelSparkReplStatementResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *CancelSparkReplStatementResponseBodyData) GetOutputType() *string {
	return s.OutputType
}

func (s *CancelSparkReplStatementResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CancelSparkReplStatementResponseBodyData) GetStatementId() *int64 {
	return s.StatementId
}

func (s *CancelSparkReplStatementResponseBodyData) SetAliyunUid(v int64) *CancelSparkReplStatementResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetCode(v string) *CancelSparkReplStatementResponseBodyData {
	s.Code = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetCodeState(v string) *CancelSparkReplStatementResponseBodyData {
	s.CodeState = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetCodeType(v string) *CancelSparkReplStatementResponseBodyData {
	s.CodeType = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetColumns(v []*string) *CancelSparkReplStatementResponseBodyData {
	s.Columns = v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetEndTime(v int64) *CancelSparkReplStatementResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetError(v string) *CancelSparkReplStatementResponseBodyData {
	s.Error = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetOutput(v string) *CancelSparkReplStatementResponseBodyData {
	s.Output = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetOutputType(v string) *CancelSparkReplStatementResponseBodyData {
	s.OutputType = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetStartTime(v int64) *CancelSparkReplStatementResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) SetStatementId(v int64) *CancelSparkReplStatementResponseBodyData {
	s.StatementId = &v
	return s
}

func (s *CancelSparkReplStatementResponseBodyData) Validate() error {
	return dara.Validate(s)
}
