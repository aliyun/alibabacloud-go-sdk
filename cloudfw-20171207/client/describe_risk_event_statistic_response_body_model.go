// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackAppCnt(v int32) *DescribeRiskEventStatisticResponseBody
	GetAttackAppCnt() *int32
	SetAttackCnt(v int32) *DescribeRiskEventStatisticResponseBody
	GetAttackCnt() *int32
	SetAttackIpCnt(v int32) *DescribeRiskEventStatisticResponseBody
	GetAttackIpCnt() *int32
	SetDropCnt(v int32) *DescribeRiskEventStatisticResponseBody
	GetDropCnt() *int32
	SetRequestId(v string) *DescribeRiskEventStatisticResponseBody
	GetRequestId() *string
}

type DescribeRiskEventStatisticResponseBody struct {
	// example:
	//
	// 10
	AttackAppCnt *int32 `json:"AttackAppCnt,omitempty" xml:"AttackAppCnt,omitempty"`
	// example:
	//
	// 5
	AttackCnt *int32 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// example:
	//
	// 10
	AttackIpCnt *int32 `json:"AttackIpCnt,omitempty" xml:"AttackIpCnt,omitempty"`
	// example:
	//
	// 20
	DropCnt *int32 `json:"DropCnt,omitempty" xml:"DropCnt,omitempty"`
	// example:
	//
	// F93A490D-9E92-5AA4-BA79-600FFC09****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskEventStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventStatisticResponseBody) GetAttackAppCnt() *int32 {
	return s.AttackAppCnt
}

func (s *DescribeRiskEventStatisticResponseBody) GetAttackCnt() *int32 {
	return s.AttackCnt
}

func (s *DescribeRiskEventStatisticResponseBody) GetAttackIpCnt() *int32 {
	return s.AttackIpCnt
}

func (s *DescribeRiskEventStatisticResponseBody) GetDropCnt() *int32 {
	return s.DropCnt
}

func (s *DescribeRiskEventStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventStatisticResponseBody) SetAttackAppCnt(v int32) *DescribeRiskEventStatisticResponseBody {
	s.AttackAppCnt = &v
	return s
}

func (s *DescribeRiskEventStatisticResponseBody) SetAttackCnt(v int32) *DescribeRiskEventStatisticResponseBody {
	s.AttackCnt = &v
	return s
}

func (s *DescribeRiskEventStatisticResponseBody) SetAttackIpCnt(v int32) *DescribeRiskEventStatisticResponseBody {
	s.AttackIpCnt = &v
	return s
}

func (s *DescribeRiskEventStatisticResponseBody) SetDropCnt(v int32) *DescribeRiskEventStatisticResponseBody {
	s.DropCnt = &v
	return s
}

func (s *DescribeRiskEventStatisticResponseBody) SetRequestId(v string) *DescribeRiskEventStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
