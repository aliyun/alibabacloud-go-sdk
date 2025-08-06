// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQoeListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribePlayQoeListShrinkRequest
	GetAppName() *string
	SetBeginTs(v int64) *DescribePlayQoeListShrinkRequest
	GetBeginTs() *int64
	SetDefinition(v string) *DescribePlayQoeListShrinkRequest
	GetDefinition() *string
	SetEndTs(v int64) *DescribePlayQoeListShrinkRequest
	GetEndTs() *int64
	SetItemConfigs(v string) *DescribePlayQoeListShrinkRequest
	GetItemConfigs() *string
	SetMetricTypesShrink(v string) *DescribePlayQoeListShrinkRequest
	GetMetricTypesShrink() *string
	SetNetwork(v string) *DescribePlayQoeListShrinkRequest
	GetNetwork() *string
	SetOrderName(v string) *DescribePlayQoeListShrinkRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribePlayQoeListShrinkRequest
	GetOrderType() *string
	SetOs(v string) *DescribePlayQoeListShrinkRequest
	GetOs() *string
	SetPageNo(v int64) *DescribePlayQoeListShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQoeListShrinkRequest
	GetPageSize() *int64
	SetTerminalType(v string) *DescribePlayQoeListShrinkRequest
	GetTerminalType() *string
}

type DescribePlayQoeListShrinkRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	BeginTs    *int64  `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	EndTs             *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	ItemConfigs       *string `json:"ItemConfigs,omitempty" xml:"ItemConfigs,omitempty"`
	MetricTypesShrink *string `json:"MetricTypes,omitempty" xml:"MetricTypes,omitempty"`
	Network           *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OrderName         *string `json:"OrderName,omitempty" xml:"OrderName,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Os                *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribePlayQoeListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQoeListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayQoeListShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayQoeListShrinkRequest) GetBeginTs() *int64 {
	return s.BeginTs
}

func (s *DescribePlayQoeListShrinkRequest) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayQoeListShrinkRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribePlayQoeListShrinkRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribePlayQoeListShrinkRequest) GetMetricTypesShrink() *string {
	return s.MetricTypesShrink
}

func (s *DescribePlayQoeListShrinkRequest) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayQoeListShrinkRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribePlayQoeListShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePlayQoeListShrinkRequest) GetOs() *string {
	return s.Os
}

func (s *DescribePlayQoeListShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQoeListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQoeListShrinkRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayQoeListShrinkRequest) SetAppName(v string) *DescribePlayQoeListShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetBeginTs(v int64) *DescribePlayQoeListShrinkRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetDefinition(v string) *DescribePlayQoeListShrinkRequest {
	s.Definition = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetEndTs(v int64) *DescribePlayQoeListShrinkRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetItemConfigs(v string) *DescribePlayQoeListShrinkRequest {
	s.ItemConfigs = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetMetricTypesShrink(v string) *DescribePlayQoeListShrinkRequest {
	s.MetricTypesShrink = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetNetwork(v string) *DescribePlayQoeListShrinkRequest {
	s.Network = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetOrderName(v string) *DescribePlayQoeListShrinkRequest {
	s.OrderName = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetOrderType(v string) *DescribePlayQoeListShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetOs(v string) *DescribePlayQoeListShrinkRequest {
	s.Os = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetPageNo(v int64) *DescribePlayQoeListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetPageSize(v int64) *DescribePlayQoeListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) SetTerminalType(v string) *DescribePlayQoeListShrinkRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayQoeListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
