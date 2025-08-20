// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputOption(v string) *GetStackInstanceRequest
	GetOutputOption() *string
	SetRegionId(v string) *GetStackInstanceRequest
	GetRegionId() *string
	SetStackGroupName(v string) *GetStackInstanceRequest
	GetStackGroupName() *string
	SetStackInstanceAccountId(v string) *GetStackInstanceRequest
	GetStackInstanceAccountId() *string
	SetStackInstanceRegionId(v string) *GetStackInstanceRequest
	GetStackInstanceRegionId() *string
}

type GetStackInstanceRequest struct {
	// Specifies whether to return the Outputs parameter. The Outputs parameter specifies the outputs of the stack. Valid values:
	//
	// 	- Enabled: returns the Outputs parameter.
	//
	// 	- Disabled (default): does not return the Outputs parameter.
	//
	// >  The Outputs parameter requires a long period of time to calculate. If you do not require the outputs of the stack, we recommend that you set OutputOption to Disabled to improve the response speed of the API operation.
	//
	// example:
	//
	// Disabled
	OutputOption *string `json:"OutputOption,omitempty" xml:"OutputOption,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the destination account to which the stack belongs.
	//
	// 	- If the stack group is granted self-managed permissions, the stack belongs to an Alibaba Cloud account.
	//
	// 	- If the stack group is granted service-managed permissions, the stack belongs to a member in a resource directory.
	//
	// > For more information about the destination account, see [Overview](https://help.aliyun.com/document_detail/154578.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 151266687691****
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	// The region ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	StackInstanceRegionId *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
}

func (s GetStackInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetStackInstanceRequest) GetOutputOption() *string {
	return s.OutputOption
}

func (s *GetStackInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackInstanceRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetStackInstanceRequest) GetStackInstanceAccountId() *string {
	return s.StackInstanceAccountId
}

func (s *GetStackInstanceRequest) GetStackInstanceRegionId() *string {
	return s.StackInstanceRegionId
}

func (s *GetStackInstanceRequest) SetOutputOption(v string) *GetStackInstanceRequest {
	s.OutputOption = &v
	return s
}

func (s *GetStackInstanceRequest) SetRegionId(v string) *GetStackInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackGroupName(v string) *GetStackInstanceRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceAccountId(v string) *GetStackInstanceRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceRegionId(v string) *GetStackInstanceRequest {
	s.StackInstanceRegionId = &v
	return s
}

func (s *GetStackInstanceRequest) Validate() error {
	return dara.Validate(s)
}
