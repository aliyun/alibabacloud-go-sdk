// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQoeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribePlayQoeListRequest
	GetAppName() *string
	SetBeginTs(v int64) *DescribePlayQoeListRequest
	GetBeginTs() *int64
	SetDefinition(v string) *DescribePlayQoeListRequest
	GetDefinition() *string
	SetEndTs(v int64) *DescribePlayQoeListRequest
	GetEndTs() *int64
	SetItemConfigs(v string) *DescribePlayQoeListRequest
	GetItemConfigs() *string
	SetMetricTypes(v []*string) *DescribePlayQoeListRequest
	GetMetricTypes() []*string
	SetNetwork(v string) *DescribePlayQoeListRequest
	GetNetwork() *string
	SetOrderName(v string) *DescribePlayQoeListRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribePlayQoeListRequest
	GetOrderType() *string
	SetOs(v string) *DescribePlayQoeListRequest
	GetOs() *string
	SetPageNo(v int64) *DescribePlayQoeListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQoeListRequest
	GetPageSize() *int64
	SetTerminalType(v string) *DescribePlayQoeListRequest
	GetTerminalType() *string
}

type DescribePlayQoeListRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	BeginTs    *int64  `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	EndTs       *int64    `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	ItemConfigs *string   `json:"ItemConfigs,omitempty" xml:"ItemConfigs,omitempty"`
	MetricTypes []*string `json:"MetricTypes,omitempty" xml:"MetricTypes,omitempty" type:"Repeated"`
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	OrderName   *string   `json:"OrderName,omitempty" xml:"OrderName,omitempty"`
	OrderType   *string   `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Os          *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribePlayQoeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQoeListRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayQoeListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayQoeListRequest) GetBeginTs() *int64 {
	return s.BeginTs
}

func (s *DescribePlayQoeListRequest) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayQoeListRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribePlayQoeListRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribePlayQoeListRequest) GetMetricTypes() []*string {
	return s.MetricTypes
}

func (s *DescribePlayQoeListRequest) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayQoeListRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribePlayQoeListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePlayQoeListRequest) GetOs() *string {
	return s.Os
}

func (s *DescribePlayQoeListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQoeListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQoeListRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayQoeListRequest) SetAppName(v string) *DescribePlayQoeListRequest {
	s.AppName = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetBeginTs(v int64) *DescribePlayQoeListRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetDefinition(v string) *DescribePlayQoeListRequest {
	s.Definition = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetEndTs(v int64) *DescribePlayQoeListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetItemConfigs(v string) *DescribePlayQoeListRequest {
	s.ItemConfigs = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetMetricTypes(v []*string) *DescribePlayQoeListRequest {
	s.MetricTypes = v
	return s
}

func (s *DescribePlayQoeListRequest) SetNetwork(v string) *DescribePlayQoeListRequest {
	s.Network = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetOrderName(v string) *DescribePlayQoeListRequest {
	s.OrderName = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetOrderType(v string) *DescribePlayQoeListRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetOs(v string) *DescribePlayQoeListRequest {
	s.Os = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetPageNo(v int64) *DescribePlayQoeListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetPageSize(v int64) *DescribePlayQoeListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQoeListRequest) SetTerminalType(v string) *DescribePlayQoeListRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayQoeListRequest) Validate() error {
	return dara.Validate(s)
}
