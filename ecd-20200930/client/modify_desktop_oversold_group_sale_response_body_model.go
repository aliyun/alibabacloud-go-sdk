// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupSaleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDesktopOversoldGroupSaleResponseBodyData) *ModifyDesktopOversoldGroupSaleResponseBody
	GetData() *ModifyDesktopOversoldGroupSaleResponseBodyData
	SetRequestId(v string) *ModifyDesktopOversoldGroupSaleResponseBody
	GetRequestId() *string
}

type ModifyDesktopOversoldGroupSaleResponseBody struct {
	Data      *ModifyDesktopOversoldGroupSaleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopOversoldGroupSaleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupSaleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupSaleResponseBody) GetData() *ModifyDesktopOversoldGroupSaleResponseBodyData {
	return s.Data
}

func (s *ModifyDesktopOversoldGroupSaleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopOversoldGroupSaleResponseBody) SetData(v *ModifyDesktopOversoldGroupSaleResponseBodyData) *ModifyDesktopOversoldGroupSaleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponseBody) SetRequestId(v string) *ModifyDesktopOversoldGroupSaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDesktopOversoldGroupSaleResponseBodyData struct {
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
}

func (s ModifyDesktopOversoldGroupSaleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupSaleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupSaleResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDesktopOversoldGroupSaleResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *ModifyDesktopOversoldGroupSaleResponseBodyData) SetOrderId(v int64) *ModifyDesktopOversoldGroupSaleResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponseBodyData) SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupSaleResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
