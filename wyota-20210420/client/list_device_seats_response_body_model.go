// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDeviceSeatsResponseBody
	GetCode() *string
	SetData(v *ListDeviceSeatsResponseBodyData) *ListDeviceSeatsResponseBody
	GetData() *ListDeviceSeatsResponseBodyData
	SetMessage(v string) *ListDeviceSeatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceSeatsResponseBody
	GetRequestId() *string
}

type ListDeviceSeatsResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDeviceSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDeviceSeatsResponseBody) GetData() *ListDeviceSeatsResponseBodyData {
	return s.Data
}

func (s *ListDeviceSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceSeatsResponseBody) SetCode(v string) *ListDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetData(v *ListDeviceSeatsResponseBodyData) *ListDeviceSeatsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetMessage(v string) *ListDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceSeatsResponseBody) SetRequestId(v string) *ListDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceSeatsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeviceSeatsResponseBodyData struct {
	DeviceSeatDTOList []*ListDeviceSeatsResponseBodyDataDeviceSeatDTOList `json:"DeviceSeatDTOList,omitempty" xml:"DeviceSeatDTOList,omitempty" type:"Repeated"`
	TotalCount        *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceSeatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBodyData) GetDeviceSeatDTOList() []*ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	return s.DeviceSeatDTOList
}

func (s *ListDeviceSeatsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeviceSeatsResponseBodyData) SetDeviceSeatDTOList(v []*ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) *ListDeviceSeatsResponseBodyData {
	s.DeviceSeatDTOList = v
	return s
}

func (s *ListDeviceSeatsResponseBodyData) SetTotalCount(v int64) *ListDeviceSeatsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDeviceSeatsResponseBodyDataDeviceSeatDTOList struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	SeatName *string `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SeatNo   *string `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetLabel() *string {
	return s.Label
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetSeatName() *string {
	return s.SeatName
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetSeatNo() *string {
	return s.SeatNo
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetSiteId() *string {
	return s.SiteId
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetSiteName() *string {
	return s.SiteName
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetLabel(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.Label = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSeatName(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SeatName = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSeatNo(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SeatNo = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSerialNo(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SerialNo = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSiteId(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SiteId = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetSiteName(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.SiteName = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) SetZoneId(v string) *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList {
	s.ZoneId = &v
	return s
}

func (s *ListDeviceSeatsResponseBodyDataDeviceSeatDTOList) Validate() error {
	return dara.Validate(s)
}
