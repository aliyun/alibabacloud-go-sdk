// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamsResponseBody
	GetCount() *int64
	SetIpams(v []*ListIpamsResponseBodyIpams) *ListIpamsResponseBody
	GetIpams() []*ListIpamsResponseBodyIpams
	SetMaxResults(v int64) *ListIpamsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamsResponseBody
	GetTotalCount() *int64
}

type ListIpamsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAMs.
	Ipams []*ListIpamsResponseBodyIpams `json:"Ipams,omitempty" xml:"Ipams,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 23CA0A0B-B0F5-5495-B355-7D9A9203A46B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamsResponseBody) GetIpams() []*ListIpamsResponseBodyIpams {
	return s.Ipams
}

func (s *ListIpamsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamsResponseBody) SetCount(v int64) *ListIpamsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamsResponseBody) SetIpams(v []*ListIpamsResponseBodyIpams) *ListIpamsResponseBody {
	s.Ipams = v
	return s
}

func (s *ListIpamsResponseBody) SetMaxResults(v int64) *ListIpamsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamsResponseBody) SetNextToken(v string) *ListIpamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamsResponseBody) SetRequestId(v string) *ListIpamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamsResponseBody) SetTotalCount(v int64) *ListIpamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamsResponseBodyIpams struct {
	// The time when the IPAM was created.
	//
	// example:
	//
	// 2022-07-01T02:05:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Default resource discovery association ID.
	//
	// example:
	//
	// ipam-res-disco-assoc-jt5fac8twugdbbgip****
	DefaultResourceDiscoveryAssociationId *string `json:"DefaultResourceDiscoveryAssociationId,omitempty" xml:"DefaultResourceDiscoveryAssociationId,omitempty"`
	// Default resource discovery instance ID.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	DefaultResourceDiscoveryId *string `json:"DefaultResourceDiscoveryId,omitempty" xml:"DefaultResourceDiscoveryId,omitempty"`
	// The description of the IPAM.
	//
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The name of the IPAM.
	//
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The status of the IPAM. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	IpamStatus *string `json:"IpamStatus,omitempty" xml:"IpamStatus,omitempty"`
	// The effective regions of the IPAM.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that owns the IPAM.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The default private scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-okoerbco6unqfr****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// The default public scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-ovb76p1g1m19dr****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// The region ID of the IPAM.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Number of resource discovery associations.
	//
	// example:
	//
	// 1
	ResourceDiscoveryAssociationCount *int32 `json:"ResourceDiscoveryAssociationCount,omitempty" xml:"ResourceDiscoveryAssociationCount,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-aek2dbprgpt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of IPAM scopes. Value: **2 to 5**.
	//
	// example:
	//
	// 2
	ScopeCount *int32 `json:"ScopeCount,omitempty" xml:"ScopeCount,omitempty"`
	// The tag list.
	Tags []*ListIpamsResponseBodyIpamsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamsResponseBodyIpams) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsResponseBodyIpams) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBodyIpams) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIpamsResponseBodyIpams) GetDefaultResourceDiscoveryAssociationId() *string {
	return s.DefaultResourceDiscoveryAssociationId
}

func (s *ListIpamsResponseBodyIpams) GetDefaultResourceDiscoveryId() *string {
	return s.DefaultResourceDiscoveryId
}

func (s *ListIpamsResponseBodyIpams) GetIpamDescription() *string {
	return s.IpamDescription
}

func (s *ListIpamsResponseBodyIpams) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamsResponseBodyIpams) GetIpamName() *string {
	return s.IpamName
}

func (s *ListIpamsResponseBodyIpams) GetIpamStatus() *string {
	return s.IpamStatus
}

func (s *ListIpamsResponseBodyIpams) GetOperatingRegionList() []*string {
	return s.OperatingRegionList
}

func (s *ListIpamsResponseBodyIpams) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamsResponseBodyIpams) GetPrivateDefaultScopeId() *string {
	return s.PrivateDefaultScopeId
}

func (s *ListIpamsResponseBodyIpams) GetPublicDefaultScopeId() *string {
	return s.PublicDefaultScopeId
}

func (s *ListIpamsResponseBodyIpams) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamsResponseBodyIpams) GetResourceDiscoveryAssociationCount() *int32 {
	return s.ResourceDiscoveryAssociationCount
}

func (s *ListIpamsResponseBodyIpams) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamsResponseBodyIpams) GetScopeCount() *int32 {
	return s.ScopeCount
}

func (s *ListIpamsResponseBodyIpams) GetTags() []*ListIpamsResponseBodyIpamsTags {
	return s.Tags
}

func (s *ListIpamsResponseBodyIpams) SetCreateTime(v string) *ListIpamsResponseBodyIpams {
	s.CreateTime = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetDefaultResourceDiscoveryAssociationId(v string) *ListIpamsResponseBodyIpams {
	s.DefaultResourceDiscoveryAssociationId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetDefaultResourceDiscoveryId(v string) *ListIpamsResponseBodyIpams {
	s.DefaultResourceDiscoveryId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamDescription(v string) *ListIpamsResponseBodyIpams {
	s.IpamDescription = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamId(v string) *ListIpamsResponseBodyIpams {
	s.IpamId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamName(v string) *ListIpamsResponseBodyIpams {
	s.IpamName = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetIpamStatus(v string) *ListIpamsResponseBodyIpams {
	s.IpamStatus = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetOperatingRegionList(v []*string) *ListIpamsResponseBodyIpams {
	s.OperatingRegionList = v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetOwnerId(v int64) *ListIpamsResponseBodyIpams {
	s.OwnerId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetPrivateDefaultScopeId(v string) *ListIpamsResponseBodyIpams {
	s.PrivateDefaultScopeId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetPublicDefaultScopeId(v string) *ListIpamsResponseBodyIpams {
	s.PublicDefaultScopeId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetRegionId(v string) *ListIpamsResponseBodyIpams {
	s.RegionId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetResourceDiscoveryAssociationCount(v int32) *ListIpamsResponseBodyIpams {
	s.ResourceDiscoveryAssociationCount = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetResourceGroupId(v string) *ListIpamsResponseBodyIpams {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetScopeCount(v int32) *ListIpamsResponseBodyIpams {
	s.ScopeCount = &v
	return s
}

func (s *ListIpamsResponseBodyIpams) SetTags(v []*ListIpamsResponseBodyIpamsTags) *ListIpamsResponseBodyIpams {
	s.Tags = v
	return s
}

func (s *ListIpamsResponseBodyIpams) Validate() error {
	return dara.Validate(s)
}

type ListIpamsResponseBodyIpamsTags struct {
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

func (s ListIpamsResponseBodyIpamsTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsResponseBodyIpamsTags) GoString() string {
	return s.String()
}

func (s *ListIpamsResponseBodyIpamsTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamsResponseBodyIpamsTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamsResponseBodyIpamsTags) SetKey(v string) *ListIpamsResponseBodyIpamsTags {
	s.Key = &v
	return s
}

func (s *ListIpamsResponseBodyIpamsTags) SetValue(v string) *ListIpamsResponseBodyIpamsTags {
	s.Value = &v
	return s
}

func (s *ListIpamsResponseBodyIpamsTags) Validate() error {
	return dara.Validate(s)
}
