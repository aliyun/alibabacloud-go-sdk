// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeployMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeployMachineResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyDeployMachineResponseBody
	GetResult() *bool
}

type ModifyDeployMachineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C37CE536-6C0F-4778-9B59-6D94C7F7EB63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the ECS instances are changed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyDeployMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeployMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeployMachineResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyDeployMachineResponseBody) SetRequestId(v string) *ModifyDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeployMachineResponseBody) SetResult(v bool) *ModifyDeployMachineResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyDeployMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
