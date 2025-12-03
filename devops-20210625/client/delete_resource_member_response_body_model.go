// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteResourceMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteResourceMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteResourceMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteResourceMemberResponseBody
	GetSuccess() *bool
}

type DeleteResourceMemberResponseBody struct {
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

func (s DeleteResourceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteResourceMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteResourceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteResourceMemberResponseBody) SetErrorCode(v string) *DeleteResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetErrorMessage(v string) *DeleteResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetRequestId(v string) *DeleteResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetSuccess(v bool) *DeleteResourceMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
