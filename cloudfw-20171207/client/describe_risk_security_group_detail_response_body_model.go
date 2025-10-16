// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskSecurityGroupDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeRiskSecurityGroupDetailResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeRiskSecurityGroupDetailResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRiskSecurityGroupDetailResponseBody
	GetRequestId() *string
	SetRiskSgDetail(v []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) *DescribeRiskSecurityGroupDetailResponseBody
	GetRiskSgDetail() []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail
	SetTotalCount(v int32) *DescribeRiskSecurityGroupDetailResponseBody
	GetTotalCount() *int32
}

type DescribeRiskSecurityGroupDetailResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 33C94306-2064-5A06-9645-01419967****
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskSgDetail []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail `json:"RiskSgDetail,omitempty" xml:"RiskSgDetail,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) GetRiskSgDetail() []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	return s.RiskSgDetail
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) SetPageNo(v int32) *DescribeRiskSecurityGroupDetailResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) SetPageSize(v int32) *DescribeRiskSecurityGroupDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) SetRequestId(v string) *DescribeRiskSecurityGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) SetRiskSgDetail(v []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) *DescribeRiskSecurityGroupDetailResponseBody {
	s.RiskSgDetail = v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) SetTotalCount(v int32) *DescribeRiskSecurityGroupDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBody) Validate() error {
	if s.RiskSgDetail != nil {
		for _, item := range s.RiskSgDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail struct {
	// example:
	//
	// 10
	EcsCount *int32                                                            `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	EcsInfo  []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo `json:"EcsInfo,omitempty" xml:"EcsInfo,omitempty" type:"Repeated"`
	// example:
	//
	// sg-2vc0p803vgxumn6r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test-instance-name
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// high
	RiskLevel *string                                                            `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RuleInfo  []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-8vbuzirdl3w1r7exw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetEcsInfo() []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo {
	return s.EcsInfo
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetRuleInfo() []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	return s.RuleInfo
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetEcsCount(v int32) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.EcsCount = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetEcsInfo(v []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.EcsInfo = v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetInstanceId(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.InstanceId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetInstanceName(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.InstanceName = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetRegionNo(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetRiskLevel(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetRuleInfo(v []*DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.RuleInfo = v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) SetVpcId(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail {
	s.VpcId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetail) Validate() error {
	if s.EcsInfo != nil {
		for _, item := range s.EcsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleInfo != nil {
		for _, item := range s.RuleInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo struct {
	// example:
	//
	// i-bp1gra23yai47d8e****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// example:
	//
	// tyf_test
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// example:
	//
	// 172.24.121.XXX
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// 47.107.141.XXX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) SetEcsInstanceId(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) SetEcsInstanceName(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) SetPrivateIp(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo {
	s.PrivateIp = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) SetPublicIp(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo {
	s.PublicIp = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailEcsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// enable
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// example:
	//
	// 4c1e72c9-6690-408b-9048-065f0f10****
	RuleUuid *string `json:"RuleUuid,omitempty" xml:"RuleUuid,omitempty"`
	// example:
	//
	// test
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetRuleUuid() *string {
	return s.RuleUuid
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetDescription(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.Description = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetRiskLevel(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetRuleName(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.RuleName = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetRuleStatus(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetRuleUuid(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.RuleUuid = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) SetSuggestion(v string) *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo {
	s.Suggestion = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponseBodyRiskSgDetailRuleInfo) Validate() error {
	return dara.Validate(s)
}
