// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAirportSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AirportSearchResponseBody
	GetCode() *string
	SetMessage(v string) *AirportSearchResponseBody
	GetMessage() *string
	SetModule(v *AirportSearchResponseBodyModule) *AirportSearchResponseBody
	GetModule() *AirportSearchResponseBodyModule
	SetRequestId(v string) *AirportSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AirportSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AirportSearchResponseBody
	GetTraceId() *string
}

type AirportSearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *AirportSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f07f316603757445272547d959f
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AirportSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchResponseBody) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *AirportSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AirportSearchResponseBody) GetModule() *AirportSearchResponseBodyModule {
	return s.Module
}

func (s *AirportSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AirportSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AirportSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AirportSearchResponseBody) SetCode(v string) *AirportSearchResponseBody {
	s.Code = &v
	return s
}

func (s *AirportSearchResponseBody) SetMessage(v string) *AirportSearchResponseBody {
	s.Message = &v
	return s
}

func (s *AirportSearchResponseBody) SetModule(v *AirportSearchResponseBodyModule) *AirportSearchResponseBody {
	s.Module = v
	return s
}

func (s *AirportSearchResponseBody) SetRequestId(v string) *AirportSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *AirportSearchResponseBody) SetSuccess(v bool) *AirportSearchResponseBody {
	s.Success = &v
	return s
}

func (s *AirportSearchResponseBody) SetTraceId(v string) *AirportSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *AirportSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type AirportSearchResponseBodyModule struct {
	Cities []*AirportSearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Nearby *bool `json:"nearby,omitempty" xml:"nearby,omitempty"`
}

func (s AirportSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBodyModule) GetCities() []*AirportSearchResponseBodyModuleCities {
	return s.Cities
}

func (s *AirportSearchResponseBodyModule) GetNearby() *bool {
	return s.Nearby
}

func (s *AirportSearchResponseBodyModule) SetCities(v []*AirportSearchResponseBodyModuleCities) *AirportSearchResponseBodyModule {
	s.Cities = v
	return s
}

func (s *AirportSearchResponseBodyModule) SetNearby(v bool) *AirportSearchResponseBodyModule {
	s.Nearby = &v
	return s
}

func (s *AirportSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type AirportSearchResponseBodyModuleCities struct {
	// example:
	//
	// HGH
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 100
	Distance *int32  `json:"distance,omitempty" xml:"distance,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 上海
	TravelName *string `json:"travel_name,omitempty" xml:"travel_name,omitempty"`
}

func (s AirportSearchResponseBodyModuleCities) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *AirportSearchResponseBodyModuleCities) GetCode() *string {
	return s.Code
}

func (s *AirportSearchResponseBodyModuleCities) GetDistance() *int32 {
	return s.Distance
}

func (s *AirportSearchResponseBodyModuleCities) GetName() *string {
	return s.Name
}

func (s *AirportSearchResponseBodyModuleCities) GetTravelName() *string {
	return s.TravelName
}

func (s *AirportSearchResponseBodyModuleCities) SetCode(v string) *AirportSearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetDistance(v int32) *AirportSearchResponseBodyModuleCities {
	s.Distance = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetName(v string) *AirportSearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) SetTravelName(v string) *AirportSearchResponseBodyModuleCities {
	s.TravelName = &v
	return s
}

func (s *AirportSearchResponseBodyModuleCities) Validate() error {
	return dara.Validate(s)
}
