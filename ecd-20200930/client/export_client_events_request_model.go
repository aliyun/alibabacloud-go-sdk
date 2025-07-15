// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportClientEventsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDesktopId(v string) *ExportClientEventsRequest
  GetDesktopId() *string 
  SetDesktopName(v string) *ExportClientEventsRequest
  GetDesktopName() *string 
  SetEndTime(v string) *ExportClientEventsRequest
  GetEndTime() *string 
  SetEndUserId(v string) *ExportClientEventsRequest
  GetEndUserId() *string 
  SetEventType(v string) *ExportClientEventsRequest
  GetEventType() *string 
  SetEventTypes(v []*string) *ExportClientEventsRequest
  GetEventTypes() []*string 
  SetLangType(v string) *ExportClientEventsRequest
  GetLangType() *string 
  SetMaxResults(v int32) *ExportClientEventsRequest
  GetMaxResults() *int32 
  SetOfficeSiteId(v string) *ExportClientEventsRequest
  GetOfficeSiteId() *string 
  SetOfficeSiteName(v string) *ExportClientEventsRequest
  GetOfficeSiteName() *string 
  SetRegionId(v string) *ExportClientEventsRequest
  GetRegionId() *string 
  SetStartTime(v string) *ExportClientEventsRequest
  GetStartTime() *string 
}

type ExportClientEventsRequest struct {
  // The cloud computer ID.
  // 
  // example:
  // 
  // ecd-gx2x1dhsmucyy****
  DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
  // The cloud computer name.
  // 
  // example:
  // 
  // testName
  DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
  // The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
  // 
  // If you do not specify a value for this parameter, the current time is used.
  // 
  // example:
  // 
  // 2022-03-23T07:11:01Z
  EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  // The ID of the endpoint user.
  // 
  // example:
  // 
  // user01
  EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
  // The type of the event that you want to query. If you provide multiple values for EventTypes, the response will include events of all the specified types. If you provide no values for EventTypes and EventType, the response will include all events in the designated region.
  // 
  // Valid values:
  // 
  // 	- DESKTOP_STOP: the shutdown event.
  // 
  // 	- GET_LITE_CONNECTION_TICKET: the event of retrieving the connection ticket.
  // 
  // 	- DESKTOP_DISCONNECT: the session disconnection event.
  // 
  // 	- CLIENT_LOGIN: the user logon event.
  // 
  // 	- GET_CONNECTION_TICKET: the connection credential retrieval event.
  // 
  // 	- DESKTOP_REBOOT: the restart event.
  // 
  // 	- DESKTOP_CONNECT: the session establishment event.
  // 
  // 	- DESKTOP_START: the start event.
  // 
  // example:
  // 
  // CLIENT_LOGIN
  EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
  // The types of the events that you want to query. You can include multiple event types, and the response will return events matching the specified types or all events if none are specified.
  EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
  // The language displayed on the frontend page. The backend uses this setting to define the language of exported files.
  // 
  // Valid values:
  // 
  // 	- zh-CN: Simplified Chinese.
  // 
  // 	- en-GB: British English.
  // 
  // example:
  // 
  // zh-CN
  LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
  // The number of entries to return on each page.
  // 
  // 	- Maximum value: 5000.
  // 
  // 	- Default value: 5000.
  // 
  // example:
  // 
  // 50
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // The office network ID.
  // 
  // example:
  // 
  // cn-hangzhou+dir-363353****
  OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
  // The office network name.
  // 
  // example:
  // 
  // test
  OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
  // The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
  // 
  // If you do not specify a value for this parameter, all events that occurred before the point in time that you specify for `EndTime` are queried.
  // 
  // example:
  // 
  // 2022-03-23T04:10:21Z
  StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ExportClientEventsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportClientEventsRequest) GoString() string {
  return s.String()
}

func (s *ExportClientEventsRequest) GetDesktopId() *string  {
  return s.DesktopId
}

func (s *ExportClientEventsRequest) GetDesktopName() *string  {
  return s.DesktopName
}

func (s *ExportClientEventsRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportClientEventsRequest) GetEndUserId() *string  {
  return s.EndUserId
}

func (s *ExportClientEventsRequest) GetEventType() *string  {
  return s.EventType
}

func (s *ExportClientEventsRequest) GetEventTypes() []*string  {
  return s.EventTypes
}

func (s *ExportClientEventsRequest) GetLangType() *string  {
  return s.LangType
}

func (s *ExportClientEventsRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *ExportClientEventsRequest) GetOfficeSiteId() *string  {
  return s.OfficeSiteId
}

func (s *ExportClientEventsRequest) GetOfficeSiteName() *string  {
  return s.OfficeSiteName
}

func (s *ExportClientEventsRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportClientEventsRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportClientEventsRequest) SetDesktopId(v string) *ExportClientEventsRequest {
  s.DesktopId = &v
  return s
}

func (s *ExportClientEventsRequest) SetDesktopName(v string) *ExportClientEventsRequest {
  s.DesktopName = &v
  return s
}

func (s *ExportClientEventsRequest) SetEndTime(v string) *ExportClientEventsRequest {
  s.EndTime = &v
  return s
}

func (s *ExportClientEventsRequest) SetEndUserId(v string) *ExportClientEventsRequest {
  s.EndUserId = &v
  return s
}

func (s *ExportClientEventsRequest) SetEventType(v string) *ExportClientEventsRequest {
  s.EventType = &v
  return s
}

func (s *ExportClientEventsRequest) SetEventTypes(v []*string) *ExportClientEventsRequest {
  s.EventTypes = v
  return s
}

func (s *ExportClientEventsRequest) SetLangType(v string) *ExportClientEventsRequest {
  s.LangType = &v
  return s
}

func (s *ExportClientEventsRequest) SetMaxResults(v int32) *ExportClientEventsRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportClientEventsRequest) SetOfficeSiteId(v string) *ExportClientEventsRequest {
  s.OfficeSiteId = &v
  return s
}

func (s *ExportClientEventsRequest) SetOfficeSiteName(v string) *ExportClientEventsRequest {
  s.OfficeSiteName = &v
  return s
}

func (s *ExportClientEventsRequest) SetRegionId(v string) *ExportClientEventsRequest {
  s.RegionId = &v
  return s
}

func (s *ExportClientEventsRequest) SetStartTime(v string) *ExportClientEventsRequest {
  s.StartTime = &v
  return s
}

func (s *ExportClientEventsRequest) Validate() error {
  return dara.Validate(s)
}

