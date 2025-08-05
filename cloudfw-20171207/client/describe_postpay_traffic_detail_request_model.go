// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *DescribePostpayTrafficDetailRequest
	GetCurrentPage() *int64
	SetEndTime(v string) *DescribePostpayTrafficDetailRequest
	GetEndTime() *string
	SetLang(v string) *DescribePostpayTrafficDetailRequest
	GetLang() *string
	SetOrder(v string) *DescribePostpayTrafficDetailRequest
	GetOrder() *string
	SetPageSize(v int64) *DescribePostpayTrafficDetailRequest
	GetPageSize() *int64
	SetRegionNo(v string) *DescribePostpayTrafficDetailRequest
	GetRegionNo() *string
	SetSearchItem(v string) *DescribePostpayTrafficDetailRequest
	GetSearchItem() *string
	SetStartTime(v string) *DescribePostpayTrafficDetailRequest
	GetStartTime() *string
	SetTrafficType(v string) *DescribePostpayTrafficDetailRequest
	GetTrafficType() *string
}

type DescribePostpayTrafficDetailRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Specify a value in the YYYYMMDD format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230130
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// 	- **resourceId**
	//
	// 	- **trafficDay**
	//
	// example:
	//
	// resourceId
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The instance ID or the IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The beginning of the time range to query. Specify a value in the YYYYMMDD format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230101
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The traffic type. This parameter is required. Valid values:
	//
	// 	- **EIP_TRAFFIC**: traffic for the Internet firewall.
	//
	// 	- **NatGateway_TRAFFIC**: traffic for NAT firewalls.
	//
	// 	- **VPC_TRAFFIC**: traffic for virtual private cloud (VPC) firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// EIP_TRAFFIC
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribePostpayTrafficDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePostpayTrafficDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePostpayTrafficDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayTrafficDetailRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribePostpayTrafficDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePostpayTrafficDetailRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePostpayTrafficDetailRequest) GetSearchItem() *string {
	return s.SearchItem
}

func (s *DescribePostpayTrafficDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePostpayTrafficDetailRequest) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribePostpayTrafficDetailRequest) SetCurrentPage(v int64) *DescribePostpayTrafficDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetEndTime(v string) *DescribePostpayTrafficDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetLang(v string) *DescribePostpayTrafficDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetOrder(v string) *DescribePostpayTrafficDetailRequest {
	s.Order = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetPageSize(v int64) *DescribePostpayTrafficDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetRegionNo(v string) *DescribePostpayTrafficDetailRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetSearchItem(v string) *DescribePostpayTrafficDetailRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetStartTime(v string) *DescribePostpayTrafficDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetTrafficType(v string) *DescribePostpayTrafficDetailRequest {
	s.TrafficType = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) Validate() error {
	return dara.Validate(s)
}
