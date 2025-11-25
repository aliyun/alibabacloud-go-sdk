// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v []*string) *DescribeRiskEventStatisticRequest
	GetAttackApp() []*string
	SetAttackType(v string) *DescribeRiskEventStatisticRequest
	GetAttackType() *string
	SetBuyVersion(v string) *DescribeRiskEventStatisticRequest
	GetBuyVersion() *string
	SetEndTime(v string) *DescribeRiskEventStatisticRequest
	GetEndTime() *string
	SetLang(v string) *DescribeRiskEventStatisticRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeRiskEventStatisticRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeRiskEventStatisticRequest
	GetStartTime() *string
}

type DescribeRiskEventStatisticRequest struct {
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// example:
	//
	// 10
	BuyVersion *string `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1534408267
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 218.76.30.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRiskEventStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventStatisticRequest) GetAttackApp() []*string {
	return s.AttackApp
}

func (s *DescribeRiskEventStatisticRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeRiskEventStatisticRequest) GetBuyVersion() *string {
	return s.BuyVersion
}

func (s *DescribeRiskEventStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskEventStatisticRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskEventStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventStatisticRequest) SetAttackApp(v []*string) *DescribeRiskEventStatisticRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetAttackType(v string) *DescribeRiskEventStatisticRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetBuyVersion(v string) *DescribeRiskEventStatisticRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetEndTime(v string) *DescribeRiskEventStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetLang(v string) *DescribeRiskEventStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetSourceIp(v string) *DescribeRiskEventStatisticRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) SetStartTime(v string) *DescribeRiskEventStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventStatisticRequest) Validate() error {
	return dara.Validate(s)
}
