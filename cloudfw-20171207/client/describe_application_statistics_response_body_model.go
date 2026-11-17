// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryList(v []*DescribeApplicationStatisticsResponseBodyCategoryList) *DescribeApplicationStatisticsResponseBody
	GetCategoryList() []*DescribeApplicationStatisticsResponseBodyCategoryList
	SetRequestId(v string) *DescribeApplicationStatisticsResponseBody
	GetRequestId() *string
	SetRiskList(v []*DescribeApplicationStatisticsResponseBodyRiskList) *DescribeApplicationStatisticsResponseBody
	GetRiskList() []*DescribeApplicationStatisticsResponseBodyRiskList
	SetTechnologyList(v []*DescribeApplicationStatisticsResponseBodyTechnologyList) *DescribeApplicationStatisticsResponseBody
	GetTechnologyList() []*DescribeApplicationStatisticsResponseBodyTechnologyList
}

type DescribeApplicationStatisticsResponseBody struct {
	// The list of application counts grouped by application category.
	CategoryList []*DescribeApplicationStatisticsResponseBodyCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of application counts grouped by risk level.
	RiskList []*DescribeApplicationStatisticsResponseBodyRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
	// The list of application counts grouped by technology type.
	TechnologyList []*DescribeApplicationStatisticsResponseBodyTechnologyList `json:"TechnologyList,omitempty" xml:"TechnologyList,omitempty" type:"Repeated"`
}

func (s DescribeApplicationStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsResponseBody) GetCategoryList() []*DescribeApplicationStatisticsResponseBodyCategoryList {
	return s.CategoryList
}

func (s *DescribeApplicationStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationStatisticsResponseBody) GetRiskList() []*DescribeApplicationStatisticsResponseBodyRiskList {
	return s.RiskList
}

func (s *DescribeApplicationStatisticsResponseBody) GetTechnologyList() []*DescribeApplicationStatisticsResponseBodyTechnologyList {
	return s.TechnologyList
}

func (s *DescribeApplicationStatisticsResponseBody) SetCategoryList(v []*DescribeApplicationStatisticsResponseBodyCategoryList) *DescribeApplicationStatisticsResponseBody {
	s.CategoryList = v
	return s
}

func (s *DescribeApplicationStatisticsResponseBody) SetRequestId(v string) *DescribeApplicationStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBody) SetRiskList(v []*DescribeApplicationStatisticsResponseBodyRiskList) *DescribeApplicationStatisticsResponseBody {
	s.RiskList = v
	return s
}

func (s *DescribeApplicationStatisticsResponseBody) SetTechnologyList(v []*DescribeApplicationStatisticsResponseBodyTechnologyList) *DescribeApplicationStatisticsResponseBody {
	s.TechnologyList = v
	return s
}

func (s *DescribeApplicationStatisticsResponseBody) Validate() error {
	if s.CategoryList != nil {
		for _, item := range s.CategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TechnologyList != nil {
		for _, item := range s.TechnologyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationStatisticsResponseBodyCategoryList struct {
	// The number of applications in the application category.
	//
	// example:
	//
	// 12
	ApplicationCount *int64 `json:"ApplicationCount,omitempty" xml:"ApplicationCount,omitempty"`
	// The name of the application category. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// Office collaboration
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s DescribeApplicationStatisticsResponseBodyCategoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsResponseBodyCategoryList) GetApplicationCount() *int64 {
	return s.ApplicationCount
}

func (s *DescribeApplicationStatisticsResponseBodyCategoryList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeApplicationStatisticsResponseBodyCategoryList) SetApplicationCount(v int64) *DescribeApplicationStatisticsResponseBodyCategoryList {
	s.ApplicationCount = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyCategoryList) SetCategoryName(v string) *DescribeApplicationStatisticsResponseBodyCategoryList {
	s.CategoryName = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyCategoryList) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationStatisticsResponseBodyRiskList struct {
	// The number of applications at the risk level.
	//
	// example:
	//
	// 20
	ApplicationCount *int64 `json:"ApplicationCount,omitempty" xml:"ApplicationCount,omitempty"`
	// The risk level. Valid values:
	//
	// - **Information**: Informational.
	//
	// - **Low**: Low.
	//
	// - **Medium**: Medium.
	//
	// - **High**: High.
	//
	// - **Critical**: Critical.
	//
	// example:
	//
	// Low
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
}

func (s DescribeApplicationStatisticsResponseBodyRiskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsResponseBodyRiskList) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsResponseBodyRiskList) GetApplicationCount() *int64 {
	return s.ApplicationCount
}

func (s *DescribeApplicationStatisticsResponseBodyRiskList) GetRisk() *string {
	return s.Risk
}

func (s *DescribeApplicationStatisticsResponseBodyRiskList) SetApplicationCount(v int64) *DescribeApplicationStatisticsResponseBodyRiskList {
	s.ApplicationCount = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyRiskList) SetRisk(v string) *DescribeApplicationStatisticsResponseBodyRiskList {
	s.Risk = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyRiskList) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationStatisticsResponseBodyTechnologyList struct {
	// The number of applications of the technology type.
	//
	// example:
	//
	// 8
	ApplicationCount *int64 `json:"ApplicationCount,omitempty" xml:"ApplicationCount,omitempty"`
	// The name of the technology type. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// Browser-based
	Technology *string `json:"Technology,omitempty" xml:"Technology,omitempty"`
}

func (s DescribeApplicationStatisticsResponseBodyTechnologyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsResponseBodyTechnologyList) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsResponseBodyTechnologyList) GetApplicationCount() *int64 {
	return s.ApplicationCount
}

func (s *DescribeApplicationStatisticsResponseBodyTechnologyList) GetTechnology() *string {
	return s.Technology
}

func (s *DescribeApplicationStatisticsResponseBodyTechnologyList) SetApplicationCount(v int64) *DescribeApplicationStatisticsResponseBodyTechnologyList {
	s.ApplicationCount = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyTechnologyList) SetTechnology(v string) *DescribeApplicationStatisticsResponseBodyTechnologyList {
	s.Technology = &v
	return s
}

func (s *DescribeApplicationStatisticsResponseBodyTechnologyList) Validate() error {
	return dara.Validate(s)
}
