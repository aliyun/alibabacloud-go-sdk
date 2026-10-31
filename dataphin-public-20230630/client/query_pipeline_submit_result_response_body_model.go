// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineSubmitResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPipelineSubmitResultResponseBody
	GetCode() *string
	SetData(v *QueryPipelineSubmitResultResponseBodyData) *QueryPipelineSubmitResultResponseBody
	GetData() *QueryPipelineSubmitResultResponseBodyData
	SetHttpStatusCode(v int32) *QueryPipelineSubmitResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryPipelineSubmitResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPipelineSubmitResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPipelineSubmitResultResponseBody
	GetSuccess() *bool
}

type QueryPipelineSubmitResultResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous execution result.
	Data *QueryPipelineSubmitResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPipelineSubmitResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPipelineSubmitResultResponseBody) GetData() *QueryPipelineSubmitResultResponseBodyData {
	return s.Data
}

func (s *QueryPipelineSubmitResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryPipelineSubmitResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPipelineSubmitResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPipelineSubmitResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPipelineSubmitResultResponseBody) SetCode(v string) *QueryPipelineSubmitResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) SetData(v *QueryPipelineSubmitResultResponseBodyData) *QueryPipelineSubmitResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) SetHttpStatusCode(v int32) *QueryPipelineSubmitResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) SetMessage(v string) *QueryPipelineSubmitResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) SetRequestId(v string) *QueryPipelineSubmitResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) SetSuccess(v bool) *QueryPipelineSubmitResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPipelineSubmitResultResponseBodyData struct {
	// The query ID for asynchronous execution.
	//
	// example:
	//
	// 123
	AsyncId *int64 `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	// The check result of the commit details. The value is a JSON character string and is returned only for the asynchronous commit process. The structure is consistent with the commit details on the product page.
	//
	// example:
	//
	// {"test01":"success"}
	CheckDetail *string `json:"CheckDetail,omitempty" xml:"CheckDetail,omitempty"`
	// The error code returned when an execution exception occurs.
	//
	// example:
	//
	// DPN.Pipeline.InnerError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when an execution exception occurs.
	//
	// example:
	//
	// NullPointException
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The server host information for request execution.
	//
	// example:
	//
	// hostName:hostIp
	HostMachine *string `json:"HostMachine,omitempty" xml:"HostMachine,omitempty"`
	// The ID of the schedule node for the pipeline task.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The primary key ID of the pipeline.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The current execution status. Valid values: SUCCESS: execution succeeded. FAILED: execution failed. RUNNING: execution in progress.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The publish ID generated after the pipeline task is submitted, which is used to publish the task through the publish domain.
	//
	// example:
	//
	// 123
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// The version number generated when the pipeline task is submitted.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryPipelineSubmitResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetAsyncId() *int64 {
	return s.AsyncId
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetCheckDetail() *string {
	return s.CheckDetail
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *QueryPipelineSubmitResultResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetAsyncId(v int64) *QueryPipelineSubmitResultResponseBodyData {
	s.AsyncId = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetCheckDetail(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.CheckDetail = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetErrorCode(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetErrorMessage(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetHostMachine(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetNodeId(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetPipelineId(v int64) *QueryPipelineSubmitResultResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetStatus(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetSubmitId(v int64) *QueryPipelineSubmitResultResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) SetVersion(v string) *QueryPipelineSubmitResultResponseBodyData {
	s.Version = &v
	return s
}

func (s *QueryPipelineSubmitResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
