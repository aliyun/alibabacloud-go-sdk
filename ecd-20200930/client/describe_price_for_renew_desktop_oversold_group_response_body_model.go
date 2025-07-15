// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForRenewDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePriceForRenewDesktopOversoldGroupResponseBodyData) *DescribePriceForRenewDesktopOversoldGroupResponseBody
	GetData() *DescribePriceForRenewDesktopOversoldGroupResponseBodyData
	SetRequestId(v string) *DescribePriceForRenewDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type DescribePriceForRenewDesktopOversoldGroupResponseBody struct {
	Data      *DescribePriceForRenewDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceForRenewDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForRenewDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBody) GetData() *DescribePriceForRenewDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBody) SetData(v *DescribePriceForRenewDesktopOversoldGroupResponseBodyData) *DescribePriceForRenewDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBody) SetRequestId(v string) *DescribePriceForRenewDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePriceForRenewDesktopOversoldGroupResponseBodyData struct {
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s DescribePriceForRenewDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForRenewDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBodyData) GetPrice() *string {
	return s.Price
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBodyData) SetPrice(v string) *DescribePriceForRenewDesktopOversoldGroupResponseBodyData {
	s.Price = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
