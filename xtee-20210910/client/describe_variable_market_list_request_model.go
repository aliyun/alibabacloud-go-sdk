// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableMarketListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableMarketListRequest
	GetLang() *string
	SetChargingMode(v string) *DescribeVariableMarketListRequest
	GetChargingMode() *string
	SetCurrentPage(v int32) *DescribeVariableMarketListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVariableMarketListRequest
	GetPageSize() *int32
	SetPaging(v string) *DescribeVariableMarketListRequest
	GetPaging() *string
	SetQueryContent(v string) *DescribeVariableMarketListRequest
	GetQueryContent() *string
	SetRegId(v string) *DescribeVariableMarketListRequest
	GetRegId() *string
	SetScenesStr(v string) *DescribeVariableMarketListRequest
	GetScenesStr() *string
	SetSource(v string) *DescribeVariableMarketListRequest
	GetSource() *string
	SetTitle(v string) *DescribeVariableMarketListRequest
	GetTitle() *string
}

type DescribeVariableMarketListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Charging mode
	//
	// example:
	//
	// FREE
	ChargingMode *string `json:"chargingMode,omitempty" xml:"chargingMode,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Paging indicator, default is true.
	//
	// example:
	//
	// true
	Paging *string `json:"paging,omitempty" xml:"paging,omitempty"`
	// Query content, supports fuzzy search.
	//
	//  Title/Description
	//
	// example:
	//
	// 手机号
	QueryContent *string `json:"queryContent,omitempty" xml:"queryContent,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Scenario
	//
	// example:
	//
	// [\\"coupon_abuse_detection\\"]
	ScenesStr *string `json:"scenesStr,omitempty" xml:"scenesStr,omitempty"`
	// Source
	//
	// example:
	//
	// SAF
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Title.
	//
	// example:
	//
	// 设备风险识别_标签
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeVariableMarketListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableMarketListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableMarketListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableMarketListRequest) GetChargingMode() *string {
	return s.ChargingMode
}

func (s *DescribeVariableMarketListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVariableMarketListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVariableMarketListRequest) GetPaging() *string {
	return s.Paging
}

func (s *DescribeVariableMarketListRequest) GetQueryContent() *string {
	return s.QueryContent
}

func (s *DescribeVariableMarketListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableMarketListRequest) GetScenesStr() *string {
	return s.ScenesStr
}

func (s *DescribeVariableMarketListRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeVariableMarketListRequest) GetTitle() *string {
	return s.Title
}

func (s *DescribeVariableMarketListRequest) SetLang(v string) *DescribeVariableMarketListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetChargingMode(v string) *DescribeVariableMarketListRequest {
	s.ChargingMode = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetCurrentPage(v int32) *DescribeVariableMarketListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetPageSize(v int32) *DescribeVariableMarketListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetPaging(v string) *DescribeVariableMarketListRequest {
	s.Paging = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetQueryContent(v string) *DescribeVariableMarketListRequest {
	s.QueryContent = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetRegId(v string) *DescribeVariableMarketListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetScenesStr(v string) *DescribeVariableMarketListRequest {
	s.ScenesStr = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetSource(v string) *DescribeVariableMarketListRequest {
	s.Source = &v
	return s
}

func (s *DescribeVariableMarketListRequest) SetTitle(v string) *DescribeVariableMarketListRequest {
	s.Title = &v
	return s
}

func (s *DescribeVariableMarketListRequest) Validate() error {
	return dara.Validate(s)
}
