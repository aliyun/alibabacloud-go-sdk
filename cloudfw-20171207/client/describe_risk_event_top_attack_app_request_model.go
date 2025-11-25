// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v []*string) *DescribeRiskEventTopAttackAppRequest
	GetAttackApp() []*string
	SetAttackType(v string) *DescribeRiskEventTopAttackAppRequest
	GetAttackType() *string
	SetBuyVersion(v string) *DescribeRiskEventTopAttackAppRequest
	GetBuyVersion() *string
	SetEndTime(v string) *DescribeRiskEventTopAttackAppRequest
	GetEndTime() *string
	SetLang(v string) *DescribeRiskEventTopAttackAppRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeRiskEventTopAttackAppRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeRiskEventTopAttackAppRequest
	GetStartTime() *string
}

type DescribeRiskEventTopAttackAppRequest struct {
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
	// 1735784888
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 182.150.22.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRiskEventTopAttackAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAppRequest) GetAttackApp() []*string {
	return s.AttackApp
}

func (s *DescribeRiskEventTopAttackAppRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeRiskEventTopAttackAppRequest) GetBuyVersion() *string {
	return s.BuyVersion
}

func (s *DescribeRiskEventTopAttackAppRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventTopAttackAppRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskEventTopAttackAppRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskEventTopAttackAppRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventTopAttackAppRequest) SetAttackApp(v []*string) *DescribeRiskEventTopAttackAppRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetAttackType(v string) *DescribeRiskEventTopAttackAppRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetBuyVersion(v string) *DescribeRiskEventTopAttackAppRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetEndTime(v string) *DescribeRiskEventTopAttackAppRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetLang(v string) *DescribeRiskEventTopAttackAppRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetSourceIp(v string) *DescribeRiskEventTopAttackAppRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) SetStartTime(v string) *DescribeRiskEventTopAttackAppRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppRequest) Validate() error {
	return dara.Validate(s)
}
