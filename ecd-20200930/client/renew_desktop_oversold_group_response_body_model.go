// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RenewDesktopOversoldGroupResponseBodyData) *RenewDesktopOversoldGroupResponseBody
	GetData() *RenewDesktopOversoldGroupResponseBodyData
	SetRequestId(v string) *RenewDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type RenewDesktopOversoldGroupResponseBody struct {
	Data      *RenewDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopOversoldGroupResponseBody) GetData() *RenewDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *RenewDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewDesktopOversoldGroupResponseBody) SetData(v *RenewDesktopOversoldGroupResponseBodyData) *RenewDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *RenewDesktopOversoldGroupResponseBody) SetRequestId(v string) *RenewDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopOversoldGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewDesktopOversoldGroupResponseBodyData struct {
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
}

func (s RenewDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewDesktopOversoldGroupResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewDesktopOversoldGroupResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *RenewDesktopOversoldGroupResponseBodyData) SetOrderId(v int64) *RenewDesktopOversoldGroupResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewDesktopOversoldGroupResponseBodyData) SetOversoldGroupId(v string) *RenewDesktopOversoldGroupResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *RenewDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
