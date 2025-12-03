// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopPipelineRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopPipelineRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopPipelineRunResponseBody
	GetSuccess() *bool
}

type StopPipelineRunResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelineRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopPipelineRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopPipelineRunResponseBody) SetErrorCode(v string) *StopPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetErrorMessage(v string) *StopPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetRequestId(v string) *StopPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetSuccess(v bool) *StopPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *StopPipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
