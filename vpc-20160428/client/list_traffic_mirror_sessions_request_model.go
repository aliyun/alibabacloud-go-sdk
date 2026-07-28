// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ListTrafficMirrorSessionsRequest
	GetEnabled() *bool
	SetMaxResults(v int32) *ListTrafficMirrorSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrafficMirrorSessionsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTrafficMirrorSessionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTrafficMirrorSessionsRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ListTrafficMirrorSessionsRequest
	GetPriority() *int32
	SetRegionId(v string) *ListTrafficMirrorSessionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTrafficMirrorSessionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListTrafficMirrorSessionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTrafficMirrorSessionsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListTrafficMirrorSessionsRequestTags) *ListTrafficMirrorSessionsRequest
	GetTags() []*ListTrafficMirrorSessionsRequestTags
	SetTrafficMirrorFilterId(v string) *ListTrafficMirrorSessionsRequest
	GetTrafficMirrorFilterId() *string
	SetTrafficMirrorSessionIds(v []*string) *ListTrafficMirrorSessionsRequest
	GetTrafficMirrorSessionIds() []*string
	SetTrafficMirrorSessionName(v string) *ListTrafficMirrorSessionsRequest
	GetTrafficMirrorSessionName() *string
	SetTrafficMirrorSourceId(v string) *ListTrafficMirrorSessionsRequest
	GetTrafficMirrorSourceId() *string
	SetTrafficMirrorTargetId(v string) *ListTrafficMirrorSessionsRequest
	GetTrafficMirrorTargetId() *string
	SetVirtualNetworkId(v int32) *ListTrafficMirrorSessionsRequest
	GetVirtualNetworkId() *int32
}

type ListTrafficMirrorSessionsRequest struct {
	// Specifies whether the traffic mirror session is enabled. Valid values:
	//
	// - **false*	- (default): The traffic mirror session is not enabled.
	//
	// - **true**: The traffic mirror session is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of entries to return in this query. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first request or if no next query exists.
	//
	// - If a next query exists, set the value to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of traffic mirror session. Valid values: **1*	- to **32766**.
	//
	// A smaller value indicates a higher priority. The priority of traffic mirror session created by the same account in the same region must be unique.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the traffic mirror session. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID. For information about the regions that support traffic mirroring, see [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the traffic mirroring session belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tags []*ListTrafficMirrorSessionsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance ID of the traffic mirror filter.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The instance IDs of traffic mirror sessions. The maximum value of **N*	- is **100**, which means you can query up to 100 traffic mirror sessions.
	//
	// example:
	//
	// tms-j6cla50buc44ap8tu****
	TrafficMirrorSessionIds []*string `json:"TrafficMirrorSessionIds,omitempty" xml:"TrafficMirrorSessionIds,omitempty" type:"Repeated"`
	// The name of the traffic mirror session.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	TrafficMirrorSessionName *string `json:"TrafficMirrorSessionName,omitempty" xml:"TrafficMirrorSessionName,omitempty"`
	// The instance ID of the traffic mirror source. Currently, elastic network interfaces (ENIs) are supported as traffic mirror sources.
	//
	// example:
	//
	// eni-j6c8znm5l1yt4sox*****
	TrafficMirrorSourceId *string `json:"TrafficMirrorSourceId,omitempty" xml:"TrafficMirrorSourceId,omitempty"`
	// The instance ID of the traffic mirror destination. Currently, elastic network interfaces (ENIs) and internal-facing SLB instances are supported as traffic mirror destinations. Elastic network interfaces are also referred to as network interface controllers (NICs).
	//
	// example:
	//
	// eni-j6c2fp57q8rr47rp****
	TrafficMirrorTargetId *string `json:"TrafficMirrorTargetId,omitempty" xml:"TrafficMirrorTargetId,omitempty"`
	// The Virtual Network Identifier (VNI) used to distinguish different mirrored data. Valid values: **0*	- to **16777215**. You can use the VNI to identify mirrored data from different sessions at the traffic mirror destination. You can specify a custom VNI value or let the system randomly assign one. To let the system randomly assign a value, do not specify this parameter.
	//
	// example:
	//
	// 10
	VirtualNetworkId *int32 `json:"VirtualNetworkId,omitempty" xml:"VirtualNetworkId,omitempty"`
}

func (s ListTrafficMirrorSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListTrafficMirrorSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrafficMirrorSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMirrorSessionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTrafficMirrorSessionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTrafficMirrorSessionsRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTrafficMirrorSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTrafficMirrorSessionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTrafficMirrorSessionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTrafficMirrorSessionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTrafficMirrorSessionsRequest) GetTags() []*ListTrafficMirrorSessionsRequestTags {
	return s.Tags
}

func (s *ListTrafficMirrorSessionsRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *ListTrafficMirrorSessionsRequest) GetTrafficMirrorSessionIds() []*string {
	return s.TrafficMirrorSessionIds
}

func (s *ListTrafficMirrorSessionsRequest) GetTrafficMirrorSessionName() *string {
	return s.TrafficMirrorSessionName
}

func (s *ListTrafficMirrorSessionsRequest) GetTrafficMirrorSourceId() *string {
	return s.TrafficMirrorSourceId
}

func (s *ListTrafficMirrorSessionsRequest) GetTrafficMirrorTargetId() *string {
	return s.TrafficMirrorTargetId
}

func (s *ListTrafficMirrorSessionsRequest) GetVirtualNetworkId() *int32 {
	return s.VirtualNetworkId
}

func (s *ListTrafficMirrorSessionsRequest) SetEnabled(v bool) *ListTrafficMirrorSessionsRequest {
	s.Enabled = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetMaxResults(v int32) *ListTrafficMirrorSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetNextToken(v string) *ListTrafficMirrorSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetOwnerAccount(v string) *ListTrafficMirrorSessionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetOwnerId(v int64) *ListTrafficMirrorSessionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetPriority(v int32) *ListTrafficMirrorSessionsRequest {
	s.Priority = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetRegionId(v string) *ListTrafficMirrorSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetResourceGroupId(v string) *ListTrafficMirrorSessionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetResourceOwnerAccount(v string) *ListTrafficMirrorSessionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetResourceOwnerId(v int64) *ListTrafficMirrorSessionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTags(v []*ListTrafficMirrorSessionsRequestTags) *ListTrafficMirrorSessionsRequest {
	s.Tags = v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTrafficMirrorFilterId(v string) *ListTrafficMirrorSessionsRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTrafficMirrorSessionIds(v []*string) *ListTrafficMirrorSessionsRequest {
	s.TrafficMirrorSessionIds = v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTrafficMirrorSessionName(v string) *ListTrafficMirrorSessionsRequest {
	s.TrafficMirrorSessionName = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTrafficMirrorSourceId(v string) *ListTrafficMirrorSessionsRequest {
	s.TrafficMirrorSourceId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetTrafficMirrorTargetId(v string) *ListTrafficMirrorSessionsRequest {
	s.TrafficMirrorTargetId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) SetVirtualNetworkId(v int32) *ListTrafficMirrorSessionsRequest {
	s.VirtualNetworkId = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequest) Validate() error {
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

type ListTrafficMirrorSessionsRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficMirrorSessionsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorSessionsRequestTags) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorSessionsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTrafficMirrorSessionsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTrafficMirrorSessionsRequestTags) SetKey(v string) *ListTrafficMirrorSessionsRequestTags {
	s.Key = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequestTags) SetValue(v string) *ListTrafficMirrorSessionsRequestTags {
	s.Value = &v
	return s
}

func (s *ListTrafficMirrorSessionsRequestTags) Validate() error {
	return dara.Validate(s)
}
