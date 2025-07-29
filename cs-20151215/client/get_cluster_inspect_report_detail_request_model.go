// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterInspectReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetClusterInspectReportDetailRequest
	GetCategory() *string
	SetEnableFilter(v bool) *GetClusterInspectReportDetailRequest
	GetEnableFilter() *bool
	SetLanguage(v string) *GetClusterInspectReportDetailRequest
	GetLanguage() *string
	SetLevel(v string) *GetClusterInspectReportDetailRequest
	GetLevel() *string
	SetMaxResults(v int32) *GetClusterInspectReportDetailRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetClusterInspectReportDetailRequest
	GetNextToken() *string
	SetTargetType(v string) *GetClusterInspectReportDetailRequest
	GetTargetType() *string
}

type GetClusterInspectReportDetailRequest struct {
	// The category of the inspection item. Valid values:
	//
	// 	- security: Security compliance
	//
	// 	- performance: Performance efficiency
	//
	// 	- stability: Business stability
	//
	// 	- limitation: Service limits
	//
	// 	- cost: Cost optimization
	//
	// example:
	//
	// cost
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Inspection results filtering. If this parameter is set to true, only abnormal inspection items are returned.
	//
	// example:
	//
	// true
	EnableFilter *bool `json:"enableFilter,omitempty" xml:"enableFilter,omitempty"`
	// The query language.
	//
	// 	- zh_CN
	//
	// 	- en_US
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The level of the inspection item. Valid values:
	//
	// 	- advice: Suggestions
	//
	// 	- warning: Low severity
	//
	// 	- error: Medium severity
	//
	// 	- critical: High severity
	//
	// example:
	//
	// error
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The maximum number of entries per page. Maximum value: 50.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that is used to display the returned tags on multiple pages.
	//
	// example:
	//
	// bb41e2bcb47b747389dce6c0d2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The type of the inspection object. Only items that meet the targetType parameter are returned.
	//
	// example:
	//
	// NodePool
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s GetClusterInspectReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectReportDetailRequest) GoString() string {
	return s.String()
}

func (s *GetClusterInspectReportDetailRequest) GetCategory() *string {
	return s.Category
}

func (s *GetClusterInspectReportDetailRequest) GetEnableFilter() *bool {
	return s.EnableFilter
}

func (s *GetClusterInspectReportDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetClusterInspectReportDetailRequest) GetLevel() *string {
	return s.Level
}

func (s *GetClusterInspectReportDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetClusterInspectReportDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetClusterInspectReportDetailRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GetClusterInspectReportDetailRequest) SetCategory(v string) *GetClusterInspectReportDetailRequest {
	s.Category = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetEnableFilter(v bool) *GetClusterInspectReportDetailRequest {
	s.EnableFilter = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetLanguage(v string) *GetClusterInspectReportDetailRequest {
	s.Language = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetLevel(v string) *GetClusterInspectReportDetailRequest {
	s.Level = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetMaxResults(v int32) *GetClusterInspectReportDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetNextToken(v string) *GetClusterInspectReportDetailRequest {
	s.NextToken = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) SetTargetType(v string) *GetClusterInspectReportDetailRequest {
	s.TargetType = &v
	return s
}

func (s *GetClusterInspectReportDetailRequest) Validate() error {
	return dara.Validate(s)
}
