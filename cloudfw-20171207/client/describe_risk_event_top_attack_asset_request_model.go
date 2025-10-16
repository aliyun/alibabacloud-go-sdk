// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v []*string) *DescribeRiskEventTopAttackAssetRequest
	GetAttackApp() []*string
	SetAttackType(v string) *DescribeRiskEventTopAttackAssetRequest
	GetAttackType() *string
	SetBuyVersion(v string) *DescribeRiskEventTopAttackAssetRequest
	GetBuyVersion() *string
	SetEndTime(v string) *DescribeRiskEventTopAttackAssetRequest
	GetEndTime() *string
	SetLang(v string) *DescribeRiskEventTopAttackAssetRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeRiskEventTopAttackAssetRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeRiskEventTopAttackAssetRequest
	GetStartTime() *string
}

type DescribeRiskEventTopAttackAssetRequest struct {
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// example:
	//
	// 10
	BuyVersion *string `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1742955867
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 125.33.253.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1754273436
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRiskEventTopAttackAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAssetRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetAttackApp() []*string {
	return s.AttackApp
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetBuyVersion() *string {
	return s.BuyVersion
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskEventTopAttackAssetRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetAttackApp(v []*string) *DescribeRiskEventTopAttackAssetRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetAttackType(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetBuyVersion(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetEndTime(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetLang(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetSourceIp(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) SetStartTime(v string) *DescribeRiskEventTopAttackAssetRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetRequest) Validate() error {
	return dara.Validate(s)
}
