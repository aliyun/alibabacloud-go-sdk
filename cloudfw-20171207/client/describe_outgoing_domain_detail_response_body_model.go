// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationNameList(v []*string) *DescribeOutgoingDomainDetailResponseBody
	GetApplicationNameList() []*string
	SetAssetList(v []*DescribeOutgoingDomainDetailResponseBodyAssetList) *DescribeOutgoingDomainDetailResponseBody
	GetAssetList() []*DescribeOutgoingDomainDetailResponseBodyAssetList
	SetRequestId(v string) *DescribeOutgoingDomainDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDomainDetailResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingDomainDetailResponseBody struct {
	ApplicationNameList []*string                                            `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	AssetList           []*DescribeOutgoingDomainDetailResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// example:
	//
	// 98AF5888-9606-59CF-888F-032A9ED0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainDetailResponseBody) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeOutgoingDomainDetailResponseBody) GetAssetList() []*DescribeOutgoingDomainDetailResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeOutgoingDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDomainDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDomainDetailResponseBody) SetApplicationNameList(v []*string) *DescribeOutgoingDomainDetailResponseBody {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBody) SetAssetList(v []*DescribeOutgoingDomainDetailResponseBodyAssetList) *DescribeOutgoingDomainDetailResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBody) SetRequestId(v string) *DescribeOutgoingDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBody) SetTotalCount(v int32) *DescribeOutgoingDomainDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBody) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDomainDetailResponseBodyAssetList struct {
	// example:
	//
	// FullCoverage
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1738548990
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 244438.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 32
	IpsHitCnt *int64 `json:"IpsHitCnt,omitempty" xml:"IpsHitCnt,omitempty"`
	// example:
	//
	// 1739326614
	LastTime *int32 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// ngw-bp1utx6wj4x9qu9tl****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// example:
	//
	// 100
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 10.21.242XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// example:
	//
	// 47.96.181.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// i-uf63wjhyoohc1g4z****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// test
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// example:
	//
	// NatEIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 08a58465-ff4d-4c47-8782-eb008301****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 10
	SessionCount *int64                                                      `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	TagList      []*DescribeOutgoingDomainDetailResponseBodyAssetListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// example:
	//
	// 321120825843
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// example:
	//
	// vpc-9dp16jgwgyvn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeOutgoingDomainDetailResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainDetailResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetFirstTime() *int32 {
	return s.FirstTime
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetIpsHitCnt() *int64 {
	return s.IpsHitCnt
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetLastTime() *int32 {
	return s.LastTime
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetTagList() []*DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	return s.TagList
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetAclCoverage(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetDomain(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetFirstTime(v int32) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.FirstTime = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetInBytes(v int64) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetIpsHitCnt(v int64) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.IpsHitCnt = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetLastTime(v int32) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.LastTime = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetNatGatewayId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetNatGatewayName(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetOutBytes(v int64) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetPrivateIP(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetPublicIP(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetRegionNo(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.RegionNo = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetResourceInstanceId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetResourceInstanceName(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetResourceType(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetRuleId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetRuleName(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetSessionCount(v int64) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetTagList(v []*DescribeOutgoingDomainDetailResponseBodyAssetListTagList) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetTotalBytes(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) SetVpcId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetList {
	s.VpcId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetList) Validate() error {
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

type DescribeOutgoingDomainDetailResponseBodyAssetListTagList struct {
	// example:
	//
	// 3
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// example:
	//
	// 0
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// example:
	//
	// FirstFlow
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDomainDetailResponseBodyAssetListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) SetClassId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) SetRiskLevel(v int32) *DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) SetTagDescribe(v string) *DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) SetTagId(v string) *DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) SetTagName(v string) *DescribeOutgoingDomainDetailResponseBodyAssetListTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingDomainDetailResponseBodyAssetListTagList) Validate() error {
	return dara.Validate(s)
}
