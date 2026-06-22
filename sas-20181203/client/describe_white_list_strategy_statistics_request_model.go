// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWhiteListStrategyStatisticsRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeWhiteListStrategyStatisticsRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeWhiteListStrategyStatisticsRequest
	GetPageSize() *int32
	SetSourceIp(v string) *DescribeWhiteListStrategyStatisticsRequest
	GetSourceIp() *string
	SetStrategyIds(v string) *DescribeWhiteListStrategyStatisticsRequest
	GetStrategyIds() *string
}

type DescribeWhiteListStrategyStatisticsRequest struct {
	// The page number of the page to return. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 1000. Default value: 20. If you leave this parameter empty, 20 entries are returned.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains this value.
	//
	// example:
	//
	// 183.63.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The policy ID.
	//
	// >You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain this parameter.
	//
	// example:
	//
	// 3645
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
}

func (s DescribeWhiteListStrategyStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListStrategyStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListStrategyStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListStrategyStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListStrategyStatisticsRequest) GetStrategyIds() *string {
	return s.StrategyIds
}

func (s *DescribeWhiteListStrategyStatisticsRequest) SetCurrentPage(v int32) *DescribeWhiteListStrategyStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsRequest) SetLang(v string) *DescribeWhiteListStrategyStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsRequest) SetPageSize(v int32) *DescribeWhiteListStrategyStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsRequest) SetSourceIp(v string) *DescribeWhiteListStrategyStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsRequest) SetStrategyIds(v string) *DescribeWhiteListStrategyStatisticsRequest {
	s.StrategyIds = &v
	return s
}

func (s *DescribeWhiteListStrategyStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
