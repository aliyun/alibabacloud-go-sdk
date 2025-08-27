// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCitySearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CitySearchResponseBody
	GetCode() *string
	SetMessage(v string) *CitySearchResponseBody
	GetMessage() *string
	SetModule(v *CitySearchResponseBodyModule) *CitySearchResponseBody
	GetModule() *CitySearchResponseBodyModule
	SetRequestId(v string) *CitySearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CitySearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CitySearchResponseBody
	GetTraceId() *string
}

type CitySearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *CitySearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CitySearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CitySearchResponseBody) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CitySearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CitySearchResponseBody) GetModule() *CitySearchResponseBodyModule {
	return s.Module
}

func (s *CitySearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CitySearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CitySearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CitySearchResponseBody) SetCode(v string) *CitySearchResponseBody {
	s.Code = &v
	return s
}

func (s *CitySearchResponseBody) SetMessage(v string) *CitySearchResponseBody {
	s.Message = &v
	return s
}

func (s *CitySearchResponseBody) SetModule(v *CitySearchResponseBodyModule) *CitySearchResponseBody {
	s.Module = v
	return s
}

func (s *CitySearchResponseBody) SetRequestId(v string) *CitySearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CitySearchResponseBody) SetSuccess(v bool) *CitySearchResponseBody {
	s.Success = &v
	return s
}

func (s *CitySearchResponseBody) SetTraceId(v string) *CitySearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *CitySearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type CitySearchResponseBodyModule struct {
	Cities []*CitySearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
}

func (s CitySearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CitySearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBodyModule) GetCities() []*CitySearchResponseBodyModuleCities {
	return s.Cities
}

func (s *CitySearchResponseBodyModule) SetCities(v []*CitySearchResponseBodyModuleCities) *CitySearchResponseBodyModule {
	s.Cities = v
	return s
}

func (s *CitySearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CitySearchResponseBodyModuleCities struct {
	// example:
	//
	// 330100
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	Region *int32 `json:"region,omitempty" xml:"region,omitempty"`
}

func (s CitySearchResponseBodyModuleCities) String() string {
	return dara.Prettify(s)
}

func (s CitySearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *CitySearchResponseBodyModuleCities) GetCode() *string {
	return s.Code
}

func (s *CitySearchResponseBodyModuleCities) GetName() *string {
	return s.Name
}

func (s *CitySearchResponseBodyModuleCities) GetRegion() *int32 {
	return s.Region
}

func (s *CitySearchResponseBodyModuleCities) SetCode(v string) *CitySearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *CitySearchResponseBodyModuleCities) SetName(v string) *CitySearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

func (s *CitySearchResponseBodyModuleCities) SetRegion(v int32) *CitySearchResponseBodyModuleCities {
	s.Region = &v
	return s
}

func (s *CitySearchResponseBodyModuleCities) Validate() error {
	return dara.Validate(s)
}
