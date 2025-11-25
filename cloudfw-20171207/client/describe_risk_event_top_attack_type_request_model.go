// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuyVersion(v string) *DescribeRiskEventTopAttackTypeRequest
	GetBuyVersion() *string
	SetDirection(v string) *DescribeRiskEventTopAttackTypeRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeRiskEventTopAttackTypeRequest
	GetEndTime() *string
	SetLang(v string) *DescribeRiskEventTopAttackTypeRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeRiskEventTopAttackTypeRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeRiskEventTopAttackTypeRequest
	GetStartTime() *string
}

type DescribeRiskEventTopAttackTypeRequest struct {
	// example:
	//
	// 2
	BuyVersion *string `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1743387943
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 140.210.153.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1670307484
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRiskEventTopAttackTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetBuyVersion() *string {
	return s.BuyVersion
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskEventTopAttackTypeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetBuyVersion(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetDirection(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetEndTime(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetLang(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetSourceIp(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) SetStartTime(v string) *DescribeRiskEventTopAttackTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeRequest) Validate() error {
	return dara.Validate(s)
}
