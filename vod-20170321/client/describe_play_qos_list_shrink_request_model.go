// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQosListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribePlayQosListShrinkRequest
	GetAppName() *string
	SetBeginTs(v int64) *DescribePlayQosListShrinkRequest
	GetBeginTs() *int64
	SetDefinition(v string) *DescribePlayQosListShrinkRequest
	GetDefinition() *string
	SetEndTs(v int64) *DescribePlayQosListShrinkRequest
	GetEndTs() *int64
	SetItemConfigs(v string) *DescribePlayQosListShrinkRequest
	GetItemConfigs() *string
	SetMetricTypesShrink(v string) *DescribePlayQosListShrinkRequest
	GetMetricTypesShrink() *string
	SetNetwork(v string) *DescribePlayQosListShrinkRequest
	GetNetwork() *string
	SetOrderName(v string) *DescribePlayQosListShrinkRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribePlayQosListShrinkRequest
	GetOrderType() *string
	SetOs(v string) *DescribePlayQosListShrinkRequest
	GetOs() *string
	SetPageNo(v int64) *DescribePlayQosListShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQosListShrinkRequest
	GetPageSize() *int64
	SetTerminalType(v string) *DescribePlayQosListShrinkRequest
	GetTerminalType() *string
}

type DescribePlayQosListShrinkRequest struct {
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

func (s DescribePlayQosListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQosListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayQosListShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayQosListShrinkRequest) GetBeginTs() *int64 {
	return s.BeginTs
}

func (s *DescribePlayQosListShrinkRequest) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayQosListShrinkRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribePlayQosListShrinkRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribePlayQosListShrinkRequest) GetMetricTypesShrink() *string {
	return s.MetricTypesShrink
}

func (s *DescribePlayQosListShrinkRequest) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayQosListShrinkRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribePlayQosListShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePlayQosListShrinkRequest) GetOs() *string {
	return s.Os
}

func (s *DescribePlayQosListShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQosListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQosListShrinkRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayQosListShrinkRequest) SetAppName(v string) *DescribePlayQosListShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetBeginTs(v int64) *DescribePlayQosListShrinkRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetDefinition(v string) *DescribePlayQosListShrinkRequest {
	s.Definition = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetEndTs(v int64) *DescribePlayQosListShrinkRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetItemConfigs(v string) *DescribePlayQosListShrinkRequest {
	s.ItemConfigs = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetMetricTypesShrink(v string) *DescribePlayQosListShrinkRequest {
	s.MetricTypesShrink = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetNetwork(v string) *DescribePlayQosListShrinkRequest {
	s.Network = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetOrderName(v string) *DescribePlayQosListShrinkRequest {
	s.OrderName = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetOrderType(v string) *DescribePlayQosListShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetOs(v string) *DescribePlayQosListShrinkRequest {
	s.Os = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetPageNo(v int64) *DescribePlayQosListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetPageSize(v int64) *DescribePlayQosListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) SetTerminalType(v string) *DescribePlayQosListShrinkRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayQosListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
