// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStopoverSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainStopoverSearchResponseBody
	GetCode() *string
	SetMessage(v string) *TrainStopoverSearchResponseBody
	GetMessage() *string
	SetModule(v []*TrainStopoverSearchResponseBodyModule) *TrainStopoverSearchResponseBody
	GetModule() []*TrainStopoverSearchResponseBodyModule
	SetRequestId(v string) *TrainStopoverSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainStopoverSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainStopoverSearchResponseBody
	GetTraceId() *string
}

type TrainStopoverSearchResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module []*TrainStopoverSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainStopoverSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainStopoverSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TrainStopoverSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainStopoverSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainStopoverSearchResponseBody) GetModule() []*TrainStopoverSearchResponseBodyModule {
	return s.Module
}

func (s *TrainStopoverSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainStopoverSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainStopoverSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainStopoverSearchResponseBody) SetCode(v string) *TrainStopoverSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TrainStopoverSearchResponseBody) SetMessage(v string) *TrainStopoverSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TrainStopoverSearchResponseBody) SetModule(v []*TrainStopoverSearchResponseBodyModule) *TrainStopoverSearchResponseBody {
	s.Module = v
	return s
}

func (s *TrainStopoverSearchResponseBody) SetRequestId(v string) *TrainStopoverSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainStopoverSearchResponseBody) SetSuccess(v bool) *TrainStopoverSearchResponseBody {
	s.Success = &v
	return s
}

func (s *TrainStopoverSearchResponseBody) SetTraceId(v string) *TrainStopoverSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainStopoverSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainStopoverSearchResponseBodyModule struct {
	// example:
	//
	// 2024-05-06 15:19:01
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime     *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	StationName *string `json:"station_name,omitempty" xml:"station_name,omitempty"`
	// example:
	//
	// 0
	StationNo *string `json:"station_no,omitempty" xml:"station_no,omitempty"`
	// example:
	//
	// 0
	StationType *string `json:"station_type,omitempty" xml:"station_type,omitempty"`
	// example:
	//
	// 22
	StopOverTime *string `json:"stop_over_time,omitempty" xml:"stop_over_time,omitempty"`
}

func (s TrainStopoverSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainStopoverSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainStopoverSearchResponseBodyModule) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainStopoverSearchResponseBodyModule) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainStopoverSearchResponseBodyModule) GetStationName() *string {
	return s.StationName
}

func (s *TrainStopoverSearchResponseBodyModule) GetStationNo() *string {
	return s.StationNo
}

func (s *TrainStopoverSearchResponseBodyModule) GetStationType() *string {
	return s.StationType
}

func (s *TrainStopoverSearchResponseBodyModule) GetStopOverTime() *string {
	return s.StopOverTime
}

func (s *TrainStopoverSearchResponseBodyModule) SetArrTime(v string) *TrainStopoverSearchResponseBodyModule {
	s.ArrTime = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) SetDepTime(v string) *TrainStopoverSearchResponseBodyModule {
	s.DepTime = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) SetStationName(v string) *TrainStopoverSearchResponseBodyModule {
	s.StationName = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) SetStationNo(v string) *TrainStopoverSearchResponseBodyModule {
	s.StationNo = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) SetStationType(v string) *TrainStopoverSearchResponseBodyModule {
	s.StationType = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) SetStopOverTime(v string) *TrainStopoverSearchResponseBodyModule {
	s.StopOverTime = &v
	return s
}

func (s *TrainStopoverSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
