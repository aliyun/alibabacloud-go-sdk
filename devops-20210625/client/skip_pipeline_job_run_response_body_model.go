// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipPipelineJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SkipPipelineJobRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SkipPipelineJobRunResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SkipPipelineJobRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkipPipelineJobRunResponseBody
	GetSuccess() *bool
}

type SkipPipelineJobRunResponseBody struct {
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

func (s SkipPipelineJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *SkipPipelineJobRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SkipPipelineJobRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SkipPipelineJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipPipelineJobRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipPipelineJobRunResponseBody) SetErrorCode(v string) *SkipPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetErrorMessage(v string) *SkipPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetRequestId(v string) *SkipPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetSuccess(v bool) *SkipPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) Validate() error {
	return dara.Validate(s)
}
