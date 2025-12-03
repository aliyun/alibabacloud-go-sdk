// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateVariableGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateVariableGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateVariableGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVariableGroupResponseBody
	GetSuccess() *bool
	SetVariableGroupId(v int64) *CreateVariableGroupResponseBody
	GetVariableGroupId() *int64
}

type CreateVariableGroupResponseBody struct {
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
	// example:
	//
	// 1234
	VariableGroupId *int64 `json:"variableGroupId,omitempty" xml:"variableGroupId,omitempty"`
}

func (s CreateVariableGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateVariableGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateVariableGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVariableGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVariableGroupResponseBody) GetVariableGroupId() *int64 {
	return s.VariableGroupId
}

func (s *CreateVariableGroupResponseBody) SetErrorCode(v string) *CreateVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetErrorMessage(v string) *CreateVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetRequestId(v string) *CreateVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetSuccess(v bool) *CreateVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetVariableGroupId(v int64) *CreateVariableGroupResponseBody {
	s.VariableGroupId = &v
	return s
}

func (s *CreateVariableGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
