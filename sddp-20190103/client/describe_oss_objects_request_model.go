// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOssObjectsRequest
	GetCurrentPage() *int32
	SetFileCategoryCode(v int64) *DescribeOssObjectsRequest
	GetFileCategoryCode() *int64
	SetInstanceId(v string) *DescribeOssObjectsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeOssObjectsRequest
	GetLang() *string
	SetLastScanTimeEnd(v int64) *DescribeOssObjectsRequest
	GetLastScanTimeEnd() *int64
	SetLastScanTimeStart(v int64) *DescribeOssObjectsRequest
	GetLastScanTimeStart() *int64
	SetMarker(v int64) *DescribeOssObjectsRequest
	GetMarker() *int64
	SetName(v string) *DescribeOssObjectsRequest
	GetName() *string
	SetPageSize(v int32) *DescribeOssObjectsRequest
	GetPageSize() *int32
	SetRiskLevelId(v int32) *DescribeOssObjectsRequest
	GetRiskLevelId() *int32
	SetRuleId(v int64) *DescribeOssObjectsRequest
	GetRuleId() *int64
	SetServiceRegionId(v string) *DescribeOssObjectsRequest
	GetServiceRegionId() *string
	SetTemplateId(v int64) *DescribeOssObjectsRequest
	GetTemplateId() *int64
}

type DescribeOssObjectsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The code of the file type.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// The ID of the instance to which the OSS object belongs.
	//
	// > You can call the **DescribeInstances*	- operation to query the instance ID.
	//
	// example:
	//
	// ins-2222
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The end time of the last scan. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	LastScanTimeEnd *int64 `json:"LastScanTimeEnd,omitempty" xml:"LastScanTimeEnd,omitempty"`
	// The start time of the last scan. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	LastScanTimeStart *int64 `json:"LastScanTimeStart,omitempty" xml:"LastScanTimeStart,omitempty"`
	// When you query data by page, use the `Marker` parameter to query the data that follows the `Marker` value.
	//
	// example:
	//
	// 1754786235714378752
	Marker *int64 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sensitivity level of the OSS object. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int32 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the OSS object hits.
	//
	// > You can call the **DescribeRules*	- operation to query the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 1222
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeOssObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOssObjectsRequest) GetFileCategoryCode() *int64 {
	return s.FileCategoryCode
}

func (s *DescribeOssObjectsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeOssObjectsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssObjectsRequest) GetLastScanTimeEnd() *int64 {
	return s.LastScanTimeEnd
}

func (s *DescribeOssObjectsRequest) GetLastScanTimeStart() *int64 {
	return s.LastScanTimeStart
}

func (s *DescribeOssObjectsRequest) GetMarker() *int64 {
	return s.Marker
}

func (s *DescribeOssObjectsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeOssObjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOssObjectsRequest) GetRiskLevelId() *int32 {
	return s.RiskLevelId
}

func (s *DescribeOssObjectsRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeOssObjectsRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeOssObjectsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeOssObjectsRequest) SetCurrentPage(v int32) *DescribeOssObjectsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetFileCategoryCode(v int64) *DescribeOssObjectsRequest {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetInstanceId(v string) *DescribeOssObjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLang(v string) *DescribeOssObjectsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeEnd(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeEnd = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeStart(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeStart = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetMarker(v int64) *DescribeOssObjectsRequest {
	s.Marker = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetName(v string) *DescribeOssObjectsRequest {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetPageSize(v int32) *DescribeOssObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetRiskLevelId(v int32) *DescribeOssObjectsRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetRuleId(v int64) *DescribeOssObjectsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetServiceRegionId(v string) *DescribeOssObjectsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetTemplateId(v int64) *DescribeOssObjectsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeOssObjectsRequest) Validate() error {
	return dara.Validate(s)
}
