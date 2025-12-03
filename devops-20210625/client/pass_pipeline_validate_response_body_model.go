// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPassPipelineValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *PassPipelineValidateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PassPipelineValidateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PassPipelineValidateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PassPipelineValidateResponseBody
	GetSuccess() *bool
}

type PassPipelineValidateResponseBody struct {
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

func (s PassPipelineValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PassPipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *PassPipelineValidateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PassPipelineValidateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PassPipelineValidateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PassPipelineValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PassPipelineValidateResponseBody) SetErrorCode(v string) *PassPipelineValidateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetErrorMessage(v string) *PassPipelineValidateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetRequestId(v string) *PassPipelineValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetSuccess(v bool) *PassPipelineValidateResponseBody {
	s.Success = &v
	return s
}

func (s *PassPipelineValidateResponseBody) Validate() error {
	return dara.Validate(s)
}
