// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanResultsByEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineState(v string) *DescribeScanResultsByEngineRequest
	GetBaselineState() *string
	SetLang(v string) *DescribeScanResultsByEngineRequest
	GetLang() *string
	SetMaxResults(v int64) *DescribeScanResultsByEngineRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeScanResultsByEngineRequest
	GetNextToken() *string
	SetPackageName(v string) *DescribeScanResultsByEngineRequest
	GetPackageName() *string
}

type DescribeScanResultsByEngineRequest struct {
	// Filters results by incremental scan baseline status. Valid values: new, unchanged, absent, updated.
	//
	// example:
	//
	// unchanged
	BaselineState *string `json:"baselineState,omitempty" xml:"baselineState,omitempty"`
	// The language. Valid values:
	//
	// 	- zh: Chinese (default).
	//
	// 	- en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not pass nextToken or pass an empty string for the first page. To retrieve the next page, pass the nextToken value from the previous response without any modification. When the nextToken in the response is empty, you have reached the last page.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Performs a fuzzy match by component name. This parameter takes effect only when engine is set to sca.
	//
	// example:
	//
	// app10000006824_fulljoin_61b2c_global_ad_budget_ctrl_v1.5.jar
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
}

func (s DescribeScanResultsByEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineRequest) GetBaselineState() *string {
	return s.BaselineState
}

func (s *DescribeScanResultsByEngineRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeScanResultsByEngineRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeScanResultsByEngineRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScanResultsByEngineRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *DescribeScanResultsByEngineRequest) SetBaselineState(v string) *DescribeScanResultsByEngineRequest {
	s.BaselineState = &v
	return s
}

func (s *DescribeScanResultsByEngineRequest) SetLang(v string) *DescribeScanResultsByEngineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScanResultsByEngineRequest) SetMaxResults(v int64) *DescribeScanResultsByEngineRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScanResultsByEngineRequest) SetNextToken(v string) *DescribeScanResultsByEngineRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeScanResultsByEngineRequest) SetPackageName(v string) *DescribeScanResultsByEngineRequest {
	s.PackageName = &v
	return s
}

func (s *DescribeScanResultsByEngineRequest) Validate() error {
	return dara.Validate(s)
}
