// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForCreateDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePriceForCreateDesktopOversoldGroupResponseBodyData) *DescribePriceForCreateDesktopOversoldGroupResponseBody
	GetData() *DescribePriceForCreateDesktopOversoldGroupResponseBodyData
	SetRequestId(v string) *DescribePriceForCreateDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type DescribePriceForCreateDesktopOversoldGroupResponseBody struct {
	Data      *DescribePriceForCreateDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceForCreateDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForCreateDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBody) GetData() *DescribePriceForCreateDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBody) SetData(v *DescribePriceForCreateDesktopOversoldGroupResponseBodyData) *DescribePriceForCreateDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBody) SetRequestId(v string) *DescribePriceForCreateDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceForCreateDesktopOversoldGroupResponseBodyData struct {
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s DescribePriceForCreateDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForCreateDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBodyData) GetPrice() *string {
	return s.Price
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBodyData) SetPrice(v string) *DescribePriceForCreateDesktopOversoldGroupResponseBodyData {
	s.Price = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
