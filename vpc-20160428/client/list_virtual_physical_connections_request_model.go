// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualPhysicalConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsConfirmed(v bool) *ListVirtualPhysicalConnectionsRequest
	GetIsConfirmed() *bool
	SetMaxResults(v int32) *ListVirtualPhysicalConnectionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVirtualPhysicalConnectionsRequest
	GetNextToken() *string
	SetPhysicalConnectionId(v string) *ListVirtualPhysicalConnectionsRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *ListVirtualPhysicalConnectionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVirtualPhysicalConnectionsRequest
	GetResourceGroupId() *string
	SetTags(v []*ListVirtualPhysicalConnectionsRequestTags) *ListVirtualPhysicalConnectionsRequest
	GetTags() []*ListVirtualPhysicalConnectionsRequestTags
	SetVirtualPhysicalConnectionAliUids(v []*string) *ListVirtualPhysicalConnectionsRequest
	GetVirtualPhysicalConnectionAliUids() []*string
	SetVirtualPhysicalConnectionBusinessStatus(v string) *ListVirtualPhysicalConnectionsRequest
	GetVirtualPhysicalConnectionBusinessStatus() *string
	SetVirtualPhysicalConnectionIds(v []*string) *ListVirtualPhysicalConnectionsRequest
	GetVirtualPhysicalConnectionIds() []*string
	SetVirtualPhysicalConnectionStatuses(v []*string) *ListVirtualPhysicalConnectionsRequest
	GetVirtualPhysicalConnectionStatuses() []*string
	SetVlanIds(v []*string) *ListVirtualPhysicalConnectionsRequest
	GetVlanIds() []*string
}

type ListVirtualPhysicalConnectionsRequest struct {
	// Indicates whether the shared Express Connect circuits have been confirmed and accepted by the tenant. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsConfirmed *bool `json:"IsConfirmed,omitempty" xml:"IsConfirmed,omitempty"`
	// The number of entries per page when you query results by page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether a next query token is available. Valid values:
	//
	// - You do not need to specify this parameter for the first query or if no next query is available.
	//
	// - If a next query is available, set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the Express Connect circuit associated with the shared Express Connect circuits.
	//
	// In the following content of this topic, the Express Connect circuit associated with the shared Express Connect circuits is referred to as the Express Connect circuit, to distinguish it from the shared Express Connect circuits.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the shared Express Connect circuits.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the shared Express Connect circuits belong.
	//
	// example:
	//
	// rg-acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*ListVirtualPhysicalConnectionsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud account information of the shared Express Connect circuits owner.
	//
	// example:
	//
	// 189xxx
	VirtualPhysicalConnectionAliUids []*string `json:"VirtualPhysicalConnectionAliUids,omitempty" xml:"VirtualPhysicalConnectionAliUids,omitempty" type:"Repeated"`
	// The business status of the shared Express Connect circuits. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: financial lock.
	//
	// - **SecurityLocked**: Locked for security reasons.
	//
	// example:
	//
	// Normal
	VirtualPhysicalConnectionBusinessStatus *string `json:"VirtualPhysicalConnectionBusinessStatus,omitempty" xml:"VirtualPhysicalConnectionBusinessStatus,omitempty"`
	// The shared Express Connect circuits information.
	//
	// example:
	//
	// pc-xxx
	VirtualPhysicalConnectionIds []*string `json:"VirtualPhysicalConnectionIds,omitempty" xml:"VirtualPhysicalConnectionIds,omitempty" type:"Repeated"`
	// The business status information of the shared Express Connect circuits.
	//
	// example:
	//
	// pc-xxx
	VirtualPhysicalConnectionStatuses []*string `json:"VirtualPhysicalConnectionStatuses,omitempty" xml:"VirtualPhysicalConnectionStatuses,omitempty" type:"Repeated"`
	// The VLAN ID of the shared Express Connect circuits.
	//
	// example:
	//
	// pc-xxx
	VlanIds []*string `json:"VlanIds,omitempty" xml:"VlanIds,omitempty" type:"Repeated"`
}

func (s ListVirtualPhysicalConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsRequest) GetIsConfirmed() *bool {
	return s.IsConfirmed
}

func (s *ListVirtualPhysicalConnectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVirtualPhysicalConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualPhysicalConnectionsRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ListVirtualPhysicalConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVirtualPhysicalConnectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVirtualPhysicalConnectionsRequest) GetTags() []*ListVirtualPhysicalConnectionsRequestTags {
	return s.Tags
}

func (s *ListVirtualPhysicalConnectionsRequest) GetVirtualPhysicalConnectionAliUids() []*string {
	return s.VirtualPhysicalConnectionAliUids
}

func (s *ListVirtualPhysicalConnectionsRequest) GetVirtualPhysicalConnectionBusinessStatus() *string {
	return s.VirtualPhysicalConnectionBusinessStatus
}

func (s *ListVirtualPhysicalConnectionsRequest) GetVirtualPhysicalConnectionIds() []*string {
	return s.VirtualPhysicalConnectionIds
}

func (s *ListVirtualPhysicalConnectionsRequest) GetVirtualPhysicalConnectionStatuses() []*string {
	return s.VirtualPhysicalConnectionStatuses
}

func (s *ListVirtualPhysicalConnectionsRequest) GetVlanIds() []*string {
	return s.VlanIds
}

func (s *ListVirtualPhysicalConnectionsRequest) SetIsConfirmed(v bool) *ListVirtualPhysicalConnectionsRequest {
	s.IsConfirmed = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetMaxResults(v int32) *ListVirtualPhysicalConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetNextToken(v string) *ListVirtualPhysicalConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetPhysicalConnectionId(v string) *ListVirtualPhysicalConnectionsRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetRegionId(v string) *ListVirtualPhysicalConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetResourceGroupId(v string) *ListVirtualPhysicalConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetTags(v []*ListVirtualPhysicalConnectionsRequestTags) *ListVirtualPhysicalConnectionsRequest {
	s.Tags = v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetVirtualPhysicalConnectionAliUids(v []*string) *ListVirtualPhysicalConnectionsRequest {
	s.VirtualPhysicalConnectionAliUids = v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetVirtualPhysicalConnectionBusinessStatus(v string) *ListVirtualPhysicalConnectionsRequest {
	s.VirtualPhysicalConnectionBusinessStatus = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetVirtualPhysicalConnectionIds(v []*string) *ListVirtualPhysicalConnectionsRequest {
	s.VirtualPhysicalConnectionIds = v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetVirtualPhysicalConnectionStatuses(v []*string) *ListVirtualPhysicalConnectionsRequest {
	s.VirtualPhysicalConnectionStatuses = v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) SetVlanIds(v []*string) *ListVirtualPhysicalConnectionsRequest {
	s.VlanIds = v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirtualPhysicalConnectionsRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag value cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVirtualPhysicalConnectionsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsRequestTags) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListVirtualPhysicalConnectionsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListVirtualPhysicalConnectionsRequestTags) SetKey(v string) *ListVirtualPhysicalConnectionsRequestTags {
	s.Key = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequestTags) SetValue(v string) *ListVirtualPhysicalConnectionsRequestTags {
	s.Value = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsRequestTags) Validate() error {
	return dara.Validate(s)
}
