// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponItemListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCouponItemListResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeCouponItemListResponseBodyData) *DescribeCouponItemListResponseBody
	GetData() []*DescribeCouponItemListResponseBodyData
	SetPageSize(v int32) *DescribeCouponItemListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCouponItemListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCouponItemListResponseBody
	GetTotalCount() *int32
}

type DescribeCouponItemListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*DescribeCouponItemListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EAE08A27-386C-579E-966D-8853EC3C5D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCouponItemListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponItemListResponseBody) GetData() []*DescribeCouponItemListResponseBodyData {
	return s.Data
}

func (s *DescribeCouponItemListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponItemListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCouponItemListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCouponItemListResponseBody) SetCurrentPage(v int32) *DescribeCouponItemListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetData(v []*DescribeCouponItemListResponseBodyData) *DescribeCouponItemListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetPageSize(v int32) *DescribeCouponItemListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetRequestId(v string) *DescribeCouponItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetTotalCount(v int32) *DescribeCouponItemListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCouponItemListResponseBodyData struct {
	// example:
	//
	// vm
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCouponItemListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DescribeCouponItemListResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeCouponItemListResponseBodyData) SetCode(v string) *DescribeCouponItemListResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribeCouponItemListResponseBodyData) SetName(v string) *DescribeCouponItemListResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeCouponItemListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
