// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v []*string) *DescribeDesktopGroupsRequest
	GetBundleId() []*string
	SetDesktopGroupId(v string) *DescribeDesktopGroupsRequest
	GetDesktopGroupId() *string
	SetDesktopGroupIds(v []*string) *DescribeDesktopGroupsRequest
	GetDesktopGroupIds() []*string
	SetDesktopGroupName(v string) *DescribeDesktopGroupsRequest
	GetDesktopGroupName() *string
	SetDesktopType(v string) *DescribeDesktopGroupsRequest
	GetDesktopType() *string
	SetEndUserIds(v []*string) *DescribeDesktopGroupsRequest
	GetEndUserIds() []*string
	SetExcludedEndUserIds(v []*string) *DescribeDesktopGroupsRequest
	GetExcludedEndUserIds() []*string
	SetImageId(v []*string) *DescribeDesktopGroupsRequest
	GetImageId() []*string
	SetMaxResults(v int32) *DescribeDesktopGroupsRequest
	GetMaxResults() *int32
	SetMultiResource(v bool) *DescribeDesktopGroupsRequest
	GetMultiResource() *bool
	SetNextToken(v string) *DescribeDesktopGroupsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeDesktopGroupsRequest
	GetOfficeSiteId() *string
	SetOwnType(v int64) *DescribeDesktopGroupsRequest
	GetOwnType() *int64
	SetPeriod(v int32) *DescribeDesktopGroupsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeDesktopGroupsRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *DescribeDesktopGroupsRequest
	GetPolicyGroupId() *string
	SetProtocolType(v string) *DescribeDesktopGroupsRequest
	GetProtocolType() *string
	SetRegionId(v string) *DescribeDesktopGroupsRequest
	GetRegionId() *string
	SetStatus(v int32) *DescribeDesktopGroupsRequest
	GetStatus() *int32
	SetTag(v []*DescribeDesktopGroupsRequestTag) *DescribeDesktopGroupsRequest
	GetTag() []*DescribeDesktopGroupsRequestTag
}

type DescribeDesktopGroupsRequest struct {
	// The IDs of the cloud computer templates.
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	// The ID of the cloud computer share.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The IDs of the cloud computer shares.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The name of the cloud computer share that you want to query. Fuzzy search is supported.
	//
	// example:
	//
	// testName
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DesktopType      *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The IDs of the users who can access the cloud computer share.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The authorized users that you want to exclude.
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	// The IDs of the images.
	//
	// if can be null:
	// false
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether the cloud computer share is a many-to-many share.
	//
	// Valid values:
	//
	// 	- true: The cloud computer share is a many-to-many share.
	//
	// 	- false: The cloud computer share is a one-to-many share.
	//
	// example:
	//
	// true
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the NextToken parameter is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network in which the cloud computer share resides.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the cloud computer share.
	//
	// >  This parameter is not publicly available.
	//
	// Valid values:
	//
	// 	- 0: a single-session many-to-many share.
	//
	// 	- 1: a multi-session many-to-many share.
	//
	// example:
	//
	// 0
	OwnType *int64 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration of the cloud computer share. The unit is specified by `PeriodUnit`.
	//
	// 	- Valid values if you set `PeriodUnit` to `Month`:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 6
	//
	// 	- Valid values if you set `PeriodUnit` to `Year`:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 4
	//
	//     	- 5
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the applied policy.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- High-definition Experience (HDX)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Adaptive Streaming Protocol (ASP)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cloud computer share.
	//
	// Valid values:
	//
	// 	- 0: The cloud computer share is unpaid.
	//
	// 	- 1: The cloud computer share is normal.
	//
	// 	- 2: The cloud computer share expired, or your account has an overdue payment.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that you want to add to the cloud computer share. You can specify 1 to 20 tags.
	Tag []*DescribeDesktopGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDesktopGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsRequest) GetBundleId() []*string {
	return s.BundleId
}

func (s *DescribeDesktopGroupsRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopGroupsRequest) GetDesktopGroupIds() []*string {
	return s.DesktopGroupIds
}

func (s *DescribeDesktopGroupsRequest) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeDesktopGroupsRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopGroupsRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeDesktopGroupsRequest) GetExcludedEndUserIds() []*string {
	return s.ExcludedEndUserIds
}

func (s *DescribeDesktopGroupsRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *DescribeDesktopGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopGroupsRequest) GetMultiResource() *bool {
	return s.MultiResource
}

func (s *DescribeDesktopGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopGroupsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopGroupsRequest) GetOwnType() *int64 {
	return s.OwnType
}

func (s *DescribeDesktopGroupsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeDesktopGroupsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeDesktopGroupsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopGroupsRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopGroupsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDesktopGroupsRequest) GetTag() []*DescribeDesktopGroupsRequestTag {
	return s.Tag
}

func (s *DescribeDesktopGroupsRequest) SetBundleId(v []*string) *DescribeDesktopGroupsRequest {
	s.BundleId = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupId(v string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupIds(v []*string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupIds = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopGroupName(v string) *DescribeDesktopGroupsRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetDesktopType(v string) *DescribeDesktopGroupsRequest {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetEndUserIds(v []*string) *DescribeDesktopGroupsRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetExcludedEndUserIds(v []*string) *DescribeDesktopGroupsRequest {
	s.ExcludedEndUserIds = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetImageId(v []*string) *DescribeDesktopGroupsRequest {
	s.ImageId = v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetMaxResults(v int32) *DescribeDesktopGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetMultiResource(v bool) *DescribeDesktopGroupsRequest {
	s.MultiResource = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetNextToken(v string) *DescribeDesktopGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetOfficeSiteId(v string) *DescribeDesktopGroupsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetOwnType(v int64) *DescribeDesktopGroupsRequest {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPeriod(v int32) *DescribeDesktopGroupsRequest {
	s.Period = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPeriodUnit(v string) *DescribeDesktopGroupsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetPolicyGroupId(v string) *DescribeDesktopGroupsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetProtocolType(v string) *DescribeDesktopGroupsRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetRegionId(v string) *DescribeDesktopGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetStatus(v int32) *DescribeDesktopGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDesktopGroupsRequest) SetTag(v []*DescribeDesktopGroupsRequestTag) *DescribeDesktopGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDesktopGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopGroupsRequestTag struct {
	// The tag key. You cannot specify an empty string as a tag key. A tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify an empty string as a tag key. A tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDesktopGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDesktopGroupsRequestTag) SetKey(v string) *DescribeDesktopGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDesktopGroupsRequestTag) SetValue(v string) *DescribeDesktopGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDesktopGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
