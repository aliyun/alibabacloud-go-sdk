// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeFlowLogsRequest
	GetDescription() *string
	SetFlowLogId(v string) *DescribeFlowLogsRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *DescribeFlowLogsRequest
	GetFlowLogName() *string
	SetLogStoreName(v string) *DescribeFlowLogsRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *DescribeFlowLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFlowLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeFlowLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowLogsRequest
	GetPageSize() *int32
	SetProjectName(v string) *DescribeFlowLogsRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeFlowLogsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeFlowLogsRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *DescribeFlowLogsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeFlowLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeFlowLogsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeFlowLogsRequest
	GetResourceType() *string
	SetStatus(v string) *DescribeFlowLogsRequest
	GetStatus() *string
	SetTags(v []*DescribeFlowLogsRequestTags) *DescribeFlowLogsRequest
	GetTags() []*DescribeFlowLogsRequestTags
	SetTrafficType(v string) *DescribeFlowLogsRequest
	GetTrafficType() *string
	SetVpcId(v string) *DescribeFlowLogsRequest
	GetVpcId() *string
}

type DescribeFlowLogsRequest struct {
	// The description of the flow log.
	//
	// The description must be 1 to 256 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Flowlog.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the flow log.
	//
	// example:
	//
	// fl-bp1f6qqhsrc2c12ta****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log.
	//
	// The name must be 1 to 128 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The Logstore that stores the captured traffic.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number, with a default value of **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page in a paginated query, with a maximum value of **50*	- and a default value of **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Project that manages the captured traffic.
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the flow log.
	//
	// You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the flow log.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID of the traffic to capture.
	//
	// example:
	//
	// eni-askldfas****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type of the traffic to capture. Values:
	//
	// - **NetworkInterface**: Elastic Network Interface (ENI).
	//
	// - **VSwitch**: All ENIs within a VSwitch.
	//
	// - **VPC**: All ENIs within a VPC.
	//
	// example:
	//
	// NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the flow log. Values:
	//
	// - **Active**: The flow log is in an active state.
	//
	// - **Activating**: The flow log is being created.
	//
	// - **Inactive**: The flow log is in an inactive state.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*DescribeFlowLogsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of traffic to collect. Values:
	//
	// - **All**: All traffic.
	//
	// - **Allow**: Traffic allowed by access control.
	//
	// - **Drop**: Traffic denied by access control.
	//
	// example:
	//
	// All
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
	// The ID of the VPC for which you want to view the flow log.
	//
	// example:
	//
	// vpc-bp1nwd16gvo1wgs****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeFlowLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogsRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowLogsRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFlowLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFlowLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowLogsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowLogsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFlowLogsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeFlowLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFlowLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeFlowLogsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeFlowLogsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowLogsRequest) GetTags() []*DescribeFlowLogsRequestTags {
	return s.Tags
}

func (s *DescribeFlowLogsRequest) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribeFlowLogsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFlowLogsRequest) SetDescription(v string) *DescribeFlowLogsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogId(v string) *DescribeFlowLogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogName(v string) *DescribeFlowLogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetLogStoreName(v string) *DescribeFlowLogsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetOwnerAccount(v string) *DescribeFlowLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetOwnerId(v int64) *DescribeFlowLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetPageNumber(v int32) *DescribeFlowLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetPageSize(v int32) *DescribeFlowLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetProjectName(v string) *DescribeFlowLogsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetRegionId(v string) *DescribeFlowLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceGroupId(v string) *DescribeFlowLogsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceId(v string) *DescribeFlowLogsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceOwnerAccount(v string) *DescribeFlowLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceOwnerId(v int64) *DescribeFlowLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceType(v string) *DescribeFlowLogsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetStatus(v string) *DescribeFlowLogsRequest {
	s.Status = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetTags(v []*DescribeFlowLogsRequestTags) *DescribeFlowLogsRequest {
	s.Tags = v
	return s
}

func (s *DescribeFlowLogsRequest) SetTrafficType(v string) *DescribeFlowLogsRequest {
	s.TrafficType = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetVpcId(v string) *DescribeFlowLogsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFlowLogsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsRequestTags struct {
	// The key of the tag. Up to 20 tag keys are supported. If you need to pass this value, it cannot be an empty string.
	//
	// A tag key can have up to 128 characters and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Up to 20 tag values are supported. If you need to pass this value, it can be an empty string.
	//
	// A tag value can have up to 128 characters and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowLogsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowLogsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowLogsRequestTags) SetKey(v string) *DescribeFlowLogsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeFlowLogsRequestTags) SetValue(v string) *DescribeFlowLogsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeFlowLogsRequestTags) Validate() error {
	return dara.Validate(s)
}
