// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdatePipelineGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePipelineGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePipelineGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePipelineGroupResponseBody
	GetSuccess() *bool
}

type UpdatePipelineGroupResponseBody struct {
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

func (s UpdatePipelineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePipelineGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePipelineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineGroupResponseBody) SetErrorCode(v string) *UpdatePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetErrorMessage(v string) *UpdatePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetRequestId(v string) *UpdatePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetSuccess(v bool) *UpdatePipelineGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
