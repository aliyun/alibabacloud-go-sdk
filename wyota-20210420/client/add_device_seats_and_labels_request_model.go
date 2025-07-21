// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceSeatsAndLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsUnique(v bool) *AddDeviceSeatsAndLabelsRequest
	GetIsUnique() *bool
	SetLabel(v string) *AddDeviceSeatsAndLabelsRequest
	GetLabel() *string
	SetLabelList(v []*string) *AddDeviceSeatsAndLabelsRequest
	GetLabelList() []*string
	SetSeatName(v string) *AddDeviceSeatsAndLabelsRequest
	GetSeatName() *string
	SetSerialNo(v string) *AddDeviceSeatsAndLabelsRequest
	GetSerialNo() *string
	SetTenantId(v string) *AddDeviceSeatsAndLabelsRequest
	GetTenantId() *string
	SetZoneId(v string) *AddDeviceSeatsAndLabelsRequest
	GetZoneId() *string
}

type AddDeviceSeatsAndLabelsRequest struct {
	IsUnique  *bool     `json:"IsUnique,omitempty" xml:"IsUnique,omitempty"`
	Label     *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	SeatName  *string   `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SerialNo  *string   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId  *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId    *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddDeviceSeatsAndLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsRequest) GetIsUnique() *bool {
	return s.IsUnique
}

func (s *AddDeviceSeatsAndLabelsRequest) GetLabel() *string {
	return s.Label
}

func (s *AddDeviceSeatsAndLabelsRequest) GetLabelList() []*string {
	return s.LabelList
}

func (s *AddDeviceSeatsAndLabelsRequest) GetSeatName() *string {
	return s.SeatName
}

func (s *AddDeviceSeatsAndLabelsRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *AddDeviceSeatsAndLabelsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddDeviceSeatsAndLabelsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddDeviceSeatsAndLabelsRequest) SetIsUnique(v bool) *AddDeviceSeatsAndLabelsRequest {
	s.IsUnique = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetLabel(v string) *AddDeviceSeatsAndLabelsRequest {
	s.Label = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetLabelList(v []*string) *AddDeviceSeatsAndLabelsRequest {
	s.LabelList = v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetSeatName(v string) *AddDeviceSeatsAndLabelsRequest {
	s.SeatName = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetSerialNo(v string) *AddDeviceSeatsAndLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetTenantId(v string) *AddDeviceSeatsAndLabelsRequest {
	s.TenantId = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) SetZoneId(v string) *AddDeviceSeatsAndLabelsRequest {
	s.ZoneId = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsRequest) Validate() error {
	return dara.Validate(s)
}
