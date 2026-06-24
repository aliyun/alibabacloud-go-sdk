// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeCategoryTemplateRuleListResponseBodyItems) *DescribeCategoryTemplateRuleListResponseBody
	GetItems() []*DescribeCategoryTemplateRuleListResponseBodyItems
	SetPageSize(v int32) *DescribeCategoryTemplateRuleListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCategoryTemplateRuleListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCategoryTemplateRuleListResponseBody
	GetTotalCount() *int32
}

type DescribeCategoryTemplateRuleListResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of template rules.
	Items []*DescribeCategoryTemplateRuleListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of template rules returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 136082B3-B21F-5E9D-B68E-991FFD205D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of rules in the template.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCategoryTemplateRuleListResponseBody) GetItems() []*DescribeCategoryTemplateRuleListResponseBodyItems {
	return s.Items
}

func (s *DescribeCategoryTemplateRuleListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCategoryTemplateRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCategoryTemplateRuleListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetItems(v []*DescribeCategoryTemplateRuleListResponseBodyItems) *DescribeCategoryTemplateRuleListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetPageSize(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetRequestId(v string) *DescribeCategoryTemplateRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetTotalCount(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCategoryTemplateRuleListResponseBodyItems struct {
	// The description of the rule.
	//
	// example:
	//
	// Template rule for identifying ID card numbers
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the template rule.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// A comma-separated list of IDs of the associated atomic models.
	//
	// example:
	//
	// 1001,1002
	IdentificationRuleIds *string `json:"IdentificationRuleIds,omitempty" xml:"IdentificationRuleIds,omitempty"`
	// The scope of data that the template rule scans. This parameter is a string converted from a JSON array. Each element in the JSON array represents a data scanning scope and contains the following fields:
	//
	// - **Asset**: A string that indicates the asset type. Valid values include RDS, DRDS, PolarDB, OTS, ADB, OceanBase, and ODPS.
	//
	// - **Content**: The specific scope of the asset to scan. This is an array of objects, where each object contains the following fields:
	//
	//   - **Range**: A string that indicates the matching range. Valid values include instance, database, table, column, project (for MaxCompute assets only), bucket (for OSS assets only), and object (for OSS assets only).
	//
	//   - **Operator**: A string that indicates the matching condition. Valid values include equals, regex (regular expression), prefix, and suffix.
	//
	//   - **Value**: A string that indicates the content to match.
	//
	// example:
	//
	// [{"Asset":"RDS","Content":[{"Range":"database","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"project","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]}]
	IdentificationScope *string `json:"IdentificationScope,omitempty" xml:"IdentificationScope,omitempty"`
	// The name of the template rule.
	//
	// example:
	//
	// ID card
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The risk level of the template rule. The value ranges from **1*	- to **11**. Valid values:
	//
	// - **1**: No risk.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// - **6**: S5.
	//
	// - **7**: S6.
	//
	// - **8**: S7.
	//
	// - **9**: S8.
	//
	// - **10**: S9.
	//
	// - **11**: S10.
	//
	// - **null**: Indicates all risk levels, including No risk, S1, S2, S3, S4, S5, S6, S7, S8, S9, and S10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the template rule. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// - **null**: Represents all statuses, including enabled and disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetIdentificationRuleIds() *string {
	return s.IdentificationRuleIds
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetIdentificationScope() *string {
	return s.IdentificationScope
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetDescription(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetId(v int64) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetIdentificationRuleIds(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.IdentificationRuleIds = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetIdentificationScope(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.IdentificationScope = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetName(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetRiskLevelId(v int64) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetStatus(v int32) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
