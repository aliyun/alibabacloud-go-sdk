// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateFlowTagGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFlowTagGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateFlowTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFlowTagGroupResponseBody
	GetSuccess() *bool
}

type UpdateFlowTagGroupResponseBody struct {
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

func (s UpdateFlowTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFlowTagGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFlowTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFlowTagGroupResponseBody) SetErrorCode(v string) *UpdateFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetErrorMessage(v string) *UpdateFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetRequestId(v string) *UpdateFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetSuccess(v bool) *UpdateFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
