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
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
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
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project that manages the captured traffic data.
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the flow log.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the flow log belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource from which traffic is captured.
	//
	// example:
	//
	// eni-askldfas****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource from which traffic is captured. Valid values:
	//
	// 	- **NetworkInterface**: elastic network interface (ENI)
	//
	// 	- **VSwitch**: all ENIs in a vSwitch
	//
	// 	- **VPC**: all ENIs in a virtual private cloud (VPC)
	//
	// example:
	//
	// NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the flow log. Valid values:
	//
	// 	- **Active**
	//
	// 	- **Activating**
	//
	// 	- **Inactive**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*DescribeFlowLogsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of traffic that is captured. Valid values:
	//
	// 	- **All**: all traffic
	//
	// 	- **Allow**: traffic that is allowed by access control
	//
	// 	- **Drop**: traffic that is denied by access control
	//
	// example:
	//
	// All
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
	// The ID of the VPC to which the flow log belongs.
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
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
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
