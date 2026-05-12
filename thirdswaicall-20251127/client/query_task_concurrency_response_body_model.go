// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTaskConcurrencyResponseBody
	GetCode() *string
	SetData(v *QueryTaskConcurrencyResponseBodyData) *QueryTaskConcurrencyResponseBody
	GetData() *QueryTaskConcurrencyResponseBodyData
	SetMessage(v string) *QueryTaskConcurrencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTaskConcurrencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTaskConcurrencyResponseBody
	GetSuccess() *bool
	SetTimestamp(v string) *QueryTaskConcurrencyResponseBody
	GetTimestamp() *string
	SetTraceId(v string) *QueryTaskConcurrencyResponseBody
	GetTraceId() *string
}

type QueryTaskConcurrencyResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryTaskConcurrencyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 202BFA44-28D8-571E-B992-BA70F2E92FB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1743387963
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// Trace ID。
	//
	// example:
	//
	// F47D4976-FC5A-5687-A890-B7923D3B429B
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s QueryTaskConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskConcurrencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTaskConcurrencyResponseBody) GetData() *QueryTaskConcurrencyResponseBodyData {
	return s.Data
}

func (s *QueryTaskConcurrencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTaskConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskConcurrencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTaskConcurrencyResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *QueryTaskConcurrencyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryTaskConcurrencyResponseBody) SetCode(v string) *QueryTaskConcurrencyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetData(v *QueryTaskConcurrencyResponseBodyData) *QueryTaskConcurrencyResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetMessage(v string) *QueryTaskConcurrencyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetRequestId(v string) *QueryTaskConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetSuccess(v bool) *QueryTaskConcurrencyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetTimestamp(v string) *QueryTaskConcurrencyResponseBody {
	s.Timestamp = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) SetTraceId(v string) *QueryTaskConcurrencyResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskConcurrencyResponseBodyData struct {
	// example:
	//
	// 5
	AvailableConcurrency *int64 `json:"AvailableConcurrency,omitempty" xml:"AvailableConcurrency,omitempty"`
	// example:
	//
	// 5
	CurrentConcurrency *int64 `json:"CurrentConcurrency,omitempty" xml:"CurrentConcurrency,omitempty"`
	// example:
	//
	// 5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
}

func (s QueryTaskConcurrencyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskConcurrencyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskConcurrencyResponseBodyData) GetAvailableConcurrency() *int64 {
	return s.AvailableConcurrency
}

func (s *QueryTaskConcurrencyResponseBodyData) GetCurrentConcurrency() *int64 {
	return s.CurrentConcurrency
}

func (s *QueryTaskConcurrencyResponseBodyData) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *QueryTaskConcurrencyResponseBodyData) SetAvailableConcurrency(v int64) *QueryTaskConcurrencyResponseBodyData {
	s.AvailableConcurrency = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBodyData) SetCurrentConcurrency(v int64) *QueryTaskConcurrencyResponseBodyData {
	s.CurrentConcurrency = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBodyData) SetMaxConcurrency(v int64) *QueryTaskConcurrencyResponseBodyData {
	s.MaxConcurrency = &v
	return s
}

func (s *QueryTaskConcurrencyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
