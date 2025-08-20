// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteOptions(v []*string) *DeleteStackRequest
	GetDeleteOptions() []*string
	SetParallelism(v int64) *DeleteStackRequest
	GetParallelism() *int64
	SetRamRoleName(v string) *DeleteStackRequest
	GetRamRoleName() *string
	SetRegionId(v string) *DeleteStackRequest
	GetRegionId() *string
	SetRetainAllResources(v bool) *DeleteStackRequest
	GetRetainAllResources() *bool
	SetRetainResources(v []*string) *DeleteStackRequest
	GetRetainResources() []*string
	SetStackId(v string) *DeleteStackRequest
	GetStackId() *string
}

type DeleteStackRequest struct {
	// The options for deleting the stack.
	DeleteOptions []*string `json:"DeleteOptions,omitempty" xml:"DeleteOptions,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	//
	//
	// > -  If you set this parameter to an integer that is greater than 0, the integer is used. If you set this parameter to 0 or leave this parameter empty, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > -  If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack, such as an update operation.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\\
	//
	// ROS assumes the role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.\\
	//
	// If you leave this parameter empty when you call the DeleteStack operation, ROS cannot assume the existing RAM role that is associated with the stack. If you want ROS to assume a RAM role, you must specify this parameter. If no RAM roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\\
	//
	// The name of the RAM role can be up to 64 bytes in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain all resources in the stack.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	// The resources that you want to retain.
	//
	// example:
	//
	// WebServer
	RetainResources []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DeleteStackRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackRequest) GetDeleteOptions() []*string {
	return s.DeleteOptions
}

func (s *DeleteStackRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *DeleteStackRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DeleteStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStackRequest) GetRetainAllResources() *bool {
	return s.RetainAllResources
}

func (s *DeleteStackRequest) GetRetainResources() []*string {
	return s.RetainResources
}

func (s *DeleteStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *DeleteStackRequest) SetDeleteOptions(v []*string) *DeleteStackRequest {
	s.DeleteOptions = v
	return s
}

func (s *DeleteStackRequest) SetParallelism(v int64) *DeleteStackRequest {
	s.Parallelism = &v
	return s
}

func (s *DeleteStackRequest) SetRamRoleName(v string) *DeleteStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *DeleteStackRequest) SetRegionId(v string) *DeleteStackRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackRequest) SetRetainAllResources(v bool) *DeleteStackRequest {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteStackRequest) SetRetainResources(v []*string) *DeleteStackRequest {
	s.RetainResources = v
	return s
}

func (s *DeleteStackRequest) SetStackId(v string) *DeleteStackRequest {
	s.StackId = &v
	return s
}

func (s *DeleteStackRequest) Validate() error {
	return dara.Validate(s)
}
