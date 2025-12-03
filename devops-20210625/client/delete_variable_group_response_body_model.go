// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteVariableGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteVariableGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteVariableGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVariableGroupResponseBody
	GetSuccess() *bool
}

type DeleteVariableGroupResponseBody struct {
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

func (s DeleteVariableGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVariableGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteVariableGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteVariableGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVariableGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVariableGroupResponseBody) SetErrorCode(v string) *DeleteVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetErrorMessage(v string) *DeleteVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetRequestId(v string) *DeleteVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetSuccess(v bool) *DeleteVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
