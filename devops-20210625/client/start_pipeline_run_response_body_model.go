// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StartPipelineRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartPipelineRunResponseBody
	GetErrorMessage() *string
	SetPipelineRunId(v int64) *StartPipelineRunResponseBody
	GetPipelineRunId() *int64
	SetRequestId(v string) *StartPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartPipelineRunResponseBody
	GetSuccess() *bool
}

type StartPipelineRunResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 1
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartPipelineRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartPipelineRunResponseBody) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *StartPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartPipelineRunResponseBody) SetErrorCode(v string) *StartPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetErrorMessage(v string) *StartPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetPipelineRunId(v int64) *StartPipelineRunResponseBody {
	s.PipelineRunId = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetRequestId(v string) *StartPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetSuccess(v bool) *StartPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *StartPipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
