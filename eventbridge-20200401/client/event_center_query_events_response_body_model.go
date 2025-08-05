// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventCenterQueryEventsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EventCenterQueryEventsResponseBody
  GetCode() *string 
  SetData(v *EventCenterQueryEventsResponseBodyData) *EventCenterQueryEventsResponseBody
  GetData() *EventCenterQueryEventsResponseBodyData 
  SetMessage(v string) *EventCenterQueryEventsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EventCenterQueryEventsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EventCenterQueryEventsResponseBody
  GetSuccess() *bool 
}

type EventCenterQueryEventsResponseBody struct {
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EventCenterQueryEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // Specified parameter Limit is not valid.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 5DAF96FB-A4B6-548C-B999-XXXXXXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EventCenterQueryEventsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsResponseBody) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EventCenterQueryEventsResponseBody) GetData() *EventCenterQueryEventsResponseBodyData  {
  return s.Data
}

func (s *EventCenterQueryEventsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EventCenterQueryEventsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EventCenterQueryEventsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EventCenterQueryEventsResponseBody) SetCode(v string) *EventCenterQueryEventsResponseBody {
  s.Code = &v
  return s
}

func (s *EventCenterQueryEventsResponseBody) SetData(v *EventCenterQueryEventsResponseBodyData) *EventCenterQueryEventsResponseBody {
  s.Data = v
  return s
}

func (s *EventCenterQueryEventsResponseBody) SetMessage(v string) *EventCenterQueryEventsResponseBody {
  s.Message = &v
  return s
}

func (s *EventCenterQueryEventsResponseBody) SetRequestId(v string) *EventCenterQueryEventsResponseBody {
  s.RequestId = &v
  return s
}

func (s *EventCenterQueryEventsResponseBody) SetSuccess(v bool) *EventCenterQueryEventsResponseBody {
  s.Success = &v
  return s
}

func (s *EventCenterQueryEventsResponseBody) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsResponseBodyData struct {
  // example:
  // 
  // 100
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // example:
  // 
  // 100
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  Table []*EventCenterQueryEventsResponseBodyDataTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
  TimeSeries []*EventCenterQueryEventsResponseBodyDataTimeSeries `json:"TimeSeries,omitempty" xml:"TimeSeries,omitempty" type:"Repeated"`
  // example:
  // 
  // 76
  TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s EventCenterQueryEventsResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsResponseBodyData) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsResponseBodyData) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EventCenterQueryEventsResponseBodyData) GetNextToken() *string  {
  return s.NextToken
}

func (s *EventCenterQueryEventsResponseBodyData) GetTable() []*EventCenterQueryEventsResponseBodyDataTable  {
  return s.Table
}

func (s *EventCenterQueryEventsResponseBodyData) GetTimeSeries() []*EventCenterQueryEventsResponseBodyDataTimeSeries  {
  return s.TimeSeries
}

func (s *EventCenterQueryEventsResponseBodyData) GetTotalCount() *int32  {
  return s.TotalCount
}

func (s *EventCenterQueryEventsResponseBodyData) SetMaxResults(v int32) *EventCenterQueryEventsResponseBodyData {
  s.MaxResults = &v
  return s
}

func (s *EventCenterQueryEventsResponseBodyData) SetNextToken(v string) *EventCenterQueryEventsResponseBodyData {
  s.NextToken = &v
  return s
}

func (s *EventCenterQueryEventsResponseBodyData) SetTable(v []*EventCenterQueryEventsResponseBodyDataTable) *EventCenterQueryEventsResponseBodyData {
  s.Table = v
  return s
}

func (s *EventCenterQueryEventsResponseBodyData) SetTimeSeries(v []*EventCenterQueryEventsResponseBodyDataTimeSeries) *EventCenterQueryEventsResponseBodyData {
  s.TimeSeries = v
  return s
}

func (s *EventCenterQueryEventsResponseBodyData) SetTotalCount(v int32) *EventCenterQueryEventsResponseBodyData {
  s.TotalCount = &v
  return s
}

func (s *EventCenterQueryEventsResponseBodyData) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsResponseBodyDataTable struct {
  // example:
  // 
  // {
  // 
  //   "id":"xxx",
  // 
  //   "source":"acs.ecs"
  // 
  // }
  RowData map[string]interface{} `json:"RowData,omitempty" xml:"RowData,omitempty"`
}

func (s EventCenterQueryEventsResponseBodyDataTable) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsResponseBodyDataTable) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsResponseBodyDataTable) GetRowData() map[string]interface{}  {
  return s.RowData
}

func (s *EventCenterQueryEventsResponseBodyDataTable) SetRowData(v map[string]interface{}) *EventCenterQueryEventsResponseBodyDataTable {
  s.RowData = v
  return s
}

func (s *EventCenterQueryEventsResponseBodyDataTable) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsResponseBodyDataTimeSeries struct {
  // example:
  // 
  // {
  // 
  //   "id":"xxx",
  // 
  //   "source":"acs.ecs"
  // 
  // }
  RowData map[string]interface{} `json:"RowData,omitempty" xml:"RowData,omitempty"`
  // example:
  // 
  // 1683561600000
  Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s EventCenterQueryEventsResponseBodyDataTimeSeries) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsResponseBodyDataTimeSeries) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsResponseBodyDataTimeSeries) GetRowData() map[string]interface{}  {
  return s.RowData
}

func (s *EventCenterQueryEventsResponseBodyDataTimeSeries) GetTime() *string  {
  return s.Time
}

func (s *EventCenterQueryEventsResponseBodyDataTimeSeries) SetRowData(v map[string]interface{}) *EventCenterQueryEventsResponseBodyDataTimeSeries {
  s.RowData = v
  return s
}

func (s *EventCenterQueryEventsResponseBodyDataTimeSeries) SetTime(v string) *EventCenterQueryEventsResponseBodyDataTimeSeries {
  s.Time = &v
  return s
}

func (s *EventCenterQueryEventsResponseBodyDataTimeSeries) Validate() error {
  return dara.Validate(s)
}

