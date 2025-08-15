// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*GetAccessKeyLastUsedEventsResponseBodyEvents) *GetAccessKeyLastUsedEventsResponseBody
	GetEvents() []*GetAccessKeyLastUsedEventsResponseBodyEvents
	SetNextToken(v string) *GetAccessKeyLastUsedEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetAccessKeyLastUsedEventsResponseBody
	GetRequestId() *string
}

type GetAccessKeyLastUsedEventsResponseBody struct {
	// The list of returned events.
	//
	// This parameter is required.
	Events []*GetAccessKeyLastUsedEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The token that determines the start point of the query.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponseBody) GetEvents() []*GetAccessKeyLastUsedEventsResponseBodyEvents {
	return s.Events
}

func (s *GetAccessKeyLastUsedEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetEvents(v []*GetAccessKeyLastUsedEventsResponseBodyEvents) *GetAccessKeyLastUsedEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccessKeyLastUsedEventsResponseBodyEvents struct {
	// An array that consists of the details about the event.
	//
	// example:
	//
	// {
	//
	//   "eventId": "239EB588-CD24-522E-B0B5-174A1A58****",
	//
	//   "eventVersion": 1,
	//
	//   "eventSource": "ecs.cn-hangzhou.aliyuncs.com",
	//
	//   "sourceIpAddress": "``10.10.**.**``",
	//
	//   "eventType": "ApiCall",
	//
	//   "userIdentity": {
	//
	//     "accountId": "104758519118****",
	//
	//     "principalId": "24549429003625****",
	//
	//     "type": "ram-user",
	//
	//     "userName": "alice"
	//
	//   },
	//
	//   "serviceName": "Ecs",
	//
	//   "apiVersion": "2016-01-20",
	//
	//   "requestId": "239EB588-CD24-522E-B0B5-174A1A588BE0",
	//
	//   "eventTime": "2021-08-05T09:21:32Z",
	//
	//   "isGlobal": false,
	//
	//   "acsRegion": "cn-hangzhou",
	//
	//   "eventName": "DescribeInstances"
	//
	// }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// DescribeInstances
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The event source.
	//
	// example:
	//
	// ManagementEvent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the event was generated.
	//
	// example:
	//
	// 1657247532000
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) GetDetail() *string {
	return s.Detail
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) GetEventName() *string {
	return s.EventName
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) GetSource() *string {
	return s.Source
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) GetUsedTimestamp() *int64 {
	return s.UsedTimestamp
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetDetail(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetEventName(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.EventName = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetSource(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
