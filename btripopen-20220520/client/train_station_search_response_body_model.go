// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStationSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainStationSearchResponseBody
	GetCode() *string
	SetMessage(v string) *TrainStationSearchResponseBody
	GetMessage() *string
	SetModule(v *TrainStationSearchResponseBodyModule) *TrainStationSearchResponseBody
	GetModule() *TrainStationSearchResponseBodyModule
	SetRequestId(v string) *TrainStationSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainStationSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainStationSearchResponseBody
	GetTraceId() *string
}

type TrainStationSearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainStationSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e847f16611516748613869de4f6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainStationSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainStationSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainStationSearchResponseBody) GetModule() *TrainStationSearchResponseBodyModule {
	return s.Module
}

func (s *TrainStationSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainStationSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainStationSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainStationSearchResponseBody) SetCode(v string) *TrainStationSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetMessage(v string) *TrainStationSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetModule(v *TrainStationSearchResponseBodyModule) *TrainStationSearchResponseBody {
	s.Module = v
	return s
}

func (s *TrainStationSearchResponseBody) SetRequestId(v string) *TrainStationSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetSuccess(v bool) *TrainStationSearchResponseBody {
	s.Success = &v
	return s
}

func (s *TrainStationSearchResponseBody) SetTraceId(v string) *TrainStationSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainStationSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainStationSearchResponseBodyModule struct {
	Cities []*TrainStationSearchResponseBodyModuleCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
}

func (s TrainStationSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBodyModule) GetCities() []*TrainStationSearchResponseBodyModuleCities {
	return s.Cities
}

func (s *TrainStationSearchResponseBodyModule) SetCities(v []*TrainStationSearchResponseBodyModuleCities) *TrainStationSearchResponseBodyModule {
	s.Cities = v
	return s
}

func (s *TrainStationSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainStationSearchResponseBodyModuleCities struct {
	// example:
	//
	// hz
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TrainStationSearchResponseBodyModuleCities) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchResponseBodyModuleCities) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponseBodyModuleCities) GetCode() *string {
	return s.Code
}

func (s *TrainStationSearchResponseBodyModuleCities) GetName() *string {
	return s.Name
}

func (s *TrainStationSearchResponseBodyModuleCities) SetCode(v string) *TrainStationSearchResponseBodyModuleCities {
	s.Code = &v
	return s
}

func (s *TrainStationSearchResponseBodyModuleCities) SetName(v string) *TrainStationSearchResponseBodyModuleCities {
	s.Name = &v
	return s
}

func (s *TrainStationSearchResponseBodyModuleCities) Validate() error {
	return dara.Validate(s)
}
