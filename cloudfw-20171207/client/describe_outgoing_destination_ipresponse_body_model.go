// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDstIPList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPList) *DescribeOutgoingDestinationIPResponseBody
	GetDstIPList() []*DescribeOutgoingDestinationIPResponseBodyDstIPList
	SetRequestId(v string) *DescribeOutgoingDestinationIPResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDestinationIPResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingDestinationIPResponseBody struct {
	// A list of destination IP addresses for outgoing connections.
	DstIPList []*DescribeOutgoingDestinationIPResponseBodyDstIPList `json:"DstIPList,omitempty" xml:"DstIPList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of outgoing IPs.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBody) GetDstIPList() []*DescribeOutgoingDestinationIPResponseBodyDstIPList {
	return s.DstIPList
}

func (s *DescribeOutgoingDestinationIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDestinationIPResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetDstIPList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPList) *DescribeOutgoingDestinationIPResponseBody {
	s.DstIPList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationIPResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) Validate() error {
	if s.DstIPList != nil {
		for _, item := range s.DstIPList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDestinationIPResponseBodyDstIPList struct {
	// Indicates whether an access control policy is applied. Valid values:
	//
	// - **Uncovered**: No policy is applied.
	//
	// - **FullCoverage**: A policy is applied.
	//
	// example:
	//
	// Uncovered
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// Details of the ACL recommendation.
	//
	// example:
	//
	// 建议放行
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
	// A list of address books that contain this destination IP address.
	AddressGroupList []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList `json:"AddressGroupList,omitempty" xml:"AddressGroupList,omitempty" type:"Repeated"`
	// The list of application ports.
	//
	// > This response returns a maximum of 99 application ports. If more than 99 ports exist, only the first 99 are returned.
	ApplicationPortList []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList `json:"ApplicationPortList,omitempty" xml:"ApplicationPortList,omitempty" type:"Repeated"`
	// The total number of assets that initiated outgoing connections to this destination IP.
	//
	// example:
	//
	// 20
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The threat intelligence category of the destination IP address. Valid values:
	//
	// - **Suspicious**
	//
	// - **Malicious**
	//
	// - **Trusted**
	//
	// example:
	//
	// Trusted
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The ID of the service category. Valid values:
	//
	// - **Aliyun**: The destination is an Alibaba Cloud product.
	//
	// - **NotAliyun**: The destination is a non-Alibaba Cloud product.
	//
	// example:
	//
	// Aliyun
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The service category of the destination IP address. Valid values:
	//
	// - **Alibaba Cloud product**
	//
	// - **non-Alibaba Cloud product**
	//
	// example:
	//
	// 阿里云产品
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The destination IP address of the outgoing connection.
	//
	// example:
	//
	// 10.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The name of the rule group.
	//
	// example:
	//
	// rules_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an access control rule exists. Valid values:
	//
	// - **true**: An access control rule exists.
	//
	// - **false**: No access control rule exists.
	//
	// example:
	//
	// true
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an ACL is recommended. Valid values:
	//
	// - **true**: An ACL is recommended.
	//
	// - **false**: No ACL is recommended.
	//
	// example:
	//
	// true
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The total inbound traffic in bytes.
	//
	// example:
	//
	// 472
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the destination IP address is added to the allowlist. Valid values:
	//
	// - **true**: The destination IP address is on the allowlist.
	//
	// - **false**: The destination IP address is not on the allowlist.
	//
	// example:
	//
	// true
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The geographical location of the destination IP address.
	//
	// example:
	//
	// 山东省青岛市
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// The total outbound traffic in bytes.
	//
	// example:
	//
	// 965
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The total number of private assets that initiated outgoing connections to this destination IP.
	//
	// example:
	//
	// 20
	PrivateAssetCount *int64 `json:"PrivateAssetCount,omitempty" xml:"PrivateAssetCount,omitempty"`
	// The UUID of the ACL rule.
	//
	// example:
	//
	// fadsfd-dfadf-df****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the ACL rule.
	//
	// example:
	//
	// 默认规则
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason for the security recommendation.
	//
	// example:
	//
	// 智能策略：该目的域名所属组织为阿里云计算有限公司，主要业务为阿里云，未发现安全风险，可用于配置外联白名单。
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The recommended security action for the outgoing connection. Valid values:
	//
	// - **pass**: Allows the connection.
	//
	// - **alert**: Rejects the connection.
	//
	// - **drop**: Drops the connection.
	//
	// example:
	//
	// pass
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 4
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// A list of tags associated with the destination IP.
	TagList []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total traffic volume in bytes.
	//
	// example:
	//
	// 800
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetAclRecommendDetail() *string {
	return s.AclRecommendDetail
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetAddressGroupList() []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	return s.AddressGroupList
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetApplicationPortList() []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	return s.ApplicationPortList
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetAssetCount() *int64 {
	return s.AssetCount
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetCategoryClassId() *string {
	return s.CategoryClassId
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetHasAcl() *string {
	return s.HasAcl
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetHasAclRecommend() *bool {
	return s.HasAclRecommend
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetIsMarkNormal() *bool {
	return s.IsMarkNormal
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetLocationName() *string {
	return s.LocationName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetPrivateAssetCount() *int64 {
	return s.PrivateAssetCount
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetSecurityReason() *string {
	return s.SecurityReason
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetSecuritySuggest() *string {
	return s.SecuritySuggest
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetTagList() []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	return s.TagList
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclCoverage(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclRecommendDetail(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclStatus(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAddressGroupList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AddressGroupList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetApplicationPortList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.ApplicationPortList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAssetCount(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AssetCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetDstIP(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAcl(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAclRecommend(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetInBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetIsMarkNormal(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetLocationName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.LocationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetOutBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetPrivateAssetCount(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.PrivateAssetCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecurityReason(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecuritySuggest(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSessionCount(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTagList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTotalBytes(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) Validate() error {
	if s.AddressGroupList != nil {
		for _, item := range s.AddressGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList struct {
	// The name of the address book.
	//
	// example:
	//
	// IP地址簿
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) GetAddressGroupName() *string {
	return s.AddressGroupName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) GetAddressGroupUUID() *string {
	return s.AddressGroupUUID
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupUUID(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupUUID = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) Validate() error {
	return dara.Validate(s)
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList struct {
	// The application protocol detected for the connection. Valid values:
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **Memcache**
	//
	// - **MongoDB**
	//
	// - **MQTT**
	//
	// - **MySQL**
	//
	// - **RDP**
	//
	// - **Redis**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **SSH**
	//
	// - **SSL_No_Cert**
	//
	// - **SSL**
	//
	// - **VNC**
	//
	// >
	//
	// example:
	//
	// HTTP
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// A list of reasons why the application protocol was not identified.
	UnknownReason []*string `json:"UnknownReason,omitempty" xml:"UnknownReason,omitempty" type:"Repeated"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GetUnknownReason() []*string {
	return s.UnknownReason
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetApplicationName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetPort(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetUnknownReason(v []*string) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.UnknownReason = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) Validate() error {
	return dara.Validate(s)
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListTagList struct {
	// The category of the threat intelligence tag. Valid values:
	//
	// - **Suspicious**
	//
	// - **Malicious**
	//
	// - **Trusted**
	//
	// example:
	//
	// Trusted
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// - **1**: Low
	//
	// - **2**: Medium
	//
	// - **3**: High
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the threat intelligence tag.
	//
	// example:
	//
	// ReleaseLabel
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the threat intelligence tag.
	//
	// example:
	//
	// ReleaseLabel
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the threat intelligence tag.
	//
	// example:
	//
	// ReleaseLabel
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetRiskLevel(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagDescribe(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) Validate() error {
	return dara.Validate(s)
}
