// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopPipelineJobRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopPipelineJobRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopPipelineJobRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopPipelineJobRunResponseBody
	GetSuccess() *bool
}

type StopPipelineJobRunResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// "\\"
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

func (s StopPipelineJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelineJobRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopPipelineJobRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopPipelineJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPipelineJobRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopPipelineJobRunResponseBody) SetErrorCode(v string) *StopPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetErrorMessage(v string) *StopPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetRequestId(v string) *StopPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetSuccess(v bool) *StopPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) Validate() error {
	return dara.Validate(s)
}
