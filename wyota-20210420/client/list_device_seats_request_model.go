// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabel(v string) *ListDeviceSeatsRequest
	GetLabel() *string
	SetSeatNo(v string) *ListDeviceSeatsRequest
	GetSeatNo() *string
	SetSerialNoList(v []*string) *ListDeviceSeatsRequest
	GetSerialNoList() []*string
	SetTenantId(v string) *ListDeviceSeatsRequest
	GetTenantId() *string
	SetZoneId(v string) *ListDeviceSeatsRequest
	GetZoneId() *string
}

type ListDeviceSeatsRequest struct {
	Label        *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	SeatNo       *string   `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId       *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDeviceSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsRequest) GetLabel() *string {
	return s.Label
}

func (s *ListDeviceSeatsRequest) GetSeatNo() *string {
	return s.SeatNo
}

func (s *ListDeviceSeatsRequest) GetSerialNoList() []*string {
	return s.SerialNoList
}

func (s *ListDeviceSeatsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDeviceSeatsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListDeviceSeatsRequest) SetLabel(v string) *ListDeviceSeatsRequest {
	s.Label = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetSeatNo(v string) *ListDeviceSeatsRequest {
	s.SeatNo = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetSerialNoList(v []*string) *ListDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

func (s *ListDeviceSeatsRequest) SetTenantId(v string) *ListDeviceSeatsRequest {
	s.TenantId = &v
	return s
}

func (s *ListDeviceSeatsRequest) SetZoneId(v string) *ListDeviceSeatsRequest {
	s.ZoneId = &v
	return s
}

func (s *ListDeviceSeatsRequest) Validate() error {
	return dara.Validate(s)
}
