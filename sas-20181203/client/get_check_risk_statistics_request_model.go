// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRiskStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetCheckRiskStatisticsRequest
	GetCurrentPage() *int32
	SetLang(v string) *GetCheckRiskStatisticsRequest
	GetLang() *string
	SetPageSize(v int32) *GetCheckRiskStatisticsRequest
	GetPageSize() *int32
}

type GetCheckRiskStatisticsRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCheckRiskStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetCheckRiskStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCheckRiskStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCheckRiskStatisticsRequest) SetCurrentPage(v int32) *GetCheckRiskStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetCheckRiskStatisticsRequest) SetLang(v string) *GetCheckRiskStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *GetCheckRiskStatisticsRequest) SetPageSize(v int32) *GetCheckRiskStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetCheckRiskStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
