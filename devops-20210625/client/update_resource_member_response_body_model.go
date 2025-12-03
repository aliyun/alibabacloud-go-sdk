// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateResourceMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateResourceMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateResourceMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateResourceMemberResponseBody
	GetSuccess() *bool
}

type UpdateResourceMemberResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
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

func (s UpdateResourceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateResourceMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateResourceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateResourceMemberResponseBody) SetErrorCode(v string) *UpdateResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetErrorMessage(v string) *UpdateResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetRequestId(v string) *UpdateResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetSuccess(v bool) *UpdateResourceMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
