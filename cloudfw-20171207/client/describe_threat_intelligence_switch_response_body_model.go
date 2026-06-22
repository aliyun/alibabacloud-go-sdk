// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatIntelligenceSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryList(v []*DescribeThreatIntelligenceSwitchResponseBodyCategoryList) *DescribeThreatIntelligenceSwitchResponseBody
	GetCategoryList() []*DescribeThreatIntelligenceSwitchResponseBodyCategoryList
	SetRequestId(v string) *DescribeThreatIntelligenceSwitchResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeThreatIntelligenceSwitchResponseBody
	GetTotalCount() *int64
}

type DescribeThreatIntelligenceSwitchResponseBody struct {
	// The list of threat intelligence categories.
	CategoryList []*DescribeThreatIntelligenceSwitchResponseBodyCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6B8E0379-2629-59A1-B811-96F3E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeThreatIntelligenceSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatIntelligenceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) GetCategoryList() []*DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	return s.CategoryList
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) SetCategoryList(v []*DescribeThreatIntelligenceSwitchResponseBodyCategoryList) *DescribeThreatIntelligenceSwitchResponseBody {
	s.CategoryList = v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) SetRequestId(v string) *DescribeThreatIntelligenceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) SetTotalCount(v int64) *DescribeThreatIntelligenceSwitchResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBody) Validate() error {
	if s.CategoryList != nil {
		for _, item := range s.CategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeThreatIntelligenceSwitchResponseBodyCategoryList struct {
	// The action of the rule. Valid values:
	//
	//  **alert**: monitor mode.
	//
	//  **drop**: Block Mode.
	//
	// example:
	//
	// alert
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The category description.
	//
	// example:
	//
	// Tor Export Malicious IP
	CategoryDescribe *string `json:"CategoryDescribe,omitempty" xml:"CategoryDescribe,omitempty"`
	// The ID of the threat intelligence category.
	//
	// example:
	//
	// 3000037
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// Tor Export Malicious IP
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The ID of the parent threat intelligence category.
	//
	// example:
	//
	// 123
	CategoryParentId *string `json:"CategoryParentId,omitempty" xml:"CategoryParentId,omitempty"`
	// The enabling status. Valid values:
	//
	//  **1**: enabled.
	//
	//  **0**: disabled.
	//
	// example:
	//
	// 1
	EnableStatus *int64 `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
}

func (s DescribeThreatIntelligenceSwitchResponseBodyCategoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetAction() *string {
	return s.Action
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetCategoryDescribe() *string {
	return s.CategoryDescribe
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetCategoryParentId() *string {
	return s.CategoryParentId
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) GetEnableStatus() *int64 {
	return s.EnableStatus
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetAction(v string) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.Action = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetCategoryDescribe(v string) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.CategoryDescribe = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetCategoryId(v string) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.CategoryId = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetCategoryName(v string) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.CategoryName = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetCategoryParentId(v string) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.CategoryParentId = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) SetEnableStatus(v int64) *DescribeThreatIntelligenceSwitchResponseBodyCategoryList {
	s.EnableStatus = &v
	return s
}

func (s *DescribeThreatIntelligenceSwitchResponseBodyCategoryList) Validate() error {
	return dara.Validate(s)
}
