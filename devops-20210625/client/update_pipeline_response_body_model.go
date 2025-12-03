// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdatePipelineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePipelineResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePipelineResponseBody
	GetSuccess() *bool
}

type UpdatePipelineResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePipelineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineResponseBody) SetErrorCode(v string) *UpdatePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetErrorMessage(v string) *UpdatePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetSuccess(v bool) *UpdatePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
