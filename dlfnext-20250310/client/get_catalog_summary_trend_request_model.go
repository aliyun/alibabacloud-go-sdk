// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogSummaryTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetCatalogSummaryTrendRequest
	GetEndDate() *string
	SetStartDate(v string) *GetCatalogSummaryTrendRequest
	GetStartDate() *string
}

type GetCatalogSummaryTrendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-01
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s GetCatalogSummaryTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogSummaryTrendRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryTrendRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetCatalogSummaryTrendRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetCatalogSummaryTrendRequest) SetEndDate(v string) *GetCatalogSummaryTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetCatalogSummaryTrendRequest) SetStartDate(v string) *GetCatalogSummaryTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetCatalogSummaryTrendRequest) Validate() error {
	return dara.Validate(s)
}
