// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkReplStatementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecuteSparkReplStatementResponseBodyData) *ExecuteSparkReplStatementResponseBody
  GetData() *ExecuteSparkReplStatementResponseBodyData 
  SetRequestId(v string) *ExecuteSparkReplStatementResponseBody
  GetRequestId() *string 
}

type ExecuteSparkReplStatementResponseBody struct {
  // The returned data.
  Data *ExecuteSparkReplStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The request ID.
  // 
  // example:
  // 
  // 1AD222E9-E606-4A42-BF6D-8A4442913CEF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSparkReplStatementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkReplStatementResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSparkReplStatementResponseBody) GetData() *ExecuteSparkReplStatementResponseBodyData  {
  return s.Data
}

func (s *ExecuteSparkReplStatementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSparkReplStatementResponseBody) SetData(v *ExecuteSparkReplStatementResponseBodyData) *ExecuteSparkReplStatementResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteSparkReplStatementResponseBody) SetRequestId(v string) *ExecuteSparkReplStatementResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteSparkReplStatementResponseBodyData struct {
  // The ID of the Alibaba Cloud account that owns the cluster.
  // 
  // example:
  // 
  // 17108278707****
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
  // 1730968125000
  EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // StackOverflow Exception
  Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
  // The code execution result, which is a JSON string that conforms to Apache Livy.
  // 
  // example:
  // 
  // {"text/plain": 2}
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
  // 1730968125000
  StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
  // The unique ID of the code block in the Spark job.
  // 
  // example:
  // 
  // 123
  StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s ExecuteSparkReplStatementResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkReplStatementResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetAliyunUid() *int64  {
  return s.AliyunUid
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetCode() *string  {
  return s.Code
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetCodeState() *string  {
  return s.CodeState
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetCodeType() *string  {
  return s.CodeType
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetColumns() []*string  {
  return s.Columns
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetEndTime() *int64  {
  return s.EndTime
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetError() *string  {
  return s.Error
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetOutput() *string  {
  return s.Output
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetOutputType() *string  {
  return s.OutputType
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetStartTime() *int64  {
  return s.StartTime
}

func (s *ExecuteSparkReplStatementResponseBodyData) GetStatementId() *int64  {
  return s.StatementId
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetAliyunUid(v int64) *ExecuteSparkReplStatementResponseBodyData {
  s.AliyunUid = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetCode(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.Code = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetCodeState(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.CodeState = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetCodeType(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.CodeType = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetColumns(v []*string) *ExecuteSparkReplStatementResponseBodyData {
  s.Columns = v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetEndTime(v int64) *ExecuteSparkReplStatementResponseBodyData {
  s.EndTime = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetError(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.Error = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetOutput(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.Output = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetOutputType(v string) *ExecuteSparkReplStatementResponseBodyData {
  s.OutputType = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetStartTime(v int64) *ExecuteSparkReplStatementResponseBodyData {
  s.StartTime = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) SetStatementId(v int64) *ExecuteSparkReplStatementResponseBodyData {
  s.StatementId = &v
  return s
}

func (s *ExecuteSparkReplStatementResponseBodyData) Validate() error {
  return dara.Validate(s)
}

