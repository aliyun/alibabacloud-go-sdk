// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageOrdersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribePackageOrdersShrinkRequest
	GetApiKey() *string
	SetCurrentPage(v int64) *DescribePackageOrdersShrinkRequest
	GetCurrentPage() *int64
	SetDesktopIdListShrink(v string) *DescribePackageOrdersShrinkRequest
	GetDesktopIdListShrink() *string
	SetOrderIdListShrink(v string) *DescribePackageOrdersShrinkRequest
	GetOrderIdListShrink() *string
	SetOrderStatusListShrink(v string) *DescribePackageOrdersShrinkRequest
	GetOrderStatusListShrink() *string
	SetPageSize(v int64) *DescribePackageOrdersShrinkRequest
	GetPageSize() *int64
	SetProductTypeListShrink(v string) *DescribePackageOrdersShrinkRequest
	GetProductTypeListShrink() *string
	SetResourceIdListShrink(v string) *DescribePackageOrdersShrinkRequest
	GetResourceIdListShrink() *string
	SetSessionId(v string) *DescribePackageOrdersShrinkRequest
	GetSessionId() *string
}

type DescribePackageOrdersShrinkRequest struct {
	// This parameter is required.
	ApiKey                *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	CurrentPage           *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DesktopIdListShrink   *string `json:"DesktopIdList,omitempty" xml:"DesktopIdList,omitempty"`
	OrderIdListShrink     *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
	OrderStatusListShrink *string `json:"OrderStatusList,omitempty" xml:"OrderStatusList,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductTypeListShrink *string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty"`
	ResourceIdListShrink  *string `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty"`
	SessionId             *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribePackageOrdersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribePackageOrdersShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePackageOrdersShrinkRequest) GetDesktopIdListShrink() *string {
	return s.DesktopIdListShrink
}

func (s *DescribePackageOrdersShrinkRequest) GetOrderIdListShrink() *string {
	return s.OrderIdListShrink
}

func (s *DescribePackageOrdersShrinkRequest) GetOrderStatusListShrink() *string {
	return s.OrderStatusListShrink
}

func (s *DescribePackageOrdersShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePackageOrdersShrinkRequest) GetProductTypeListShrink() *string {
	return s.ProductTypeListShrink
}

func (s *DescribePackageOrdersShrinkRequest) GetResourceIdListShrink() *string {
	return s.ResourceIdListShrink
}

func (s *DescribePackageOrdersShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePackageOrdersShrinkRequest) SetApiKey(v string) *DescribePackageOrdersShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetCurrentPage(v int64) *DescribePackageOrdersShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetDesktopIdListShrink(v string) *DescribePackageOrdersShrinkRequest {
	s.DesktopIdListShrink = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetOrderIdListShrink(v string) *DescribePackageOrdersShrinkRequest {
	s.OrderIdListShrink = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetOrderStatusListShrink(v string) *DescribePackageOrdersShrinkRequest {
	s.OrderStatusListShrink = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetPageSize(v int64) *DescribePackageOrdersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetProductTypeListShrink(v string) *DescribePackageOrdersShrinkRequest {
	s.ProductTypeListShrink = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetResourceIdListShrink(v string) *DescribePackageOrdersShrinkRequest {
	s.ResourceIdListShrink = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) SetSessionId(v string) *DescribePackageOrdersShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePackageOrdersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
