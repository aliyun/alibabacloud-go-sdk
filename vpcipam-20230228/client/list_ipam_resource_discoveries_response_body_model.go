// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListIpamResourceDiscoveriesResponseBody
	GetCount() *int32
	SetIpamResourceDiscoveries(v []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) *ListIpamResourceDiscoveriesResponseBody
	GetIpamResourceDiscoveries() []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries
	SetMaxResults(v int32) *ListIpamResourceDiscoveriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamResourceDiscoveriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamResourceDiscoveriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamResourceDiscoveriesResponseBody
	GetTotalCount() *int64
}

type ListIpamResourceDiscoveriesResponseBody struct {
	// The maximum number of entries on each page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resource discovery instances.
	IpamResourceDiscoveries []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries `json:"IpamResourceDiscoveries,omitempty" xml:"IpamResourceDiscoveries,omitempty" type:"Repeated"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 86137597-443F-5B66-B9B6-8514E0C50B8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetIpamResourceDiscoveries() []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	return s.IpamResourceDiscoveries
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamResourceDiscoveriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetCount(v int32) *ListIpamResourceDiscoveriesResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetIpamResourceDiscoveries(v []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) *ListIpamResourceDiscoveriesResponseBody {
	s.IpamResourceDiscoveries = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetMaxResults(v int32) *ListIpamResourceDiscoveriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetNextToken(v string) *ListIpamResourceDiscoveriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetRequestId(v string) *ListIpamResourceDiscoveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) SetTotalCount(v int64) *ListIpamResourceDiscoveriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBody) Validate() error {
	if s.IpamResourceDiscoveries != nil {
		for _, item := range s.IpamResourceDiscoveries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries struct {
	// The time when the resource was discovered.
	//
	// example:
	//
	// 2022-07-01T02:05:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource discovery.
	//
	// example:
	//
	// test description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// The status of the resource discovery instance. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	IpamResourceDiscoveryStatus *string `json:"IpamResourceDiscoveryStatus,omitempty" xml:"IpamResourceDiscoveryStatus,omitempty"`
	// The list of resource discovery regions.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that owns the resource discovery.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the queried resource discovery instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group that resource discovery belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing status of the resource.
	//
	// 	- If the value is empty, the resource is as an average instance.
	//
	// 	- If the value is Shared, the resource discovery comes from a shared source.
	//
	// 	- If the value is Sharing, the resource discovery is being shared.
	//
	// example:
	//
	// Shared
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag list.
	Tags []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of resource discovery.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetIpamResourceDiscoveryDescription() *string {
	return s.IpamResourceDiscoveryDescription
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetIpamResourceDiscoveryName() *string {
	return s.IpamResourceDiscoveryName
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetIpamResourceDiscoveryStatus() *string {
	return s.IpamResourceDiscoveryStatus
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetOperatingRegionList() []*string {
	return s.OperatingRegionList
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetShareType() *string {
	return s.ShareType
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetTags() []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags {
	return s.Tags
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) GetType() *string {
	return s.Type
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetCreateTime(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.CreateTime = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryDescription(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryName(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetIpamResourceDiscoveryStatus(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.IpamResourceDiscoveryStatus = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetOperatingRegionList(v []*string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.OperatingRegionList = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetOwnerId(v int64) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetRegionId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetResourceGroupId(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetShareType(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.ShareType = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetTags(v []*ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.Tags = v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) SetType(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries {
	s.Type = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveries) Validate() error {
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

type ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) SetKey(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags {
	s.Key = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) SetValue(v string) *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags {
	s.Value = &v
	return s
}

func (s *ListIpamResourceDiscoveriesResponseBodyIpamResourceDiscoveriesTags) Validate() error {
	return dara.Validate(s)
}
