// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateResourceMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateResourceMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateResourceMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateResourceMemberResponseBody
	GetSuccess() *bool
}

type CreateResourceMemberResponseBody struct {
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

func (s CreateResourceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateResourceMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateResourceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResourceMemberResponseBody) SetErrorCode(v string) *CreateResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetErrorMessage(v string) *CreateResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetRequestId(v string) *CreateResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetSuccess(v bool) *CreateResourceMemberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourceMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
