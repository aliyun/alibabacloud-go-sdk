// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeScriptStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRealtimeScriptStatsResponseBody
	GetCode() *string
	SetData(v *GetRealtimeScriptStatsResponseBodyData) *GetRealtimeScriptStatsResponseBody
	GetData() *GetRealtimeScriptStatsResponseBodyData
	SetHttpStatusCode(v int32) *GetRealtimeScriptStatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRealtimeScriptStatsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetRealtimeScriptStatsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetRealtimeScriptStatsResponseBody
	GetRequestId() *string
}

type GetRealtimeScriptStatsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRealtimeScriptStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeScriptStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeScriptStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeScriptStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRealtimeScriptStatsResponseBody) GetData() *GetRealtimeScriptStatsResponseBodyData {
	return s.Data
}

func (s *GetRealtimeScriptStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRealtimeScriptStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRealtimeScriptStatsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetRealtimeScriptStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeScriptStatsResponseBody) SetCode(v string) *GetRealtimeScriptStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) SetData(v *GetRealtimeScriptStatsResponseBodyData) *GetRealtimeScriptStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) SetHttpStatusCode(v int32) *GetRealtimeScriptStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) SetMessage(v string) *GetRealtimeScriptStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) SetParams(v []*string) *GetRealtimeScriptStatsResponseBody {
	s.Params = v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) SetRequestId(v string) *GetRealtimeScriptStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRealtimeScriptStatsResponseBodyData struct {
	// example:
	//
	// 13
	ConfiguredConcurrency *int32 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// example:
	//
	// 1774881658931
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 10
	UsedConcurrency *int32 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s GetRealtimeScriptStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeScriptStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeScriptStatsResponseBodyData) GetConfiguredConcurrency() *int32 {
	return s.ConfiguredConcurrency
}

func (s *GetRealtimeScriptStatsResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetRealtimeScriptStatsResponseBodyData) GetUsedConcurrency() *int32 {
	return s.UsedConcurrency
}

func (s *GetRealtimeScriptStatsResponseBodyData) SetConfiguredConcurrency(v int32) *GetRealtimeScriptStatsResponseBodyData {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBodyData) SetStatsTime(v int64) *GetRealtimeScriptStatsResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBodyData) SetUsedConcurrency(v int32) *GetRealtimeScriptStatsResponseBodyData {
	s.UsedConcurrency = &v
	return s
}

func (s *GetRealtimeScriptStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
