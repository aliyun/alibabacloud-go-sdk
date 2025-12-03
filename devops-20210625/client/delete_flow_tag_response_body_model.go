// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteFlowTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFlowTagResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFlowTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFlowTagResponseBody
	GetSuccess() *bool
}

type DeleteFlowTagResponseBody struct {
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

func (s DeleteFlowTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFlowTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFlowTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlowTagResponseBody) SetErrorCode(v string) *DeleteFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetErrorMessage(v string) *DeleteFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetRequestId(v string) *DeleteFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetSuccess(v bool) *DeleteFlowTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowTagResponseBody) Validate() error {
	return dara.Validate(s)
}
