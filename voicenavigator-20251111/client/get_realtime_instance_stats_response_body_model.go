// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRealtimeInstanceStatsResponseBody
	GetCode() *string
	SetData(v *GetRealtimeInstanceStatsResponseBodyData) *GetRealtimeInstanceStatsResponseBody
	GetData() *GetRealtimeInstanceStatsResponseBodyData
	SetHttpStatusCode(v int32) *GetRealtimeInstanceStatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRealtimeInstanceStatsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetRealtimeInstanceStatsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetRealtimeInstanceStatsResponseBody
	GetRequestId() *string
}

type GetRealtimeInstanceStatsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRealtimeInstanceStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3a530dc0-7cfa-48f6-9539-bf9001e77b16
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeInstanceStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRealtimeInstanceStatsResponseBody) GetData() *GetRealtimeInstanceStatsResponseBodyData {
	return s.Data
}

func (s *GetRealtimeInstanceStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRealtimeInstanceStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRealtimeInstanceStatsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetRealtimeInstanceStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeInstanceStatsResponseBody) SetCode(v string) *GetRealtimeInstanceStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) SetData(v *GetRealtimeInstanceStatsResponseBodyData) *GetRealtimeInstanceStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) SetHttpStatusCode(v int32) *GetRealtimeInstanceStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) SetMessage(v string) *GetRealtimeInstanceStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) SetParams(v []*string) *GetRealtimeInstanceStatsResponseBody {
	s.Params = v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) SetRequestId(v string) *GetRealtimeInstanceStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRealtimeInstanceStatsResponseBodyData struct {
	// example:
	//
	// 13
	ConfiguredConcurrency   *int32                                                             `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	RealtimeScriptStatsList []*GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList `json:"RealtimeScriptStatsList,omitempty" xml:"RealtimeScriptStatsList,omitempty" type:"Repeated"`
	// example:
	//
	// 1774881208361
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 10
	UsedConcurrency *int32 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s GetRealtimeInstanceStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatsResponseBodyData) GetConfiguredConcurrency() *int32 {
	return s.ConfiguredConcurrency
}

func (s *GetRealtimeInstanceStatsResponseBodyData) GetRealtimeScriptStatsList() []*GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	return s.RealtimeScriptStatsList
}

func (s *GetRealtimeInstanceStatsResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetRealtimeInstanceStatsResponseBodyData) GetUsedConcurrency() *int32 {
	return s.UsedConcurrency
}

func (s *GetRealtimeInstanceStatsResponseBodyData) SetConfiguredConcurrency(v int32) *GetRealtimeInstanceStatsResponseBodyData {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyData) SetRealtimeScriptStatsList(v []*GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) *GetRealtimeInstanceStatsResponseBodyData {
	s.RealtimeScriptStatsList = v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyData) SetStatsTime(v int64) *GetRealtimeInstanceStatsResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyData) SetUsedConcurrency(v int32) *GetRealtimeInstanceStatsResponseBodyData {
	s.UsedConcurrency = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyData) Validate() error {
	if s.RealtimeScriptStatsList != nil {
		for _, item := range s.RealtimeScriptStatsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList struct {
	// example:
	//
	// 13
	ConfiguredConcurrency *int32 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// example:
	//
	// 92836ced-6f3a-4cec-bc3d-c3893d3c7efa
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// TypeCombinationTest
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// example:
	//
	// 1774881208361
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 10
	UsedConcurrency *int32 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GetConfiguredConcurrency() *int32 {
	return s.ConfiguredConcurrency
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) GetUsedConcurrency() *int32 {
	return s.UsedConcurrency
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) SetConfiguredConcurrency(v int32) *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) SetScriptId(v string) *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	s.ScriptId = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) SetScriptName(v string) *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	s.ScriptName = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) SetStatsTime(v int64) *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	s.StatsTime = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) SetUsedConcurrency(v int32) *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList {
	s.UsedConcurrency = &v
	return s
}

func (s *GetRealtimeInstanceStatsResponseBodyDataRealtimeScriptStatsList) Validate() error {
	return dara.Validate(s)
}
