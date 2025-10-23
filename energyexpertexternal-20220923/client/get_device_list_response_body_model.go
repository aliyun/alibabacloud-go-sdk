// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceListResponseBody
	GetCode() *string
	SetData(v *GetDeviceListResponseBodyData) *GetDeviceListResponseBody
	GetData() *GetDeviceListResponseBodyData
	SetHttpCode(v int32) *GetDeviceListResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeviceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeviceListResponseBody
	GetSuccess() *bool
}

type GetDeviceListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceListResponseBody) GetData() *GetDeviceListResponseBodyData {
	return s.Data
}

func (s *GetDeviceListResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeviceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeviceListResponseBody) SetCode(v string) *GetDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBody) SetData(v *GetDeviceListResponseBodyData) *GetDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceListResponseBody) SetHttpCode(v int32) *GetDeviceListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBody) SetRequestId(v string) *GetDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceListResponseBody) SetSuccess(v bool) *GetDeviceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeviceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceListResponseBodyData struct {
	// The code returned for the request.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The devices.
	DeviceList []*GetDeviceListResponseBodyDataDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	// The ID of the site.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetDeviceListResponseBodyData) GetDeviceList() []*GetDeviceListResponseBodyDataDeviceList {
	return s.DeviceList
}

func (s *GetDeviceListResponseBodyData) GetFactoryId() *string {
	return s.FactoryId
}

func (s *GetDeviceListResponseBodyData) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeviceListResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeviceListResponseBodyData) SetCode(v string) *GetDeviceListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetDeviceList(v []*GetDeviceListResponseBodyDataDeviceList) *GetDeviceListResponseBodyData {
	s.DeviceList = v
	return s
}

func (s *GetDeviceListResponseBodyData) SetFactoryId(v string) *GetDeviceListResponseBodyData {
	s.FactoryId = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetHttpCode(v int32) *GetDeviceListResponseBodyData {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetSuccess(v bool) *GetDeviceListResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetDeviceListResponseBodyData) Validate() error {
	if s.DeviceList != nil {
		for _, item := range s.DeviceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDeviceListResponseBodyDataDeviceList struct {
	// The device ID.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The device name.
	//
	// example:
	//
	// Main transformer 4#
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// The level 1 meter type.
	//
	// example:
	//
	// Electric meter
	FirstTypeName *string `json:"firstTypeName,omitempty" xml:"firstTypeName,omitempty"`
	// The device information.
	Info *GetDeviceListResponseBodyDataDeviceListInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	// The ID of the parent device.
	//
	// example:
	//
	// pn_6987
	ParentDevice *string `json:"parentDevice,omitempty" xml:"parentDevice,omitempty"`
	// The level 2 meter type.
	//
	// example:
	//
	// Gateway meter
	SecondTypeName *string `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceList) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceList) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetFirstTypeName() *string {
	return s.FirstTypeName
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetInfo() *GetDeviceListResponseBodyDataDeviceListInfo {
	return s.Info
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetParentDevice() *string {
	return s.ParentDevice
}

func (s *GetDeviceListResponseBodyDataDeviceList) GetSecondTypeName() *string {
	return s.SecondTypeName
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceId(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetFirstTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetInfo(v *GetDeviceListResponseBodyDataDeviceListInfo) *GetDeviceListResponseBodyDataDeviceList {
	s.Info = v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetParentDevice(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.ParentDevice = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetSecondTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.SecondTypeName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) Validate() error {
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceListResponseBodyDataDeviceListInfo struct {
	// The rated capacity.
	//
	// Unit is kVA.
	//
	// example:
	//
	// 100
	ConstKva *int32 `json:"constKva,omitempty" xml:"constKva,omitempty"`
	// The transformation ratio of current.
	//
	// example:
	//
	// 1
	Ct *int32 `json:"ct,omitempty" xml:"ct,omitempty"`
	// The magnification ratio.
	//
	// example:
	//
	// 80
	Magnification *int32 `json:"magnification,omitempty" xml:"magnification,omitempty"`
	// The high and low voltage.
	//
	// example:
	//
	// 0
	Pressure *int32 `json:"pressure,omitempty" xml:"pressure,omitempty"`
	// The transformation ratio of voltage.
	//
	// example:
	//
	// 80
	Pt *int32 `json:"pt,omitempty" xml:"pt,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) GetConstKva() *int32 {
	return s.ConstKva
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) GetCt() *int32 {
	return s.Ct
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) GetMagnification() *int32 {
	return s.Magnification
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) GetPressure() *int32 {
	return s.Pressure
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) GetPt() *int32 {
	return s.Pt
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetConstKva(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.ConstKva = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetCt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Ct = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetMagnification(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Magnification = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPressure(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pressure = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pt = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) Validate() error {
	return dara.Validate(s)
}
