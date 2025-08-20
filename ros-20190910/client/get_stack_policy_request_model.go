// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetStackPolicyRequest
	GetRegionId() *string
	SetStackId(v string) *GetStackPolicyRequest
	GetStackId() *string
}

type GetStackPolicyRequest struct {
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

func (s GetStackPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetStackPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackPolicyRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetStackPolicyRequest) SetRegionId(v string) *GetStackPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackPolicyRequest) SetStackId(v string) *GetStackPolicyRequest {
	s.StackId = &v
	return s
}

func (s *GetStackPolicyRequest) Validate() error {
	return dara.Validate(s)
}
