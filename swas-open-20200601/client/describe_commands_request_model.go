// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeCommandsRequest
	GetCommandId() *string
	SetName(v string) *DescribeCommandsRequest
	GetName() *string
	SetPageNumber(v string) *DescribeCommandsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeCommandsRequest
	GetPageSize() *string
	SetProvider(v string) *DescribeCommandsRequest
	GetProvider() *string
	SetRegionId(v string) *DescribeCommandsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCommandsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeCommandsRequestTag) *DescribeCommandsRequest
	GetTag() []*DescribeCommandsRequestTag
	SetType(v string) *DescribeCommandsRequest
	GetType() *string
}

type DescribeCommandsRequest struct {
	// The command ID.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The name of the command. Fuzzy match is not supported.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider of the public command. Valid values:
	//
	// - AlibabaCloud: queries all public commands provided by Alibaba Cloud.
	//
	// - User: queries all custom commands that you created.
	//
	// This parameter is required.
	//
	// example:
	//
	// AlibabaCloud
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the command.
	Tag []*DescribeCommandsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The language type of the command. Valid values:
	//
	// - RunBatScript: Bat command for Windows instances.
	//
	// - RunPowerShellScript: PowerShell command for Windows instances.
	//
	// - RunShellScript: Shell command for Linux instances.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCommandsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeCommandsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCommandsRequest) GetProvider() *string {
	return s.Provider
}

func (s *DescribeCommandsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommandsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommandsRequest) GetTag() []*DescribeCommandsRequestTag {
	return s.Tag
}

func (s *DescribeCommandsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeCommandsRequest) SetCommandId(v string) *DescribeCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsRequest) SetName(v string) *DescribeCommandsRequest {
	s.Name = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageNumber(v string) *DescribeCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageSize(v string) *DescribeCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsRequest) SetProvider(v string) *DescribeCommandsRequest {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsRequest) SetRegionId(v string) *DescribeCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommandsRequest) SetResourceGroupId(v string) *DescribeCommandsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsRequest) SetTag(v []*DescribeCommandsRequestTag) *DescribeCommandsRequest {
	s.Tag = v
	return s
}

func (s *DescribeCommandsRequest) SetType(v string) *DescribeCommandsRequest {
	s.Type = &v
	return s
}

func (s *DescribeCommandsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandsRequestTag struct {
	// The tag key of the command. The tag key must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command. The tag value must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommandsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCommandsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCommandsRequestTag) SetKey(v string) *DescribeCommandsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCommandsRequestTag) SetValue(v string) *DescribeCommandsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCommandsRequestTag) Validate() error {
	return dara.Validate(s)
}
