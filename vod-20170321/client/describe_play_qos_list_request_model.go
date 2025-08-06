// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQosListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribePlayQosListRequest
	GetAppName() *string
	SetBeginTs(v int64) *DescribePlayQosListRequest
	GetBeginTs() *int64
	SetDefinition(v string) *DescribePlayQosListRequest
	GetDefinition() *string
	SetEndTs(v int64) *DescribePlayQosListRequest
	GetEndTs() *int64
	SetItemConfigs(v string) *DescribePlayQosListRequest
	GetItemConfigs() *string
	SetMetricTypes(v []*string) *DescribePlayQosListRequest
	GetMetricTypes() []*string
	SetNetwork(v string) *DescribePlayQosListRequest
	GetNetwork() *string
	SetOrderName(v string) *DescribePlayQosListRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribePlayQosListRequest
	GetOrderType() *string
	SetOs(v string) *DescribePlayQosListRequest
	GetOs() *string
	SetPageNo(v int64) *DescribePlayQosListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQosListRequest
	GetPageSize() *int64
	SetTerminalType(v string) *DescribePlayQosListRequest
	GetTerminalType() *string
}

type DescribePlayQosListRequest struct {
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

func (s DescribePlayQosListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQosListRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayQosListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayQosListRequest) GetBeginTs() *int64 {
	return s.BeginTs
}

func (s *DescribePlayQosListRequest) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayQosListRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribePlayQosListRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribePlayQosListRequest) GetMetricTypes() []*string {
	return s.MetricTypes
}

func (s *DescribePlayQosListRequest) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayQosListRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribePlayQosListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePlayQosListRequest) GetOs() *string {
	return s.Os
}

func (s *DescribePlayQosListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQosListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQosListRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayQosListRequest) SetAppName(v string) *DescribePlayQosListRequest {
	s.AppName = &v
	return s
}

func (s *DescribePlayQosListRequest) SetBeginTs(v int64) *DescribePlayQosListRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayQosListRequest) SetDefinition(v string) *DescribePlayQosListRequest {
	s.Definition = &v
	return s
}

func (s *DescribePlayQosListRequest) SetEndTs(v int64) *DescribePlayQosListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayQosListRequest) SetItemConfigs(v string) *DescribePlayQosListRequest {
	s.ItemConfigs = &v
	return s
}

func (s *DescribePlayQosListRequest) SetMetricTypes(v []*string) *DescribePlayQosListRequest {
	s.MetricTypes = v
	return s
}

func (s *DescribePlayQosListRequest) SetNetwork(v string) *DescribePlayQosListRequest {
	s.Network = &v
	return s
}

func (s *DescribePlayQosListRequest) SetOrderName(v string) *DescribePlayQosListRequest {
	s.OrderName = &v
	return s
}

func (s *DescribePlayQosListRequest) SetOrderType(v string) *DescribePlayQosListRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePlayQosListRequest) SetOs(v string) *DescribePlayQosListRequest {
	s.Os = &v
	return s
}

func (s *DescribePlayQosListRequest) SetPageNo(v int64) *DescribePlayQosListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQosListRequest) SetPageSize(v int64) *DescribePlayQosListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQosListRequest) SetTerminalType(v string) *DescribePlayQosListRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayQosListRequest) Validate() error {
	return dara.Validate(s)
}
