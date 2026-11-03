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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The batch information.
	Data *ConfirmPipelineBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the confirmation to start the next batch was successful. Valid values:
	//
	// - **true**: The confirmation was successful.
	//
	// - **false**: The confirmation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to accurately query call information.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmPipelineBatchResponseBodyData struct {
	// The batch ID.
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
