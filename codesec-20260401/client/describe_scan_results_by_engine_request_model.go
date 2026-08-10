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
	// 按 baseline_state 过滤 findings。值为 new/unchanged/absent/updated 之一。
	BaselineState *string `json:"baselineState,omitempty" xml:"baselineState,omitempty"`
	Lang          *string `json:"lang,omitempty" xml:"lang,omitempty"`
	MaxResults    *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PackageName   *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
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
