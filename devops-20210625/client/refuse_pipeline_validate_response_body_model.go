// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefusePipelineValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RefusePipelineValidateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RefusePipelineValidateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RefusePipelineValidateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefusePipelineValidateResponseBody
	GetSuccess() *bool
}

type RefusePipelineValidateResponseBody struct {
	// example:
	//
	// ""
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

func (s RefusePipelineValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefusePipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *RefusePipelineValidateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RefusePipelineValidateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RefusePipelineValidateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefusePipelineValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefusePipelineValidateResponseBody) SetErrorCode(v string) *RefusePipelineValidateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetErrorMessage(v string) *RefusePipelineValidateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetRequestId(v string) *RefusePipelineValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetSuccess(v bool) *RefusePipelineValidateResponseBody {
	s.Success = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) Validate() error {
	return dara.Validate(s)
}
