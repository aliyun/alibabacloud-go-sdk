// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryPipelineJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RetryPipelineJobRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RetryPipelineJobRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RetryPipelineJobRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryPipelineJobRunResponseBody
	GetSuccess() *bool
}

type RetryPipelineJobRunResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryPipelineJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *RetryPipelineJobRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryPipelineJobRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetryPipelineJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryPipelineJobRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryPipelineJobRunResponseBody) SetErrorCode(v string) *RetryPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetErrorMessage(v string) *RetryPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetRequestId(v string) *RetryPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetSuccess(v bool) *RetryPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) Validate() error {
	return dara.Validate(s)
}
