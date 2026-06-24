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
	// A list of risk levels.
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
	// The description of the risk level. You can customize the description.
	//
	// The following list describes the mappings between risk level names and their default descriptions:
	//
	// - **NA**: No sensitive data is detected
	//
	// - **S1**: Level-1 sensitive data
	//
	// - **S2**: Level-2 sensitive data
	//
	// - **S3**: Level-3 sensitive data
	//
	// - **S4**: Level-4 sensitive data
	//
	// - **S5**: Level-5 sensitive data
	//
	// - **S6**: Level-6 sensitive data
	//
	// - **S7**: Level-7 sensitive data
	//
	// - **S8**: Level-8 sensitive data
	//
	// - **S9**: Level-9 sensitive data
	//
	// - **S10**: Level-10 sensitive data
	//
	// example:
	//
	// S3
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the risk level. Valid values: 1 to 11. Each risk level ID corresponds to a risk level name. For example, the risk level ID 2 corresponds to the risk level S1.
	//
	// For more information about the mappings, see the description of the Name parameter.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the risk level for the sensitive data. The maximum risk level is S10 and varies based on the data classification template. The maximum risk level is S4 for the **built-in data security classification template for Alibaba and Ant Group**, and S5 for the **built-in data classification template for the finance industry (with reference to JR/T 0197-2020 Financial Data Security - Guidelines for Data Security Classification)*	- and the **built-in data classification standard for the energy industry**. If you use a copied template, the maximum risk level is S10.
	//
	// The following list describes the mappings between risk level names and IDs:
	//
	// - **NA**: 1
	//
	// - **S1**: 2
	//
	// - **S2**: 3
	//
	// - **S3**: 4
	//
	// - **S4**: 5
	//
	// - **S5**: 6
	//
	// - **S6**: 7
	//
	// - **S7**: 8
	//
	// - **S8**: 9
	//
	// - **S9**: 10
	//
	// - **S10**: 11
	//
	// example:
	//
	// S1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of times the risk level is referenced in the template. The default value is 0.
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
