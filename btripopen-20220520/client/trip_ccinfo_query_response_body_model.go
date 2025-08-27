// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripCCInfoQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TripCCInfoQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TripCCInfoQueryResponseBody
	GetMessage() *string
	SetModule(v []*TripCCInfoQueryResponseBodyModule) *TripCCInfoQueryResponseBody
	GetModule() []*TripCCInfoQueryResponseBodyModule
	SetRequestId(v string) *TripCCInfoQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TripCCInfoQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TripCCInfoQueryResponseBody
	GetTraceId() *string
}

type TripCCInfoQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module []*TripCCInfoQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s TripCCInfoQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TripCCInfoQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TripCCInfoQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TripCCInfoQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TripCCInfoQueryResponseBody) GetModule() []*TripCCInfoQueryResponseBodyModule {
	return s.Module
}

func (s *TripCCInfoQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TripCCInfoQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TripCCInfoQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TripCCInfoQueryResponseBody) SetCode(v string) *TripCCInfoQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TripCCInfoQueryResponseBody) SetMessage(v string) *TripCCInfoQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TripCCInfoQueryResponseBody) SetModule(v []*TripCCInfoQueryResponseBodyModule) *TripCCInfoQueryResponseBody {
	s.Module = v
	return s
}

func (s *TripCCInfoQueryResponseBody) SetRequestId(v string) *TripCCInfoQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TripCCInfoQueryResponseBody) SetSuccess(v bool) *TripCCInfoQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TripCCInfoQueryResponseBody) SetTraceId(v string) *TripCCInfoQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TripCCInfoQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TripCCInfoQueryResponseBodyModule struct {
	// example:
	//
	// user_12138
	Notifier *string `json:"notifier,omitempty" xml:"notifier,omitempty"`
	// example:
	//
	// 1525104000
	NotifyStartTime *int64 `json:"notify_start_time,omitempty" xml:"notify_start_time,omitempty"`
}

func (s TripCCInfoQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TripCCInfoQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TripCCInfoQueryResponseBodyModule) GetNotifier() *string {
	return s.Notifier
}

func (s *TripCCInfoQueryResponseBodyModule) GetNotifyStartTime() *int64 {
	return s.NotifyStartTime
}

func (s *TripCCInfoQueryResponseBodyModule) SetNotifier(v string) *TripCCInfoQueryResponseBodyModule {
	s.Notifier = &v
	return s
}

func (s *TripCCInfoQueryResponseBodyModule) SetNotifyStartTime(v int64) *TripCCInfoQueryResponseBodyModule {
	s.NotifyStartTime = &v
	return s
}

func (s *TripCCInfoQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
