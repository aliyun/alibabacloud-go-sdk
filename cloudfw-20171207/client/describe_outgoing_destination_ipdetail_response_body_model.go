// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationIPDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*DescribeOutgoingDestinationIPDetailResponseBodyAssetList) *DescribeOutgoingDestinationIPDetailResponseBody
	GetAssetList() []*DescribeOutgoingDestinationIPDetailResponseBodyAssetList
	SetIspName(v string) *DescribeOutgoingDestinationIPDetailResponseBody
	GetIspName() *string
	SetLocationName(v string) *DescribeOutgoingDestinationIPDetailResponseBody
	GetLocationName() *string
	SetRequestId(v string) *DescribeOutgoingDestinationIPDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDestinationIPDetailResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingDestinationIPDetailResponseBody struct {
	AssetList    []*DescribeOutgoingDestinationIPDetailResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	IspName      *string                                                     `json:"IspName,omitempty" xml:"IspName,omitempty"`
	LocationName *string                                                     `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDestinationIPDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) GetAssetList() []*DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) GetIspName() *string {
	return s.IspName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) GetLocationName() *string {
	return s.LocationName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) SetAssetList(v []*DescribeOutgoingDestinationIPDetailResponseBodyAssetList) *DescribeOutgoingDestinationIPDetailResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) SetIspName(v string) *DescribeOutgoingDestinationIPDetailResponseBody {
	s.IspName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) SetLocationName(v string) *DescribeOutgoingDestinationIPDetailResponseBody {
	s.LocationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationIPDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationIPDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBody) Validate() error {
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

type DescribeOutgoingDestinationIPDetailResponseBodyAssetList struct {
	// example:
	//
	// FullCoverage
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// example:
	//
	// 1744682438
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 0.0
	InBytes *int64  `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// example:
	//
	// 1739326614
	LastTime     *int32  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
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
	// 0.0
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
	// cn-shanghai
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
	// EcsPublicIP
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
	// 2
	SessionCount *int64                                                             `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	TagList      []*DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// example:
	//
	// 458681
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// example:
	//
	// vpc-9dp16jgwgyvn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeOutgoingDestinationIPDetailResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetAclCoverage() *string {
	return s.AclCoverage
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetFirstTime() *int32 {
	return s.FirstTime
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetIspName() *string {
	return s.IspName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetLastTime() *int32 {
	return s.LastTime
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetLocationName() *string {
	return s.LocationName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetTagList() []*DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	return s.TagList
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetAclCoverage(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetFirstTime(v int32) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.FirstTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetInBytes(v int64) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetIspName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.IspName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetLastTime(v int32) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.LastTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetLocationName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.LocationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetNatGatewayId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetNatGatewayName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetOutBytes(v int64) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetPrivateIP(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetPublicIP(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetRegionNo(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.RegionNo = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetResourceInstanceId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetResourceInstanceName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetResourceType(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetRuleId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetRuleName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetSessionCount(v int64) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetTagList(v []*DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetTotalBytes(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) SetVpcId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetList {
	s.VpcId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetList) Validate() error {
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

type DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList struct {
	// example:
	//
	// 1
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

func (s DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) SetClassId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) SetRiskLevel(v int32) *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) SetTagDescribe(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) SetTagId(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) SetTagName(v string) *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPDetailResponseBodyAssetListTagList) Validate() error {
	return dara.Validate(s)
}
