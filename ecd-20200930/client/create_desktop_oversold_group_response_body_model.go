// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDesktopOversoldGroupResponseBodyData) *CreateDesktopOversoldGroupResponseBody
	GetData() *CreateDesktopOversoldGroupResponseBodyData
	SetRequestId(v string) *CreateDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type CreateDesktopOversoldGroupResponseBody struct {
	Data      *CreateDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopOversoldGroupResponseBody) GetData() *CreateDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *CreateDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDesktopOversoldGroupResponseBody) SetData(v *CreateDesktopOversoldGroupResponseBodyData) *CreateDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateDesktopOversoldGroupResponseBody) SetRequestId(v string) *CreateDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopOversoldGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDesktopOversoldGroupResponseBodyData struct {
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
}

func (s CreateDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDesktopOversoldGroupResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateDesktopOversoldGroupResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *CreateDesktopOversoldGroupResponseBodyData) SetOrderId(v int64) *CreateDesktopOversoldGroupResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopOversoldGroupResponseBodyData) SetOversoldGroupId(v string) *CreateDesktopOversoldGroupResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *CreateDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
