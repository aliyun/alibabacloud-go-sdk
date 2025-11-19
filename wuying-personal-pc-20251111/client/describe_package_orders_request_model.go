// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribePackageOrdersRequest
	GetApiKey() *string
	SetCurrentPage(v int64) *DescribePackageOrdersRequest
	GetCurrentPage() *int64
	SetDesktopIdList(v []*string) *DescribePackageOrdersRequest
	GetDesktopIdList() []*string
	SetOrderIdList(v []*string) *DescribePackageOrdersRequest
	GetOrderIdList() []*string
	SetOrderStatusList(v []*string) *DescribePackageOrdersRequest
	GetOrderStatusList() []*string
	SetPageSize(v int64) *DescribePackageOrdersRequest
	GetPageSize() *int64
	SetProductTypeList(v []*string) *DescribePackageOrdersRequest
	GetProductTypeList() []*string
	SetResourceIdList(v []*string) *DescribePackageOrdersRequest
	GetResourceIdList() []*string
	SetSessionId(v string) *DescribePackageOrdersRequest
	GetSessionId() *string
}

type DescribePackageOrdersRequest struct {
	// This parameter is required.
	ApiKey          *string   `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	CurrentPage     *int64    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DesktopIdList   []*string `json:"DesktopIdList,omitempty" xml:"DesktopIdList,omitempty" type:"Repeated"`
	OrderIdList     []*string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty" type:"Repeated"`
	OrderStatusList []*string `json:"OrderStatusList,omitempty" xml:"OrderStatusList,omitempty" type:"Repeated"`
	PageSize        *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductTypeList []*string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty" type:"Repeated"`
	ResourceIdList  []*string `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty" type:"Repeated"`
	SessionId       *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribePackageOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribePackageOrdersRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePackageOrdersRequest) GetDesktopIdList() []*string {
	return s.DesktopIdList
}

func (s *DescribePackageOrdersRequest) GetOrderIdList() []*string {
	return s.OrderIdList
}

func (s *DescribePackageOrdersRequest) GetOrderStatusList() []*string {
	return s.OrderStatusList
}

func (s *DescribePackageOrdersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePackageOrdersRequest) GetProductTypeList() []*string {
	return s.ProductTypeList
}

func (s *DescribePackageOrdersRequest) GetResourceIdList() []*string {
	return s.ResourceIdList
}

func (s *DescribePackageOrdersRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePackageOrdersRequest) SetApiKey(v string) *DescribePackageOrdersRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribePackageOrdersRequest) SetCurrentPage(v int64) *DescribePackageOrdersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackageOrdersRequest) SetDesktopIdList(v []*string) *DescribePackageOrdersRequest {
	s.DesktopIdList = v
	return s
}

func (s *DescribePackageOrdersRequest) SetOrderIdList(v []*string) *DescribePackageOrdersRequest {
	s.OrderIdList = v
	return s
}

func (s *DescribePackageOrdersRequest) SetOrderStatusList(v []*string) *DescribePackageOrdersRequest {
	s.OrderStatusList = v
	return s
}

func (s *DescribePackageOrdersRequest) SetPageSize(v int64) *DescribePackageOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageOrdersRequest) SetProductTypeList(v []*string) *DescribePackageOrdersRequest {
	s.ProductTypeList = v
	return s
}

func (s *DescribePackageOrdersRequest) SetResourceIdList(v []*string) *DescribePackageOrdersRequest {
	s.ResourceIdList = v
	return s
}

func (s *DescribePackageOrdersRequest) SetSessionId(v string) *DescribePackageOrdersRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePackageOrdersRequest) Validate() error {
	return dara.Validate(s)
}
