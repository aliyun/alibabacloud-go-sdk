// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterId(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostIds(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostIds() *string
	SetDedicatedHostName(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostName() *string
	SetDedicatedHostType(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostType() *string
	SetLockReason(v string) *DescribeDedicatedHostsRequest
	GetLockReason() *string
	SetMaxResults(v int32) *DescribeDedicatedHostsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDedicatedHostsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDedicatedHostsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDedicatedHostsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedHostsRequest
	GetPageSize() *int32
	SetQueryInventory(v bool) *DescribeDedicatedHostsRequest
	GetQueryInventory() *bool
	SetRegionId(v string) *DescribeDedicatedHostsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedHostsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetResourceOwnerId() *int64
	SetSocketDetails(v string) *DescribeDedicatedHostsRequest
	GetSocketDetails() *string
	SetStatus(v string) *DescribeDedicatedHostsRequest
	GetStatus() *string
	SetTag(v []*DescribeDedicatedHostsRequestTag) *DescribeDedicatedHostsRequest
	GetTag() []*DescribeDedicatedHostsRequestTag
	SetZoneId(v string) *DescribeDedicatedHostsRequest
	GetZoneId() *string
}

type DescribeDedicatedHostsRequest struct {
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The list of dedicated host IDs. You can specify up to 100 IDs, separated by commas (,).
	//
	// example:
	//
	// ["dh-bp165p6xk2tlw61e****", "dh-bp1f9vxmno7emy96****"]
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	// The name of the dedicated host.
	//
	// example:
	//
	// MyDDHTestName
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The type of the dedicated host. You can call [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) to query the most recent list of dedicated host types.
	//
	// example:
	//
	// ddh.g5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// The reason why the dedicated host is locked. Valid values:
	//
	// - financial: The dedicated host is locked due to an overdue payment.
	//
	// - security: The dedicated host is locked for security reasons.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maximum number of entries per page for a paging query. If you set this parameter, it indicates that you are using the MaxResults and NextToken paging method.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is about to go offline. Use NextToken and MaxResults to perform paging query operations.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is about to go offline. Use NextToken and MaxResults to perform paging query operations.
	//
	// example:
	//
	// 10
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryInventory *bool  `json:"QueryInventory,omitempty" xml:"QueryInventory,omitempty"`
	// The region ID of the dedicated host. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host belongs. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-aek3b6jzp66****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to display socket-level capacity information. You can use socket-level capacity information to view remaining resources (vCPU, memory usage, remaining capacity, and total capacity) to determine whether an ECS instance of a specific instance type can be created. Valid values:
	//
	// - true: Display socket-level capacity information. Only specific dedicated host types support displaying socket-level resource information. For more information, see [View and export DDH information](https://help.aliyun.com/document_detail/68989.html).
	//
	// - false: Do not display socket-level capacity information.
	//
	// 	Notice:
	//
	// Each dedicated host typically has two CPUs, numbered Socket 0 and Socket 1. On a dedicated host, ECS instances are not created across sockets to ensure maximum performance. An ECS instance is created based on a single socket only.
	//
	// - If the remaining computing resources of one socket are greater than or equal to the instance type to be created, the ECS instance is created.
	//
	// - If the remaining computing resources of each socket are less than the instance type to be created, the ECS instance fails to be created, even if the combined remaining resources of both sockets exceed the instance type requirements.
	//
	// </notice>
	//
	// example:
	//
	// true
	SocketDetails *string `json:"SocketDetails,omitempty" xml:"SocketDetails,omitempty"`
	// The usage status of the dedicated host. Valid values:
	//
	// - Available: The dedicated host is running normally.
	//
	// - UnderAssessment: The physical machine is at risk. The physical machine is available but may cause issues for ECS instances on the dedicated host.
	//
	// - PermanentFailure: The dedicated host has a permanent failure and is unavailable.
	//
	// - TempUnavailable: The dedicated host is temporarily unavailable.
	//
	// - Redeploying: The dedicated host is being restored.
	//
	// Default value: Available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags. Valid values of N: 0 to 20.
	Tag []*DescribeDedicatedHostsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostsRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDedicatedHostsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDedicatedHostsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedHostsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedHostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedHostsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedHostsRequest) GetQueryInventory() *bool {
	return s.QueryInventory
}

func (s *DescribeDedicatedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostsRequest) GetSocketDetails() *string {
	return s.SocketDetails
}

func (s *DescribeDedicatedHostsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedHostsRequest) GetTag() []*DescribeDedicatedHostsRequestTag {
	return s.Tag
}

func (s *DescribeDedicatedHostsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostClusterId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostIds(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostName(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostType(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetLockReason(v string) *DescribeDedicatedHostsRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetMaxResults(v int32) *DescribeDedicatedHostsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetNextToken(v string) *DescribeDedicatedHostsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageNumber(v int32) *DescribeDedicatedHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageSize(v int32) *DescribeDedicatedHostsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetQueryInventory(v bool) *DescribeDedicatedHostsRequest {
	s.QueryInventory = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetRegionId(v string) *DescribeDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceGroupId(v string) *DescribeDedicatedHostsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetSocketDetails(v string) *DescribeDedicatedHostsRequest {
	s.SocketDetails = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetStatus(v string) *DescribeDedicatedHostsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetTag(v []*DescribeDedicatedHostsRequestTag) *DescribeDedicatedHostsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetZoneId(v string) *DescribeDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) Validate() error {
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

type DescribeDedicatedHostsRequestTag struct {
	// The tag key of the dedicated host. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated host. If you specify this parameter, the value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDedicatedHostsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDedicatedHostsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDedicatedHostsRequestTag) SetKey(v string) *DescribeDedicatedHostsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDedicatedHostsRequestTag) SetValue(v string) *DescribeDedicatedHostsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDedicatedHostsRequestTag) Validate() error {
	return dara.Validate(s)
}
