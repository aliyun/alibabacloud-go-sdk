// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSyntheticTaskMonitorsResponseBody
	GetCode() *string
	SetData(v []*GetSyntheticTaskMonitorsResponseBodyData) *GetSyntheticTaskMonitorsResponseBody
	GetData() []*GetSyntheticTaskMonitorsResponseBodyData
	SetMsg(v string) *GetSyntheticTaskMonitorsResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetSyntheticTaskMonitorsResponseBody
	GetRequestId() *string
}

type GetSyntheticTaskMonitorsResponseBody struct {
	// The status code returned.
	//
	// 	- 1001: The request was successful.
	//
	// 	- 1002: The request failed.
	//
	// 	- 1003: Parameter errors occurred.
	//
	// 	- 1004: Authentication failed.
	//
	// 	- 1006: The task does not exist.
	//
	// 	- 1099: Internal errors occurred.
	//
	// example:
	//
	// 1001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the monitoring point.
	Data []*GetSyntheticTaskMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message that is returned when the request failed.
	//
	// example:
	//
	// null
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSyntheticTaskMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSyntheticTaskMonitorsResponseBody) GetData() []*GetSyntheticTaskMonitorsResponseBodyData {
	return s.Data
}

func (s *GetSyntheticTaskMonitorsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetSyntheticTaskMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetCode(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetData(v []*GetSyntheticTaskMonitorsResponseBodyData) *GetSyntheticTaskMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetMsg(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.Msg = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetRequestId(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskMonitorsResponseBodyData struct {
	// The task status.
	//
	// 	- 0: active
	//
	// 	- 1: busy
	//
	// example:
	//
	// 0
	Busy *int64 `json:"Busy,omitempty" xml:"Busy,omitempty"`
	// The name of the city to which the monitoring point belongs.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The ID of the city to which the monitoring point belongs.
	//
	// example:
	//
	// 1100101
	CityCode *int64 `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The client type:
	//
	// 	- 1: IDC
	//
	// 	- 2: Last mile
	//
	// example:
	//
	// 1
	ClientType *int64 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The region to which the monitoring point belongs.
	//
	// example:
	//
	// Beijing
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	// The ID of the carrier.
	//
	// example:
	//
	// 18
	NetServiceId *int64 `json:"NetServiceId,omitempty" xml:"NetServiceId,omitempty"`
	// The name of the carrier.
	//
	// example:
	//
	// XXX
	NetServiceName *string `json:"NetServiceName,omitempty" xml:"NetServiceName,omitempty"`
}

func (s GetSyntheticTaskMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetBusy() *int64 {
	return s.Busy
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetCity() *string {
	return s.City
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetCityCode() *int64 {
	return s.CityCode
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetClientType() *int64 {
	return s.ClientType
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetDistrict() *string {
	return s.District
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetNetServiceId() *int64 {
	return s.NetServiceId
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) GetNetServiceName() *string {
	return s.NetServiceName
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetBusy(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.Busy = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetCity(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.City = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetCityCode(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetClientType(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetDistrict(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.District = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetNetServiceId(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.NetServiceId = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetNetServiceName(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.NetServiceName = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
