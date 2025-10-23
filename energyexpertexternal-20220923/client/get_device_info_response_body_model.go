// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceInfoResponseBody
	GetCode() *string
	SetData(v *GetDeviceInfoResponseBodyData) *GetDeviceInfoResponseBody
	GetData() *GetDeviceInfoResponseBodyData
	SetHttpCode(v int32) *GetDeviceInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeviceInfoResponseBody
	GetSuccess() *bool
}

type GetDeviceInfoResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDeviceInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceInfoResponseBody) GetData() *GetDeviceInfoResponseBodyData {
	return s.Data
}

func (s *GetDeviceInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeviceInfoResponseBody) SetCode(v string) *GetDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetData(v *GetDeviceInfoResponseBodyData) *GetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceInfoResponseBody) SetHttpCode(v int32) *GetDeviceInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetRequestId(v string) *GetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetSuccess(v bool) *GetDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceInfoResponseBodyData struct {
	// The ID of the device.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The name of the device.
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
	// The device parameters.
	RecordList []*GetDeviceInfoResponseBodyDataRecordList `json:"recordList,omitempty" xml:"recordList,omitempty" type:"Repeated"`
	// The level 2 meter type.
	//
	// example:
	//
	// Gateway meter
	SecondTypeName *string `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceInfoResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetDeviceInfoResponseBodyData) GetFirstTypeName() *string {
	return s.FirstTypeName
}

func (s *GetDeviceInfoResponseBodyData) GetRecordList() []*GetDeviceInfoResponseBodyDataRecordList {
	return s.RecordList
}

func (s *GetDeviceInfoResponseBodyData) GetSecondTypeName() *string {
	return s.SecondTypeName
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceId(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceName(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetFirstTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetRecordList(v []*GetDeviceInfoResponseBodyDataRecordList) *GetDeviceInfoResponseBodyData {
	s.RecordList = v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetSecondTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.SecondTypeName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) Validate() error {
	if s.RecordList != nil {
		for _, item := range s.RecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDeviceInfoResponseBodyDataRecordList struct {
	// The device identifier.
	//
	// example:
	//
	// Ia
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// Phase A current
	ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
	// The date on which the statistics were collected.
	//
	// example:
	//
	// 2022-07-26 00:00:00
	StatisticsDate *string `json:"statisticsDate,omitempty" xml:"statisticsDate,omitempty"`
	// The type of the measuring point.
	//
	// example:
	//
	// DOUBLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The unit of the parameter value.
	//
	// example:
	//
	// A
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The value of the measuring point.
	//
	// example:
	//
	// 20.00
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDeviceInfoResponseBodyDataRecordList) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceInfoResponseBodyDataRecordList) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetParamName() *string {
	return s.ParamName
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetStatisticsDate() *string {
	return s.StatisticsDate
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetType() *string {
	return s.Type
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetUnit() *string {
	return s.Unit
}

func (s *GetDeviceInfoResponseBodyDataRecordList) GetValue() *float64 {
	return s.Value
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetIdentifier(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Identifier = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetParamName(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.ParamName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetStatisticsDate(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.StatisticsDate = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetType(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Type = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetUnit(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Unit = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetValue(v float64) *GetDeviceInfoResponseBodyDataRecordList {
	s.Value = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) Validate() error {
	return dara.Validate(s)
}
