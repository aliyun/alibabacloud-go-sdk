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
	// The page number. Default value: **1**. Pages start from page 1.
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
	// The number of entries per page. Maximum value: 1000. Default value: 20. If you leave this parameter empty, 20 data entries are returned per page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 183.63.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain the ID.
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
