// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v string) *SetDeletionProtectionRequest
	GetDeletionProtection() *string
	SetRegionId(v string) *SetDeletionProtectionRequest
	GetRegionId() *string
	SetStackId(v string) *SetDeletionProtectionRequest
	GetStackId() *string
}

type SetDeletionProtectionRequest struct {
	// Indicates whether stack deletion protection is enabled. Valid values:
	//
	// 	- Enabled: enables the stack deletion protection.
	//
	// 	- Disabled (default): Resource stack deletion protection is Disabled. You can use the console or API(DeleteStack) to release the stack resources.
	//
	// >  The deletion of nested stacks is the same as the root stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
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
	// The delete protection attribute of a nested stack is determined by the root stack and remains unchanged from the root stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s SetDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionRequest) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *SetDeletionProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDeletionProtectionRequest) GetStackId() *string {
	return s.StackId
}

func (s *SetDeletionProtectionRequest) SetDeletionProtection(v string) *SetDeletionProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetRegionId(v string) *SetDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetStackId(v string) *SetDeletionProtectionRequest {
	s.StackId = &v
	return s
}

func (s *SetDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
