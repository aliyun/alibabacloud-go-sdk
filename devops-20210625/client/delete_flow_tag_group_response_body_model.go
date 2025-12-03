// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteFlowTagGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFlowTagGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFlowTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFlowTagGroupResponseBody
	GetSuccess() *bool
}

type DeleteFlowTagGroupResponseBody struct {
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

func (s DeleteFlowTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFlowTagGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFlowTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlowTagGroupResponseBody) SetErrorCode(v string) *DeleteFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetErrorMessage(v string) *DeleteFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetRequestId(v string) *DeleteFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetSuccess(v bool) *DeleteFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
