// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationInterval(v int32) *CreateFlowLogRequest
	GetAggregationInterval() *int32
	SetDescription(v string) *CreateFlowLogRequest
	GetDescription() *string
	SetFlowLogName(v string) *CreateFlowLogRequest
	GetFlowLogName() *string
	SetIpVersion(v string) *CreateFlowLogRequest
	GetIpVersion() *string
	SetLogStoreName(v string) *CreateFlowLogRequest
	GetLogStoreName() *string
	SetOwnerAccount(v string) *CreateFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFlowLogRequest
	GetOwnerId() *int64
	SetProjectName(v string) *CreateFlowLogRequest
	GetProjectName() *string
	SetRegionId(v string) *CreateFlowLogRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateFlowLogRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *CreateFlowLogRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *CreateFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowLogRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *CreateFlowLogRequest
	GetResourceType() *string
	SetTag(v []*CreateFlowLogRequestTag) *CreateFlowLogRequest
	GetTag() []*CreateFlowLogRequestTag
	SetTrafficPath(v []*string) *CreateFlowLogRequest
	GetTrafficPath() []*string
	SetTrafficType(v string) *CreateFlowLogRequest
	GetTrafficType() *string
}

type CreateFlowLogRequest struct {
	// The sampling interval of the flow log. Unit: seconds. Valid values: **1**, **5**, and **10*	- (default).
	//
	// example:
	//
	// 10
	AggregationInterval *int32 `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	// The description of the flow log.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Flowlog.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the flow log.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	IpVersion   *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The name of the Logstore that stores the captured traffic data.
	//
	// 	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the project that stores the captured traffic data.
	//
	// 	- The name can contain only lowercase letters, digits, and hyphens (-).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where you want to create the flow log. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource whose traffic you want to capture.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-askldfas****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource whose traffic you want to capture. Valid values:
	//
	// 	- **NetworkInterface**: elastic network interface (ENI)
	//
	// 	- **VSwitch**: all ENIs in a vSwitch
	//
	// 	- **VPC**: all ENIs in a virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag of the resource.
	Tag []*CreateFlowLogRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The scope of the traffic that you want to capture. Valid values:
	//
	// 	- **all**: all traffic.
	//
	// 	- **internetGateway**: Internet traffic.
	TrafficPath []*string `json:"TrafficPath,omitempty" xml:"TrafficPath,omitempty" type:"Repeated"`
	// The type of traffic that you want to capture. Valid values:
	//
	// 	- **All**: all traffic
	//
	// 	- **Allow**: traffic that is allowed
	//
	// 	- **Drop**: traffic that is rejected
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s CreateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequest) GetAggregationInterval() *int32 {
	return s.AggregationInterval
}

func (s *CreateFlowLogRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowLogRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *CreateFlowLogRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateFlowLogRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowLogRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFlowLogRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFlowLogRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowLogRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateFlowLogRequest) GetTag() []*CreateFlowLogRequestTag {
	return s.Tag
}

func (s *CreateFlowLogRequest) GetTrafficPath() []*string {
	return s.TrafficPath
}

func (s *CreateFlowLogRequest) GetTrafficType() *string {
	return s.TrafficType
}

func (s *CreateFlowLogRequest) SetAggregationInterval(v int32) *CreateFlowLogRequest {
	s.AggregationInterval = &v
	return s
}

func (s *CreateFlowLogRequest) SetDescription(v string) *CreateFlowLogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowLogRequest) SetFlowLogName(v string) *CreateFlowLogRequest {
	s.FlowLogName = &v
	return s
}

func (s *CreateFlowLogRequest) SetIpVersion(v string) *CreateFlowLogRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateFlowLogRequest) SetLogStoreName(v string) *CreateFlowLogRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateFlowLogRequest) SetOwnerAccount(v string) *CreateFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFlowLogRequest) SetOwnerId(v int64) *CreateFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowLogRequest) SetProjectName(v string) *CreateFlowLogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowLogRequest) SetRegionId(v string) *CreateFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceGroupId(v string) *CreateFlowLogRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceId(v string) *CreateFlowLogRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceOwnerAccount(v string) *CreateFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceOwnerId(v int64) *CreateFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceType(v string) *CreateFlowLogRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateFlowLogRequest) SetTag(v []*CreateFlowLogRequestTag) *CreateFlowLogRequest {
	s.Tag = v
	return s
}

func (s *CreateFlowLogRequest) SetTrafficPath(v []*string) *CreateFlowLogRequest {
	s.TrafficPath = v
	return s
}

func (s *CreateFlowLogRequest) SetTrafficType(v string) *CreateFlowLogRequest {
	s.TrafficType = &v
	return s
}

func (s *CreateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}

type CreateFlowLogRequestTag struct {
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFlowLogRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFlowLogRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFlowLogRequestTag) SetKey(v string) *CreateFlowLogRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFlowLogRequestTag) SetValue(v string) *CreateFlowLogRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFlowLogRequestTag) Validate() error {
	return dara.Validate(s)
}
