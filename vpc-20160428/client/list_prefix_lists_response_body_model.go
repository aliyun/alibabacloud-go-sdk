// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrefixListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListPrefixListsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListPrefixListsResponseBody
	GetNextToken() *string
	SetPrefixLists(v []*ListPrefixListsResponseBodyPrefixLists) *ListPrefixListsResponseBody
	GetPrefixLists() []*ListPrefixListsResponseBodyPrefixLists
	SetRequestId(v string) *ListPrefixListsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPrefixListsResponseBody
	GetTotalCount() *int64
}

type ListPrefixListsResponseBody struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value indicates the token that is used for the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the prefix lists.
	PrefixLists []*ListPrefixListsResponseBodyPrefixLists `json:"PrefixLists,omitempty" xml:"PrefixLists,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DF72F7BB-5DFA-529C-887E-B0BB70D89C4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPrefixListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrefixListsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPrefixListsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrefixListsResponseBody) GetPrefixLists() []*ListPrefixListsResponseBodyPrefixLists {
	return s.PrefixLists
}

func (s *ListPrefixListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrefixListsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPrefixListsResponseBody) SetMaxResults(v int64) *ListPrefixListsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPrefixListsResponseBody) SetNextToken(v string) *ListPrefixListsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPrefixListsResponseBody) SetPrefixLists(v []*ListPrefixListsResponseBodyPrefixLists) *ListPrefixListsResponseBody {
	s.PrefixLists = v
	return s
}

func (s *ListPrefixListsResponseBody) SetRequestId(v string) *ListPrefixListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrefixListsResponseBody) SetTotalCount(v int64) *ListPrefixListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrefixListsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrefixListsResponseBodyPrefixLists struct {
	// The CIDR block specified in the prefix list.
	CidrBlocks []*string `json:"CidrBlocks,omitempty" xml:"CidrBlocks,omitempty" type:"Repeated"`
	// The time when the prefix list was created.
	//
	// example:
	//
	// 2022-07-12T14:22:32Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IP version of the prefix list. Valid values:
	//
	// 	- **IPV4**
	//
	// 	- **IPV6**
	//
	// example:
	//
	// IPV4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The maximum number of CIDR blocks that you can specify in the prefix list.
	//
	// example:
	//
	// 10
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The Alibaba Cloud account to which the prefix list belongs.
	//
	// example:
	//
	// 1210123456123456
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the prefix list.
	//
	// example:
	//
	// Created with oss service by system.
	PrefixListDescription *string `json:"PrefixListDescription,omitempty" xml:"PrefixListDescription,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-m5estsqsdqwg88hjf****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list.
	//
	// example:
	//
	// test
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The status of the prefix list. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// 	- **Modifying**
	//
	// >  This parameter is the same as the **Status*	- parameter.
	//
	// example:
	//
	// Created
	PrefixListStatus *string `json:"PrefixListStatus,omitempty" xml:"PrefixListStatus,omitempty"`
	// The type of the prefix list.
	//
	// example:
	//
	// Custom
	PrefixListType *string `json:"PrefixListType,omitempty" xml:"PrefixListType,omitempty"`
	// The region ID of the prefix list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the prefix list is shared. Valid values:
	//
	// 	- **Shared**: The prefix list is shared.
	//
	// 	- If an empty value is returned, the prefix list is not shared.
	//
	// example:
	//
	// Shared
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The status of the prefix list. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListPrefixListsResponseBodyPrefixListsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPrefixListsResponseBodyPrefixLists) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsResponseBodyPrefixLists) GoString() string {
	return s.String()
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetCidrBlocks() []*string {
	return s.CidrBlocks
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetPrefixListDescription() *string {
	return s.PrefixListDescription
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetPrefixListStatus() *string {
	return s.PrefixListStatus
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetPrefixListType() *string {
	return s.PrefixListType
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetShareType() *string {
	return s.ShareType
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetStatus() *string {
	return s.Status
}

func (s *ListPrefixListsResponseBodyPrefixLists) GetTags() []*ListPrefixListsResponseBodyPrefixListsTags {
	return s.Tags
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetCidrBlocks(v []*string) *ListPrefixListsResponseBodyPrefixLists {
	s.CidrBlocks = v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetCreationTime(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.CreationTime = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetIpVersion(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.IpVersion = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetMaxEntries(v int32) *ListPrefixListsResponseBodyPrefixLists {
	s.MaxEntries = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetOwnerId(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.OwnerId = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetPrefixListDescription(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.PrefixListDescription = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetPrefixListId(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.PrefixListId = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetPrefixListName(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.PrefixListName = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetPrefixListStatus(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.PrefixListStatus = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetPrefixListType(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.PrefixListType = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetRegionId(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.RegionId = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetResourceGroupId(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetShareType(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.ShareType = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetStatus(v string) *ListPrefixListsResponseBodyPrefixLists {
	s.Status = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) SetTags(v []*ListPrefixListsResponseBodyPrefixListsTags) *ListPrefixListsResponseBodyPrefixLists {
	s.Tags = v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixLists) Validate() error {
	return dara.Validate(s)
}

type ListPrefixListsResponseBodyPrefixListsTags struct {
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

func (s ListPrefixListsResponseBodyPrefixListsTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsResponseBodyPrefixListsTags) GoString() string {
	return s.String()
}

func (s *ListPrefixListsResponseBodyPrefixListsTags) GetKey() *string {
	return s.Key
}

func (s *ListPrefixListsResponseBodyPrefixListsTags) GetValue() *string {
	return s.Value
}

func (s *ListPrefixListsResponseBodyPrefixListsTags) SetKey(v string) *ListPrefixListsResponseBodyPrefixListsTags {
	s.Key = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixListsTags) SetValue(v string) *ListPrefixListsResponseBodyPrefixListsTags {
	s.Value = &v
	return s
}

func (s *ListPrefixListsResponseBodyPrefixListsTags) Validate() error {
	return dara.Validate(s)
}
