// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPipelineBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfirmPipelineBatchResponseBody
	GetCode() *string
	SetData(v *ConfirmPipelineBatchResponseBodyData) *ConfirmPipelineBatchResponseBody
	GetData() *ConfirmPipelineBatchResponseBodyData
	SetErrorCode(v string) *ConfirmPipelineBatchResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ConfirmPipelineBatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmPipelineBatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfirmPipelineBatchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ConfirmPipelineBatchResponseBody
	GetTraceId() *string
}

type ConfirmPipelineBatchResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The batch information.
	Data *ConfirmPipelineBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the processing of the next batch started as required. Valid values:
	//
	// 	- **true**: The processing started.
	//
	// 	- **false**: The processing could not start.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ConfirmPipelineBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPipelineBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfirmPipelineBatchResponseBody) GetData() *ConfirmPipelineBatchResponseBodyData {
	return s.Data
}

func (s *ConfirmPipelineBatchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ConfirmPipelineBatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmPipelineBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmPipelineBatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmPipelineBatchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ConfirmPipelineBatchResponseBody) SetCode(v string) *ConfirmPipelineBatchResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetData(v *ConfirmPipelineBatchResponseBodyData) *ConfirmPipelineBatchResponseBody {
	s.Data = v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetErrorCode(v string) *ConfirmPipelineBatchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetMessage(v string) *ConfirmPipelineBatchResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetRequestId(v string) *ConfirmPipelineBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetSuccess(v bool) *ConfirmPipelineBatchResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetTraceId(v string) *ConfirmPipelineBatchResponseBody {
	s.TraceId = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) Validate() error {
	return dara.Validate(s)
}

type ConfirmPipelineBatchResponseBodyData struct {
	// The ID of the batch.
	//
	// example:
	//
	// e2e-vds-feh-***
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ConfirmPipelineBatchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPipelineBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponseBodyData) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ConfirmPipelineBatchResponseBodyData) SetPipelineId(v string) *ConfirmPipelineBatchResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
