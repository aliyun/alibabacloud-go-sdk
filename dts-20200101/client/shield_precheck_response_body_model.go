// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShieldPrecheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ShieldPrecheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ShieldPrecheckResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ShieldPrecheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ShieldPrecheckResponseBody
	GetSuccess() *bool
}

type ShieldPrecheckResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6F4B5BC4-34B1-49C9-9C8F-C8F16AC4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ShieldPrecheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShieldPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ShieldPrecheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ShieldPrecheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShieldPrecheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ShieldPrecheckResponseBody) SetErrCode(v string) *ShieldPrecheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetErrMessage(v string) *ShieldPrecheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetRequestId(v string) *ShieldPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetSuccess(v bool) *ShieldPrecheckResponseBody {
	s.Success = &v
	return s
}

func (s *ShieldPrecheckResponseBody) Validate() error {
	return dara.Validate(s)
}
