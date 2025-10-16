// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCfwRiskLevelSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCfwRiskLevelSummaryResponseBody
	GetRequestId() *string
	SetRiskList(v []*DescribeCfwRiskLevelSummaryResponseBodyRiskList) *DescribeCfwRiskLevelSummaryResponseBody
	GetRiskList() []*DescribeCfwRiskLevelSummaryResponseBodyRiskList
}

type DescribeCfwRiskLevelSummaryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F1F30435-FA0A-52DA-A5DE-C16FD6C171DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of risks.
	RiskList []*DescribeCfwRiskLevelSummaryResponseBodyRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
}

func (s DescribeCfwRiskLevelSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) GetRiskList() []*DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	return s.RiskList
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) SetRequestId(v string) *DescribeCfwRiskLevelSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) SetRiskList(v []*DescribeCfwRiskLevelSummaryResponseBodyRiskList) *DescribeCfwRiskLevelSummaryResponseBody {
	s.RiskList = v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) Validate() error {
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCfwRiskLevelSummaryResponseBodyRiskList struct {
	// The risk levels. Valid values:
	//
	// 	- **medium**
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of at-risk Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// 50
	Num *string `json:"Num,omitempty" xml:"Num,omitempty"`
	// The type.
	//
	// example:
	//
	// ResourceNotProtected
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryResponseBodyRiskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponseBodyRiskList) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) GetLevel() *string {
	return s.Level
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) GetNum() *string {
	return s.Num
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) GetType() *string {
	return s.Type
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetLevel(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Level = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetNum(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Num = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetType(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Type = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) Validate() error {
	return dara.Validate(s)
}
