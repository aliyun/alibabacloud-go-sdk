// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainList(v []*DescribeOutgoingDomainResponseBodyDomainList) *DescribeOutgoingDomainResponseBody
	GetDomainList() []*DescribeOutgoingDomainResponseBodyDomainList
	SetRequestId(v string) *DescribeOutgoingDomainResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDomainResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingDomainResponseBody struct {
	// The domain names in outbound connections.
	DomainList []*DescribeOutgoingDomainResponseBodyDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the domain names in outbound connections.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBody) GetDomainList() []*DescribeOutgoingDomainResponseBodyDomainList {
	return s.DomainList
}

func (s *DescribeOutgoingDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDomainResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDomainResponseBody) SetDomainList(v []*DescribeOutgoingDomainResponseBodyDomainList) *DescribeOutgoingDomainResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetRequestId(v string) *DescribeOutgoingDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetTotalCount(v int32) *DescribeOutgoingDomainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOutgoingDomainResponseBodyDomainList struct {
	// Indicates whether an access control policy is configured. Valid values:
	//
	// 	- **Uncovered**: no
	//
	// 	- **FullCoverage**: yes
	//
	// example:
	//
	// Uncovered
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The suggestion in an access control policy.
	//
	// example:
	//
	// Allows the traffic.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The state of the access control policy. Valid values:
	//
	// 	- **normal**: healthy
	//
	// 	- **abnormal**: unhealthy
	//
	// example:
	//
	// Normal
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// The address book for outbound connections
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	//
	// example:
	//
	// fdad-fdafa-dafa-dfa****
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The outbound asset count.
	//
	// example:
	//
	// 20
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The website service.
	//
	// example:
	//
	// Alibaba Cloud
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// The type of the tag. Valid values:
	//
	// 	- **Suspicious**
	//
	// 	- **Malicious**
	//
	// 	- **Trusted**
	//
	// example:
	//
	// Trusted
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The type ID of the service to which the domain name belongs. Valid values:
	//
	// 	- **Aliyun**: Alibaba Cloud services
	//
	// 	- **NotAliyun**: third-party services
	//
	// example:
	//
	// Aliyun
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the service to which the domain name belongs. Valid values:
	//
	// 	- **Alibaba Cloud services**
	//
	// 	- **Third-party services**
	//
	// example:
	//
	// Alibaba Cloud services
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The domain name in outbound connections.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the group to which the access control policy belongs.
	//
	// example:
	//
	// Group of addresses in outbound connections
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an `access control policy` is configured for the domain name. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an access control policy is recommended. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The volume of inbound traffic.
	//
	// example:
	//
	// 3214
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the domain name is marked as normal. Valid values:
	//
	// 	- **true**: normal
	//
	// 	- **false**: abnormal
	//
	// example:
	//
	// true
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The volume of outbound traffic.
	//
	// example:
	//
	// 4582
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The outbound private asset count.
	//
	// example:
	//
	// 20
	PrivateAssetCount *int64 `json:"PrivateAssetCount,omitempty" xml:"PrivateAssetCount,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// add-dfadf-f****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// Default rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason why the domain name is secure.
	//
	// example:
	//
	// Intelligent policy: The destination domain name belongs to Alibaba Cloud Computing Co., Ltd. The domain name mainly provides services for Alibaba Cloud. No security risks are found, and you can add the domain name to the whitelist.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The suggestion to handle the traffic of the domain name in outbound connections. Valid values:
	//
	// 	- **pass**: allow
	//
	// 	- **alert**: monitor
	//
	// 	- **drop**: deny
	//
	// example:
	//
	// pass
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 12
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// An array that consists of tags.
	TagList []*DescribeOutgoingDomainResponseBodyDomainListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total volume of traffic. Unit: bytes.
	//
	// example:
	//
	// 800
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAclRecommendDetail() *string {
	return s.AclRecommendDetail
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAddressGroupName() *string {
	return s.AddressGroupName
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAddressGroupUUID() *string {
	return s.AddressGroupUUID
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetAssetCount() *int64 {
	return s.AssetCount
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetBusiness() *string {
	return s.Business
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetCategoryClassId() *string {
	return s.CategoryClassId
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetHasAcl() *string {
	return s.HasAcl
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetHasAclRecommend() *bool {
	return s.HasAclRecommend
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetIsMarkNormal() *bool {
	return s.IsMarkNormal
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetOrganization() *string {
	return s.Organization
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetPrivateAssetCount() *int64 {
	return s.PrivateAssetCount
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetSecurityReason() *string {
	return s.SecurityReason
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetSecuritySuggest() *string {
	return s.SecuritySuggest
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetTagList() []*DescribeOutgoingDomainResponseBodyDomainListTagList {
	return s.TagList
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclCoverage(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclRecommendDetail(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclStatus(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupUUID(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupUUID = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetApplicationNameList(v []*string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAssetCount(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AssetCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetBusiness(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Business = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryClassId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetDomain(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAcl(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAclRecommend(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetInBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetIsMarkNormal(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOrganization(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Organization = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOutBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetPrivateAssetCount(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.PrivateAssetCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecurityReason(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecuritySuggest(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSessionCount(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTagList(v []*DescribeOutgoingDomainResponseBodyDomainListTagList) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTotalBytes(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) Validate() error {
	return dara.Validate(s)
}

type DescribeOutgoingDomainResponseBodyDomainListTagList struct {
	// The type of the tag. Valid values:
	//
	// 	- **Suspicious**
	//
	// 	- **Malicious**
	//
	// 	- **Trusted**
	//
	// example:
	//
	// Trusted
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the tag.
	//
	// example:
	//
	// Tag indicating that the domain name is added to the whitelist
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the tag.
	//
	// example:
	//
	// AliYun
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// Tag indicating that the domain name is added to the whitelist
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetClassId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetRiskLevel(v int32) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagDescribe(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagName(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) Validate() error {
	return dara.Validate(s)
}
