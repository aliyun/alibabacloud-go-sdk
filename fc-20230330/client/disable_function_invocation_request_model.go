// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFunctionInvocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbortOngoingRequest(v bool) *DisableFunctionInvocationRequest
	GetAbortOngoingRequest() *bool
	SetReason(v string) *DisableFunctionInvocationRequest
	GetReason() *string
}

type DisableFunctionInvocationRequest struct {
	// Specifies whether to immediately terminate all ongoing requests.
	//
	// example:
	//
	// false
	AbortOngoingRequest *bool `json:"abortOngoingRequest,omitempty" xml:"abortOngoingRequest,omitempty"`
	// The reason for disabling the function\\"s invocation.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s DisableFunctionInvocationRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableFunctionInvocationRequest) GoString() string {
	return s.String()
}

func (s *DisableFunctionInvocationRequest) GetAbortOngoingRequest() *bool {
	return s.AbortOngoingRequest
}

func (s *DisableFunctionInvocationRequest) GetReason() *string {
	return s.Reason
}

func (s *DisableFunctionInvocationRequest) SetAbortOngoingRequest(v bool) *DisableFunctionInvocationRequest {
	s.AbortOngoingRequest = &v
	return s
}

func (s *DisableFunctionInvocationRequest) SetReason(v string) *DisableFunctionInvocationRequest {
	s.Reason = &v
	return s
}

func (s *DisableFunctionInvocationRequest) Validate() error {
	return dara.Validate(s)
}
