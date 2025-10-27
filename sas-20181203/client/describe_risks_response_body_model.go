// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRisksResponseBody
	GetRequestId() *string
	SetRisks(v []*DescribeRisksResponseBodyRisks) *DescribeRisksResponseBody
	GetRisks() []*DescribeRisksResponseBodyRisks
	SetTotalCount(v int32) *DescribeRisksResponseBody
	GetTotalCount() *int32
}

type DescribeRisksResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EA54FE21-B006-5DFF-8D64-C4FFECDA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The baselines.
	Risks []*DescribeRisksResponseBodyRisks `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRisksResponseBody) GetRisks() []*DescribeRisksResponseBodyRisks {
	return s.Risks
}

func (s *DescribeRisksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRisksResponseBody) SetRequestId(v string) *DescribeRisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRisksResponseBody) SetRisks(v []*DescribeRisksResponseBodyRisks) *DescribeRisksResponseBody {
	s.Risks = v
	return s
}

func (s *DescribeRisksResponseBody) SetTotalCount(v int32) *DescribeRisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRisksResponseBody) Validate() error {
	if s.Risks != nil {
		for _, item := range s.Risks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRisksResponseBodyRisks struct {
	// The description of the baseline.
	//
	// example:
	//
	// Ubuntu 14,Ubuntu 16 ,Ubuntu 18,Ubuntu 20 baseline based on Alibaba Cloud best security practices
	RiskDetail *string `json:"RiskDetail,omitempty" xml:"RiskDetail,omitempty"`
	// The baseline ID.
	//
	// example:
	//
	// 54
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Alibaba Cloud Standard - Ubuntu Security Baseline
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The name of the baseline type.
	//
	// example:
	//
	// cis
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The name of the baseline subtype.
	//
	// example:
	//
	// hc_ubuntu
	SubRiskType *string `json:"SubRiskType,omitempty" xml:"SubRiskType,omitempty"`
	// The display name of the baseline subtype.
	//
	// example:
	//
	// Alibaba Cloud Standard - Ubuntu Security Baseline
	SubTypeAlias *string `json:"SubTypeAlias,omitempty" xml:"SubTypeAlias,omitempty"`
	// The display name of the baseline type.
	//
	// example:
	//
	// Best security practices
	TypeAlias *string `json:"TypeAlias,omitempty" xml:"TypeAlias,omitempty"`
}

func (s DescribeRisksResponseBodyRisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRisksResponseBodyRisks) GoString() string {
	return s.String()
}

func (s *DescribeRisksResponseBodyRisks) GetRiskDetail() *string {
	return s.RiskDetail
}

func (s *DescribeRisksResponseBodyRisks) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeRisksResponseBodyRisks) GetRiskName() *string {
	return s.RiskName
}

func (s *DescribeRisksResponseBodyRisks) GetRiskType() *string {
	return s.RiskType
}

func (s *DescribeRisksResponseBodyRisks) GetSubRiskType() *string {
	return s.SubRiskType
}

func (s *DescribeRisksResponseBodyRisks) GetSubTypeAlias() *string {
	return s.SubTypeAlias
}

func (s *DescribeRisksResponseBodyRisks) GetTypeAlias() *string {
	return s.TypeAlias
}

func (s *DescribeRisksResponseBodyRisks) SetRiskDetail(v string) *DescribeRisksResponseBodyRisks {
	s.RiskDetail = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetRiskId(v int64) *DescribeRisksResponseBodyRisks {
	s.RiskId = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetRiskName(v string) *DescribeRisksResponseBodyRisks {
	s.RiskName = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetRiskType(v string) *DescribeRisksResponseBodyRisks {
	s.RiskType = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetSubRiskType(v string) *DescribeRisksResponseBodyRisks {
	s.SubRiskType = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetSubTypeAlias(v string) *DescribeRisksResponseBodyRisks {
	s.SubTypeAlias = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) SetTypeAlias(v string) *DescribeRisksResponseBodyRisks {
	s.TypeAlias = &v
	return s
}

func (s *DescribeRisksResponseBodyRisks) Validate() error {
	return dara.Validate(s)
}
