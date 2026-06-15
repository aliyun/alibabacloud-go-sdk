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
	// The IDs of dedicated hosts. You can specify up to 100 dedicated host IDs in a JSON array.
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
	// The dedicated host type. Call the [`DescribeDedicatedHostTypes`](https://help.aliyun.com/document_detail/134240.html) operation to get the latest list of dedicated host types.
	//
	// example:
	//
	// ddh.g5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// The reason that the dedicated host is locked. Valid values:
	//
	// - `financial`: The dedicated host is locked due to an overdue payment.
	//
	// - `security`: The dedicated host is locked for security reasons.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maximum number of results to return per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Do not set this parameter for the first request. For subsequent requests, set this parameter to the `NextToken` value returned from the previous response.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter is deprecated. Use `NextToken` and `MaxResults` for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter is deprecated. Use `NextToken` and `MaxResults` for pagination.
	//
	// example:
	//
	// 10
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryInventory *bool  `json:"QueryInventory,omitempty" xml:"QueryInventory,omitempty"`
	// The ID of the region where the dedicated host resides. Call the [`DescribeRegions`](https://help.aliyun.com/document_detail/25609.html) operation to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host belongs. When you use this parameter to filter resources, the number of resources cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-aek3b6jzp66****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to return socket-level capacity information. You can use the information to check the remaining vCPU and memory resources and determine whether an ECS instance of a specific instance type can be created on the dedicated host. Valid values:
	//
	// - `true`: returns the information. Only specific dedicated host types support this feature. For more information, see [View and export information about dedicated hosts](https://help.aliyun.com/document_detail/68989.html).
	//
	// - `false`: does not return the information.
	//
	// 	Notice:
	//
	// A dedicated host typically has two CPUs, which correspond to Socket 0 and Socket 1. To maximize performance, an ECS instance created on a dedicated host is allocated to a single socket and does not span sockets.
	//
	// - If the remaining resources on a socket are sufficient for the specified ECS instance type, the instance can be created.
	//
	// - If the remaining resources on each socket are insufficient for the specified ECS instance type, the instance cannot be created, even if the total remaining resources on both sockets are sufficient.
	//
	// example:
	//
	// true
	SocketDetails *string `json:"SocketDetails,omitempty" xml:"SocketDetails,omitempty"`
	// The state of the dedicated host. Valid values:
	//
	// - `Available`: The dedicated host is running as expected.
	//
	// - `UnderAssessment`: The dedicated host is being assessed for physical hardware risks. The host is available but may have hardware issues that could affect its ECS instances.
	//
	// - `PermanentFailure`: The dedicated host has a permanent failure and is unavailable.
	//
	// - `TempUnavailable`: The dedicated host is temporarily unavailable.
	//
	// - `Redeploying`: The dedicated host is being redeployed.
	//
	// The default value is `Available`.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags used to filter dedicated hosts. You can specify up to 20 tags.
	Tag []*DescribeDedicatedHostsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID. Call the [`DescribeZones`](https://help.aliyun.com/document_detail/25610.html) operation to get the latest list of Alibaba Cloud zones.
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
	// The tag key. The key can be up to 128 characters long. It cannot be an empty string, start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 128 characters long and cannot contain `http://` or `https://`. You can leave the value empty.
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
