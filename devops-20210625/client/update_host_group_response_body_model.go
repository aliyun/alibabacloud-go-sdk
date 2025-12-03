// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateHostGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateHostGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateHostGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHostGroupResponseBody
	GetSuccess() *bool
}

type UpdateHostGroupResponseBody struct {
	// example:
	//
	// ""
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

func (s UpdateHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateHostGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHostGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHostGroupResponseBody) SetErrorCode(v string) *UpdateHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetErrorMessage(v string) *UpdateHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetRequestId(v string) *UpdateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetSuccess(v bool) *UpdateHostGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHostGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
