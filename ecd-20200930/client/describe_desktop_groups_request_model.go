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
	// The list of cloud computer template IDs.
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	// The ID of the shared cloud computer.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The list of shared cloud computer IDs.
	DesktopGroupIds []*string `json:"DesktopGroupIds,omitempty" xml:"DesktopGroupIds,omitempty" type:"Repeated"`
	// The name of the shared cloud computer to query. Fuzzy match is supported.
	//
	// example:
	//
	// CloudComputerPool01
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The cloud computer specifications. You can call [DescribeDesktopTypes](~~DescribeDesktopTypes~~) to query the supported specification IDs.
	//
	// example:
	//
	// eds.enterprise_office.16c64g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The list of authorized user IDs for the shared cloud computer.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The list of authorized users to exclude.
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	// The list of image IDs.
	//
	// if can be null:
	// false
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether the shared cloud computer is a multi-host type.
	//
	// Valid values:
	//
	// - true: Multi-host shared cloud computer.
	//
	// - false: Single-host shared cloud computer.
	//
	// example:
	//
	// true
	MultiResource *bool `json:"MultiResource,omitempty" xml:"MultiResource,omitempty"`
	// The token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network to which the shared cloud computers belong.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the shared cloud computer.
	//
	// example:
	//
	// 0
	OwnType *int64 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The subscription duration of the shared cloud computer. The unit is specified by `PeriodUnit`.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the duration for the subscription billing method.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the policy associated with the shared cloud computer.
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
	// The ID of the QoS rule.
	//
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the shared cloud computer.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags. You can specify 1 to 20 tags.
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
	// The tag key. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
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
