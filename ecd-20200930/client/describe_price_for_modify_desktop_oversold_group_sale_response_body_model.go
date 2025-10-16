// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForModifyDesktopOversoldGroupSaleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody
	GetData() *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData
	SetRequestId(v string) *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody
	GetRequestId() *string
}

type DescribePriceForModifyDesktopOversoldGroupSaleResponseBody struct {
	Data      *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) GetData() *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData {
	return s.Data
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) SetData(v *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody {
	s.Data = v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) SetRequestId(v string) *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData struct {
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) GetPrice() *string {
	return s.Price
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) SetPrice(v string) *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData {
	s.Price = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
