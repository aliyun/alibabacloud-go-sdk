// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *TerminateCallRequest
	GetCallId() *string
	SetInstanceId(v string) *TerminateCallRequest
	GetInstanceId() *string
}

type TerminateCallRequest struct {
	// Call ID
	//
	// > You can obtain it by calling the DescribeJob API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1528189846043
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TerminateCallRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateCallRequest) GoString() string {
	return s.String()
}

func (s *TerminateCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *TerminateCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateCallRequest) SetCallId(v string) *TerminateCallRequest {
	s.CallId = &v
	return s
}

func (s *TerminateCallRequest) SetInstanceId(v string) *TerminateCallRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateCallRequest) Validate() error {
	return dara.Validate(s)
}
