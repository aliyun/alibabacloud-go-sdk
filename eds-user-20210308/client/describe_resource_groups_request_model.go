// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *DescribeResourceGroupsRequest
	GetAgentType() *string
	SetAliyunResourceGroupIds(v []*string) *DescribeResourceGroupsRequest
	GetAliyunResourceGroupIds() []*string
	SetBusinessChannel(v string) *DescribeResourceGroupsRequest
	GetBusinessChannel() *string
	SetNeedContainResourceGroupWithOfficeSite(v int64) *DescribeResourceGroupsRequest
	GetNeedContainResourceGroupWithOfficeSite() *int64
	SetPageNumber(v int32) *DescribeResourceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResourceGroupsRequest
	GetPageSize() *int32
	SetPlatform(v string) *DescribeResourceGroupsRequest
	GetPlatform() *string
	SetResourceClassification(v string) *DescribeResourceGroupsRequest
	GetResourceClassification() *string
	SetResourceGroupIds(v []*string) *DescribeResourceGroupsRequest
	GetResourceGroupIds() []*string
	SetResourceGroupName(v string) *DescribeResourceGroupsRequest
	GetResourceGroupName() *string
}

type DescribeResourceGroupsRequest struct {
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The cloud platform resource group ID.
	AliyunResourceGroupIds []*string `json:"AliyunResourceGroupIds,omitempty" xml:"AliyunResourceGroupIds,omitempty" type:"Repeated"`
	// The channel tag.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	NeedContainResourceGroupWithOfficeSite *int64 `json:"NeedContainResourceGroupWithOfficeSite,omitempty" xml:"NeedContainResourceGroupWithOfficeSite,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// > Fixed value: AliyunConsole.
	//
	// > - For Elastic Desktop Service Enterprise Edition, set this parameter to AliyunConsole. Other platforms are not publicly available.
	//
	// example:
	//
	// AliyunConsole
	Platform               *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceClassification *string `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// The list of resource group IDs.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The name of the resource group.
	//
	// example:
	//
	// 部门A资源组
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeResourceGroupsRequest) GetAliyunResourceGroupIds() []*string {
	return s.AliyunResourceGroupIds
}

func (s *DescribeResourceGroupsRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeResourceGroupsRequest) GetNeedContainResourceGroupWithOfficeSite() *int64 {
	return s.NeedContainResourceGroupWithOfficeSite
}

func (s *DescribeResourceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResourceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResourceGroupsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeResourceGroupsRequest) GetResourceClassification() *string {
	return s.ResourceClassification
}

func (s *DescribeResourceGroupsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *DescribeResourceGroupsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeResourceGroupsRequest) SetAgentType(v string) *DescribeResourceGroupsRequest {
	s.AgentType = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetAliyunResourceGroupIds(v []*string) *DescribeResourceGroupsRequest {
	s.AliyunResourceGroupIds = v
	return s
}

func (s *DescribeResourceGroupsRequest) SetBusinessChannel(v string) *DescribeResourceGroupsRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetNeedContainResourceGroupWithOfficeSite(v int64) *DescribeResourceGroupsRequest {
	s.NeedContainResourceGroupWithOfficeSite = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageNumber(v int32) *DescribeResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageSize(v int32) *DescribeResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPlatform(v string) *DescribeResourceGroupsRequest {
	s.Platform = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceClassification(v string) *DescribeResourceGroupsRequest {
	s.ResourceClassification = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupIds(v []*string) *DescribeResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupName(v string) *DescribeResourceGroupsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
