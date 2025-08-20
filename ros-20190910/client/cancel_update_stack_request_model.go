// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpdateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCancelType(v string) *CancelUpdateStackRequest
	GetCancelType() *string
	SetRegionId(v string) *CancelUpdateStackRequest
	GetRegionId() *string
	SetStackId(v string) *CancelUpdateStackRequest
	GetStackId() *string
}

type CancelUpdateStackRequest struct {
	// The method to cancel the update operation. Valid values:
	//
	// 	- Quick: cancels the update of a stack as soon as possible.
	//
	// 	- Safe: cancels the update of a stack as safely as possible.
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
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CancelUpdateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelUpdateStackRequest) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackRequest) GetCancelType() *string {
	return s.CancelType
}

func (s *CancelUpdateStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelUpdateStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *CancelUpdateStackRequest) SetCancelType(v string) *CancelUpdateStackRequest {
	s.CancelType = &v
	return s
}

func (s *CancelUpdateStackRequest) SetRegionId(v string) *CancelUpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CancelUpdateStackRequest) SetStackId(v string) *CancelUpdateStackRequest {
	s.StackId = &v
	return s
}

func (s *CancelUpdateStackRequest) Validate() error {
	return dara.Validate(s)
}
