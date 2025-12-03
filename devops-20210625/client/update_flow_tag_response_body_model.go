// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateFlowTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFlowTagResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateFlowTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFlowTagResponseBody
	GetSuccess() *bool
}

type UpdateFlowTagResponseBody struct {
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

func (s UpdateFlowTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFlowTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFlowTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFlowTagResponseBody) SetErrorCode(v string) *UpdateFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetErrorMessage(v string) *UpdateFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetRequestId(v string) *UpdateFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetSuccess(v bool) *UpdateFlowTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFlowTagResponseBody) Validate() error {
	return dara.Validate(s)
}
