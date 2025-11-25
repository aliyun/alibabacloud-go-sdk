// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApps(v []*DescribeRiskEventTopAttackAppResponseBodyAttackApps) *DescribeRiskEventTopAttackAppResponseBody
	GetAttackApps() []*DescribeRiskEventTopAttackAppResponseBodyAttackApps
	SetRequestId(v string) *DescribeRiskEventTopAttackAppResponseBody
	GetRequestId() *string
}

type DescribeRiskEventTopAttackAppResponseBody struct {
	AttackApps []*DescribeRiskEventTopAttackAppResponseBodyAttackApps `json:"AttackApps,omitempty" xml:"AttackApps,omitempty" type:"Repeated"`
	// example:
	//
	// C9DDAD29-C6B3-5997-B757-FFB3F1C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskEventTopAttackAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAppResponseBody) GetAttackApps() []*DescribeRiskEventTopAttackAppResponseBodyAttackApps {
	return s.AttackApps
}

func (s *DescribeRiskEventTopAttackAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventTopAttackAppResponseBody) SetAttackApps(v []*DescribeRiskEventTopAttackAppResponseBodyAttackApps) *DescribeRiskEventTopAttackAppResponseBody {
	s.AttackApps = v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponseBody) SetRequestId(v string) *DescribeRiskEventTopAttackAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponseBody) Validate() error {
	if s.AttackApps != nil {
		for _, item := range s.AttackApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskEventTopAttackAppResponseBodyAttackApps struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 20
	AttackCnt *int32 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// example:
	//
	// 15
	DropCnt *int32 `json:"DropCnt,omitempty" xml:"DropCnt,omitempty"`
}

func (s DescribeRiskEventTopAttackAppResponseBodyAttackApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAppResponseBodyAttackApps) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) GetApp() *string {
	return s.App
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) GetAttackCnt() *int32 {
	return s.AttackCnt
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) GetDropCnt() *int32 {
	return s.DropCnt
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) SetApp(v string) *DescribeRiskEventTopAttackAppResponseBodyAttackApps {
	s.App = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) SetAttackCnt(v int32) *DescribeRiskEventTopAttackAppResponseBodyAttackApps {
	s.AttackCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) SetDropCnt(v int32) *DescribeRiskEventTopAttackAppResponseBodyAttackApps {
	s.DropCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponseBodyAttackApps) Validate() error {
	return dara.Validate(s)
}
