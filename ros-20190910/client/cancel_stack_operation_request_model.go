// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStackOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedStackOperations(v []*string) *CancelStackOperationRequest
	GetAllowedStackOperations() []*string
	SetCancelType(v string) *CancelStackOperationRequest
	GetCancelType() *string
	SetRegionId(v string) *CancelStackOperationRequest
	GetRegionId() *string
	SetStackId(v string) *CancelStackOperationRequest
	GetStackId() *string
}

type CancelStackOperationRequest struct {
	// The operations that you want to cancel on the stack.
	AllowedStackOperations []*string `json:"AllowedStackOperations,omitempty" xml:"AllowedStackOperations,omitempty" type:"Repeated"`
	// The method that you want to use to cancel the operations. Valid values:
	//
	// 	- Quick: cancels the operations on the stack at the earliest opportunity. In this case, Resource Orchestration Service (ROS) stops scheduling new resources and stops running resources at the earliest opportunity. If you use this method, the resource status may become invalid and subsequent stack operations may be affected.
	//
	// 	- Safe (default): cancels the operations on the stack in a secure manner. In this case, ROS stops scheduling new resources and waits for running resources to be stopped.
	//
	// example:
	//
	// Safe
	CancelType *string `json:"CancelType,omitempty" xml:"CancelType,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CancelStackOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelStackOperationRequest) GoString() string {
	return s.String()
}

func (s *CancelStackOperationRequest) GetAllowedStackOperations() []*string {
	return s.AllowedStackOperations
}

func (s *CancelStackOperationRequest) GetCancelType() *string {
	return s.CancelType
}

func (s *CancelStackOperationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelStackOperationRequest) GetStackId() *string {
	return s.StackId
}

func (s *CancelStackOperationRequest) SetAllowedStackOperations(v []*string) *CancelStackOperationRequest {
	s.AllowedStackOperations = v
	return s
}

func (s *CancelStackOperationRequest) SetCancelType(v string) *CancelStackOperationRequest {
	s.CancelType = &v
	return s
}

func (s *CancelStackOperationRequest) SetRegionId(v string) *CancelStackOperationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelStackOperationRequest) SetStackId(v string) *CancelStackOperationRequest {
	s.StackId = &v
	return s
}

func (s *CancelStackOperationRequest) Validate() error {
	return dara.Validate(s)
}
