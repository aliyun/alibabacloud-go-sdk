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
	SetQosRuleId(v string) *DescribeDesktopGroupsRequest
	GetQosRuleId() *string
	SetRegionId(v string) *DescribeDesktopGroupsRequest
	GetRegionId() *string
	SetStatus(v int32) *DescribeDesktopGroupsRequest
	GetStatus() *int32
	SetTag(v []*DescribeDesktopGroupsRequestTag) *DescribeDesktopGroupsRequest
	GetTag() []*DescribeDesktopGroupsRequestTag
}

type DescribeDesktopGroupsRequest struct {
	// The cloud computer template IDs.
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	// The ID of the cloud computer pool.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The IDs of cloud computer pools.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The name of the cloud computer pool. Fuzzy search is supported.
	//
	// example:
	//
	// CloudComputerPool01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DesktopType      *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The IDs of the authorized users of the cloud computer pool.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The IDs of the users that you want to exclude from the authorized user list.
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	// The image IDs.
	//
	// if can be null:
	// false
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The number of entries to return on each page.<br>Maximum value: 100.<br>Default value: 10.<br><br>
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether to query multi-desktop cloud computer pools.
	//
	// example:
	//
	// true
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the cloud computer pool.
	//
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	OwnType *int64 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration of the subscription cloud computer pool. The unit is specified by the `PeriodUnit` parameter.
	//
	// - Valid values when `PeriodUnit` is set to `Month`:
	//
	//   - 1
	//
	//   - 2
	//
	//   - 3
	//
	//   - 6
	//
	// - Valid values when `PeriodUnit` is set to `Year`:
	//
	//   - 1
	//
	//   - 2
	//
	//   - 3
	//
	//   - 4
	//
	//   - 5
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
	// The ID of the policy that is associated with the cloud computer pool.
	//
	// example:
	//
	// pg-53iyi2aar0nd6****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	QosRuleId    *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cloud computer pool.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags. You can specify up to 20 tags.
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

func (s *DescribeDesktopGroupsRequest) GetQosRuleId() *string {
	return s.QosRuleId
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

func (s *DescribeDesktopGroupsRequest) SetQosRuleId(v string) *DescribeDesktopGroupsRequest {
	s.QosRuleId = &v
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

type DescribeDesktopGroupsRequestTag struct {
	// The key of the tag. The key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
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
