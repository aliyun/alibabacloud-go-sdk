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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of rules.
	Items []*DescribeCategoryTemplateRuleListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
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
	// Rule for identifying ID card numbers
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the rule.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of sensitive data types. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 1001,1002
	IdentificationRuleIds *string `json:"IdentificationRuleIds,omitempty" xml:"IdentificationRuleIds,omitempty"`
	// The scan scope of the rule. The value is a JSON array of the STRING type. Each element in a JSON array indicates a scan scope that contains the following fields:
	//
	// 	- **Asset**: the data asset type. Valid values: RDS, DRDS, PolarDB, OTS, ADB, and OceanBase. The value is of the STRING type.
	//
	// 	- **Content**: the scan scope. The value is of the STRING type. Each element in a JSON array indicates a scan scope that contains the following fields:
	//
	//     	- **Range**: the matching condition. Valid values: Instance, database, table, column, project, bucket, and object. The value project is valid only for data assets in MaxCompute. The values bucket and object are valid only for data assets in Object Storage Service (OSS). The value of this parameter is of the STRING type.
	//
	//     	- **Operator**: the operator. Valid values: equals, regex, prefix, and suffix. The operator equals indicates a full match. The operator regex indicates a match by regular expression. The operator prefix indicates a match by prefix. The operator suffix indicates a match by suffix.
	//
	//     	- **Value**: the matching content. The value is of the STRING type.
	//
	// example:
	//
	// [{"Asset":"RDS","Content":[{"Range":"database","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"project","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]}]
	IdentificationScope *string `json:"IdentificationScope,omitempty" xml:"IdentificationScope,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ID card number
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sensitivity level of the data that is not compliant with the rule. Valid values: **1*	- to **11**.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: indicates the S1 sensitivity level.
	//
	// 	- **3**: indicates the S2 sensitivity level.
	//
	// 	- **4**: indicates the S3 sensitivity level.
	//
	// 	- **5**: indicates the S4 sensitivity level.
	//
	// 	- **6**: indicates the S5 sensitivity level.
	//
	// 	- **7**: indicates the S6 sensitivity level.
	//
	// 	- **8**: indicates the S7 sensitivity level.
	//
	// 	- **9**: indicates the S8 sensitivity level.
	//
	// 	- **10**: indicates the S9 sensitivity level.
	//
	// 	- **11**: indicates the S10 sensitivity level.
	//
	// 	- **null**: indicates all preceding sensitivity levels.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **null**: all states
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
