// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLevelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRiskLevelsResponseBody
	GetRequestId() *string
	SetRiskLevelList(v []*DescribeRiskLevelsResponseBodyRiskLevelList) *DescribeRiskLevelsResponseBody
	GetRiskLevelList() []*DescribeRiskLevelsResponseBodyRiskLevelList
}

type DescribeRiskLevelsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 136082B3-B21F-5E9D-B68E-991FFD205D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of sensitivity levels.
	RiskLevelList []*DescribeRiskLevelsResponseBodyRiskLevelList `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
}

func (s DescribeRiskLevelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskLevelsResponseBody) GetRiskLevelList() []*DescribeRiskLevelsResponseBodyRiskLevelList {
	return s.RiskLevelList
}

func (s *DescribeRiskLevelsResponseBody) SetRequestId(v string) *DescribeRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskLevelsResponseBody) SetRiskLevelList(v []*DescribeRiskLevelsResponseBodyRiskLevelList) *DescribeRiskLevelsResponseBody {
	s.RiskLevelList = v
	return s
}

func (s *DescribeRiskLevelsResponseBody) Validate() error {
	if s.RiskLevelList != nil {
		for _, item := range s.RiskLevelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskLevelsResponseBodyRiskLevelList struct {
	// The description of the sensitivity level. You can enter a custom description.
	//
	// The following list describes the sensitivity level names and the corresponding default description:
	//
	// 	- **NA**: which indicates that no sensitive data is detected.
	//
	// 	- **S1**: which indicates the sensitive data at sensitivity level 1.
	//
	// 	- **S2**: which indicates the sensitive data at sensitivity level 2.
	//
	// 	- **S3**: which indicates the sensitive data at sensitivity level 3.
	//
	// 	- **S4**: which indicates the sensitive data at sensitivity level 4.
	//
	// 	- **S5**: which indicates the sensitive data at sensitivity level 5.
	//
	// 	- **S6**: which indicates the sensitive data at sensitivity level 6.
	//
	// 	- **S7**: which indicates the sensitive data at sensitivity level 7.
	//
	// 	- **S8**: which indicates the sensitive data at sensitivity level 8.
	//
	// 	- **S9**: which indicates the sensitive data at sensitivity level 9.
	//
	// 	- **S10**: which indicates the sensitive data at sensitivity level 10.
	//
	// example:
	//
	// Sensitive data at sensitivity level 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the sensitivity level. Valid values: 1 to 11. Each sensitivity level ID maps a sensitivity level. For example, the sensitivity level ID of 2 corresponds to the sensitivity level S1.
	//
	// For more information, see the description of the Name parameter.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the sensitivity level. The highest sensitivity level varies based on rule templates. The highest sensitivity level is S10. The highest sensitivity level of the **Built-in data security classification templates for Alibaba and Ant Group*	- is S4, and that of the **built-in classification templates for financial data*	- and **built-in classification templates for assets*	- is S5. For more information about the built-in classification templates for financial data, see Guidelines for Security Classification of Financial Data and Security Data - JRT 0197-2020. For a copied rule template, the highest sensitivity level is S10. The following list describes the sensitivity level names and the corresponding IDs:
	//
	// 	- **NA**: 1
	//
	// 	- **S1**: 2
	//
	// 	- **S2**: 3
	//
	// 	- **S3**: 4
	//
	// 	- **S4**: 5
	//
	// 	- **S5**: 6
	//
	// 	- **S6**: 7
	//
	// 	- **S7**: 8
	//
	// 	- **S8**: 9
	//
	// 	- **S9**: 10
	//
	// 	- **S10**: 11
	//
	// example:
	//
	// S1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of times that each sensitivity level is referenced in the rule template. Default value: 0.
	//
	// example:
	//
	// 20
	ReferenceNum *int32 `json:"ReferenceNum,omitempty" xml:"ReferenceNum,omitempty"`
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) GetDescription() *string {
	return s.Description
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) GetId() *int64 {
	return s.Id
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) GetName() *string {
	return s.Name
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) GetReferenceNum() *int32 {
	return s.ReferenceNum
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetDescription(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Description = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetId(v int64) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Id = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetName(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Name = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetReferenceNum(v int32) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.ReferenceNum = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) Validate() error {
	return dara.Validate(s)
}
