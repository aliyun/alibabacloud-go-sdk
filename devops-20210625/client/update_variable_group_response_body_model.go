// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateVariableGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateVariableGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateVariableGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVariableGroupResponseBody
	GetSuccess() *bool
}

type UpdateVariableGroupResponseBody struct {
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
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVariableGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateVariableGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateVariableGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVariableGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVariableGroupResponseBody) SetErrorCode(v string) *UpdateVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetErrorMessage(v string) *UpdateVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetRequestId(v string) *UpdateVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetSuccess(v bool) *UpdateVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
