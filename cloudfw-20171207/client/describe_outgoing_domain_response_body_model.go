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
	// The list of Outbound Domain names.
	DomainList []*DescribeOutgoingDomainResponseBodyDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Outbound Domain names.
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
	if s.DomainList != nil {
		for _, item := range s.DomainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDomainResponseBodyDomainList struct {
	// Indicates whether the access control policy covers the domain name. Valid values:
	//
	// - **Uncovered**: Not covered.
	//
	// - **FullCoverage**: Covered.
	//
	// example:
	//
	// Uncovered
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The ACL recommendation details.
	//
	// example:
	//
	// RecommendedRelease
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The health status of the access control policy. Valid values:
	//
	// - **Normal**: Healthy.
	//
	// - **Abnormal**: Unhealthy.
	//
	// example:
	//
	// Normal
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// Outreach Address Book
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	//
	// example:
	//
	// fdad-fdafa-dafa-dfa****
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	// The application names.
	ApplicationNameList []*string                                                          `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	ApplicationPortList []*DescribeOutgoingDomainResponseBodyDomainListApplicationPortList `json:"ApplicationPortList,omitempty" xml:"ApplicationPortList,omitempty" type:"Repeated"`
	// The total number of assets that initiate outbound connections.
	//
	// example:
	//
	// 20
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The website business.
	//
	// example:
	//
	// Aliyun
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// The categorization of the intelligence tags label. Valid values:
	//
	// - **Suspicious**: Suspicious.
	//
	// - **Malicious**: Malicious.
	//
	// - **Trusted**: Trusted.
	//
	// example:
	//
	// Trusted
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The product category ID. Valid values:
	//
	// - **Aliyun**: Alibaba Cloud product.
	//
	// - **NotAliyun**: Non-Alibaba Cloud product.
	//
	// example:
	//
	// Aliyun
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name of the product. Valid values:
	//
	// - **Alibaba Cloud product**
	//
	// - **Non-Alibaba Cloud product**
	//
	// example:
	//
	// Alibaba Cloud product
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The domain name of outbound connections.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The group name of the rule.
	//
	// example:
	//
	// group-name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an ACL already covers this domain name. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an ACL recommendation exists. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The inbound traffic.
	//
	// example:
	//
	// 3214
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the Outbound Domain is marked as normal. Valid values:
	//
	// - **true**: Normal.
	//
	// - **false**: Abnormal.
	//
	// example:
	//
	// true
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The organization name.
	//
	// example:
	//
	// Alibaba Cloud Computing Limited
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The outbound traffic.
	//
	// example:
	//
	// 4582
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The total number of private network assets that initiate outbound connections.
	//
	// example:
	//
	// 20
	PrivateAssetCount *int64 `json:"PrivateAssetCount,omitempty" xml:"PrivateAssetCount,omitempty"`
	// The ACL rule ID.
	//
	// example:
	//
	// add-dfadf-f****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ACL rule name.
	//
	// example:
	//
	// acl-name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The security reason.
	//
	// example:
	//
	// Smart policy: The target domain name belongs to Aliyun Computing Co., Ltd., and its main business is Aliyun. No security risks have been found. It can be used to configure an outreach whitelist.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The security policy for the Outbound Domain of outbound connections. Valid values:
	//
	// - **pass**: Allow.
	//
	// - **alert**: Monitor.
	//
	// - **drop**: Deny.
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
	// The tag list.
	TagList []*DescribeOutgoingDomainResponseBodyDomainListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total traffic. Unit: bytes.
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

func (s *DescribeOutgoingDomainResponseBodyDomainList) GetApplicationPortList() []*DescribeOutgoingDomainResponseBodyDomainListApplicationPortList {
	return s.ApplicationPortList
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

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetApplicationPortList(v []*DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) *DescribeOutgoingDomainResponseBodyDomainList {
	s.ApplicationPortList = v
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
	if s.ApplicationPortList != nil {
		for _, item := range s.ApplicationPortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDomainResponseBodyDomainListApplicationPortList struct {
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) SetApplicationName(v string) *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) SetPort(v int32) *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListApplicationPortList) Validate() error {
	return dara.Validate(s)
}

type DescribeOutgoingDomainResponseBodyDomainListTagList struct {
	// The categorization of the intelligence tags label. Valid values:
	//
	// - **Suspicious**: Suspicious.
	//
	// - **Malicious**: Malicious.
	//
	// - **Trusted**: Trusted.
	//
	// example:
	//
	// Trusted
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk assessment level. Valid values:
	//
	// - **1**: Low.
	//
	// - **2**: Medium.
	//
	// - **3**: High.
	//
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The tag description.
	//
	// example:
	//
	// tag-describe
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The intelligence tags label ID.
	//
	// example:
	//
	// AliYun
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// example:
	//
	// tag-name
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
