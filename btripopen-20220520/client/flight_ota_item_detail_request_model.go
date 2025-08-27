// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaItemDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightOtaItemDetailRequest
	GetIsvName() *string
	SetOtaItemId(v string) *FlightOtaItemDetailRequest
	GetOtaItemId() *string
}

type FlightOtaItemDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cheshi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68cdc6b37c87484c98b479b49306ffbb_0
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
}

func (s FlightOtaItemDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailRequest) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightOtaItemDetailRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightOtaItemDetailRequest) SetIsvName(v string) *FlightOtaItemDetailRequest {
	s.IsvName = &v
	return s
}

func (s *FlightOtaItemDetailRequest) SetOtaItemId(v string) *FlightOtaItemDetailRequest {
	s.OtaItemId = &v
	return s
}

func (s *FlightOtaItemDetailRequest) Validate() error {
	return dara.Validate(s)
}
