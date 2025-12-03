// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDeliveryTime(v string) *DescribeCouponListRequest
	GetEndDeliveryTime() *string
	SetPageNum(v int32) *DescribeCouponListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeCouponListRequest
	GetPageSize() *int32
	SetStartDeliveryTime(v string) *DescribeCouponListRequest
	GetStartDeliveryTime() *string
	SetStatus(v string) *DescribeCouponListRequest
	GetStatus() *string
}

type DescribeCouponListRequest struct {
	EndDeliveryTime   *string `json:"EndDeliveryTime,omitempty" xml:"EndDeliveryTime,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDeliveryTime *string `json:"StartDeliveryTime,omitempty" xml:"StartDeliveryTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponListRequest) GetEndDeliveryTime() *string {
	return s.EndDeliveryTime
}

func (s *DescribeCouponListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCouponListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponListRequest) GetStartDeliveryTime() *string {
	return s.StartDeliveryTime
}

func (s *DescribeCouponListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCouponListRequest) SetEndDeliveryTime(v string) *DescribeCouponListRequest {
	s.EndDeliveryTime = &v
	return s
}

func (s *DescribeCouponListRequest) SetPageNum(v int32) *DescribeCouponListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCouponListRequest) SetPageSize(v int32) *DescribeCouponListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponListRequest) SetStartDeliveryTime(v string) *DescribeCouponListRequest {
	s.StartDeliveryTime = &v
	return s
}

func (s *DescribeCouponListRequest) SetStatus(v string) *DescribeCouponListRequest {
	s.Status = &v
	return s
}

func (s *DescribeCouponListRequest) Validate() error {
	return dara.Validate(s)
}
