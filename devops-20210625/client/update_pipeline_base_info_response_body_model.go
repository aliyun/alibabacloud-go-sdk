// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdatePipelineBaseInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePipelineBaseInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePipelineBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePipelineBaseInfoResponseBody
	GetSuccess() *bool
}

type UpdatePipelineBaseInfoResponseBody struct {
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
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePipelineBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePipelineBaseInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePipelineBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineBaseInfoResponseBody) SetErrorCode(v string) *UpdatePipelineBaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetErrorMessage(v string) *UpdatePipelineBaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetRequestId(v string) *UpdatePipelineBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetSuccess(v bool) *UpdatePipelineBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
