// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillPreQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsigneeInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetConsigneeInfoShrink() *string
	SetCpCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetCpCode() *string
	SetOrderChannels(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetOrderChannels() *string
	SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetOuterOrderCode() *string
	SetPreWeight(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetPreWeight() *string
	SetSenderInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest
	GetSenderInfoShrink() *string
}

type CreatePickUpWaybillPreQueryShrinkRequest struct {
	// The consignee information.
	//
	// This parameter is required.
	ConsigneeInfoShrink *string `json:"ConsigneeInfo,omitempty" xml:"ConsigneeInfo,omitempty"`
	// The code of the courier company. If no courier company is specified, the system allocates a courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The identifier of the external channel source. It cannot contain underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The order number of the access system.
	//
	// example:
	//
	// 787DFHHDS989****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The estimated weight. Unit: gram.
	//
	// >  If you need to query the estimated price, this parameter is required.
	//
	// example:
	//
	// 2000
	PreWeight *string `json:"PreWeight,omitempty" xml:"PreWeight,omitempty"`
	// The sender information.
	//
	// This parameter is required.
	SenderInfoShrink *string `json:"SenderInfo,omitempty" xml:"SenderInfo,omitempty"`
}

func (s CreatePickUpWaybillPreQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetConsigneeInfoShrink() *string {
	return s.ConsigneeInfoShrink
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetCpCode() *string {
	return s.CpCode
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetOrderChannels() *string {
	return s.OrderChannels
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetOuterOrderCode() *string {
	return s.OuterOrderCode
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetPreWeight() *string {
	return s.PreWeight
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) GetSenderInfoShrink() *string {
	return s.SenderInfoShrink
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetConsigneeInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.ConsigneeInfoShrink = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetCpCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetOrderChannels(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetPreWeight(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.PreWeight = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) SetSenderInfoShrink(v string) *CreatePickUpWaybillPreQueryShrinkRequest {
	s.SenderInfoShrink = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
