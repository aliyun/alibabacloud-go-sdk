// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteAbacAuthorizationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteAbacAuthorizationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteAbacAuthorizationResponseBody
	GetRequestId() *string
	SetResult(v string) *DeleteAbacAuthorizationResponseBody
	GetResult() *string
	SetSuccess(v bool) *DeleteAbacAuthorizationResponseBody
	GetSuccess() *bool
}

type DeleteAbacAuthorizationResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 207176D7-A9B3-55CE-A9DA-14E223A31913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAbacAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAbacAuthorizationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAbacAuthorizationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAbacAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAbacAuthorizationResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteAbacAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAbacAuthorizationResponseBody) SetErrorCode(v string) *DeleteAbacAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAbacAuthorizationResponseBody) SetErrorMessage(v string) *DeleteAbacAuthorizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAbacAuthorizationResponseBody) SetRequestId(v string) *DeleteAbacAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAbacAuthorizationResponseBody) SetResult(v string) *DeleteAbacAuthorizationResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteAbacAuthorizationResponseBody) SetSuccess(v bool) *DeleteAbacAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAbacAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
