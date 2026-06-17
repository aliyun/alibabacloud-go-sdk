// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRiskEventTopAttackTypeResponseBody
	GetRequestId() *string
	SetTopAttackTypeList(v []*DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) *DescribeRiskEventTopAttackTypeResponseBody
	GetTopAttackTypeList() []*DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList
	SetTotalAttackCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBody
	GetTotalAttackCnt() *int64
	SetTotalProtectCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBody
	GetTotalProtectCnt() *int64
}

type DescribeRiskEventTopAttackTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BECDBF66-91DA-5B40-8B05-0D26541A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of the top attack types.
	TopAttackTypeList []*DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList `json:"TopAttackTypeList,omitempty" xml:"TopAttackTypeList,omitempty" type:"Repeated"`
	// The total number of attacks.
	//
	// example:
	//
	// 47
	TotalAttackCnt *int64 `json:"TotalAttackCnt,omitempty" xml:"TotalAttackCnt,omitempty"`
	// The total number of protection triggers.
	//
	// example:
	//
	// 65
	TotalProtectCnt *int64 `json:"TotalProtectCnt,omitempty" xml:"TotalProtectCnt,omitempty"`
}

func (s DescribeRiskEventTopAttackTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) GetTopAttackTypeList() []*DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList {
	return s.TopAttackTypeList
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) GetTotalAttackCnt() *int64 {
	return s.TotalAttackCnt
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) GetTotalProtectCnt() *int64 {
	return s.TotalProtectCnt
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) SetRequestId(v string) *DescribeRiskEventTopAttackTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) SetTopAttackTypeList(v []*DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) *DescribeRiskEventTopAttackTypeResponseBody {
	s.TopAttackTypeList = v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) SetTotalAttackCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBody {
	s.TotalAttackCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) SetTotalProtectCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBody {
	s.TotalProtectCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBody) Validate() error {
	if s.TopAttackTypeList != nil {
		for _, item := range s.TopAttackTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList struct {
	// The number of attacks.
	//
	// example:
	//
	// 38
	AttackCnt *int64 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// The attack type of the intrusion prevention event. Valid values:
	//
	// - **1**: abnormal connection
	//
	// - **2**: command execution
	//
	// - **3**: brute-force attack
	//
	// - **4**: scan
	//
	// - **5**: other
	//
	// - **6**: information leakage
	//
	// - **7**: DoS attack
	//
	// - **8**: overflow attack
	//
	// - **9**: web attack
	//
	// - **10**: trojan and backdoor
	//
	// - **11**: virus and worm
	//
	// - **12**: cryptomining
	//
	// - **13**: reverse shell
	//
	// > By default, this API queries for all attack types.
	//
	// example:
	//
	// 2
	AttackType *int64 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The number of protection triggers.
	//
	// example:
	//
	// 42
	ProtectCnt *int64 `json:"ProtectCnt,omitempty" xml:"ProtectCnt,omitempty"`
}

func (s DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) GetAttackCnt() *int64 {
	return s.AttackCnt
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) GetAttackType() *int64 {
	return s.AttackType
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) GetProtectCnt() *int64 {
	return s.ProtectCnt
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) SetAttackCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList {
	s.AttackCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) SetAttackType(v int64) *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) SetProtectCnt(v int64) *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList {
	s.ProtectCnt = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponseBodyTopAttackTypeList) Validate() error {
	return dara.Validate(s)
}
