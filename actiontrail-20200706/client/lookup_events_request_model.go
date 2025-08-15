// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *LookupEventsRequest
	GetDirection() *string
	SetEndTime(v string) *LookupEventsRequest
	GetEndTime() *string
	SetLookupAttribute(v []*LookupEventsRequestLookupAttribute) *LookupEventsRequest
	GetLookupAttribute() []*LookupEventsRequestLookupAttribute
	SetMaxResults(v string) *LookupEventsRequest
	GetMaxResults() *string
	SetNextToken(v string) *LookupEventsRequest
	GetNextToken() *string
	SetStartTime(v string) *LookupEventsRequest
	GetStartTime() *string
}

type LookupEventsRequest struct {
	// The order in which details of events are to be retrieved. Valid values:
	//
	// 	- FORWARD: ascending order.
	//
	// 	- BACKWARD: descending order. This is the default value.
	//
	// example:
	//
	// BACKWARD
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. The default time is the current time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2020-10-15T11:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query conditions.
	LookupAttribute []*LookupEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	// The maximum number of entries to be returned.
	//
	// Valid values: 0 to 50.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to request the next page of query results.
	//
	// > The request parameters must be the same as those of the last request.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query. The default time is seven days prior to the current time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2020-10-08T11:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s LookupEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupEventsRequest) GetDirection() *string {
	return s.Direction
}

func (s *LookupEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *LookupEventsRequest) GetLookupAttribute() []*LookupEventsRequestLookupAttribute {
	return s.LookupAttribute
}

func (s *LookupEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *LookupEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *LookupEventsRequest) SetDirection(v string) *LookupEventsRequest {
	s.Direction = &v
	return s
}

func (s *LookupEventsRequest) SetEndTime(v string) *LookupEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupEventsRequest) SetLookupAttribute(v []*LookupEventsRequestLookupAttribute) *LookupEventsRequest {
	s.LookupAttribute = v
	return s
}

func (s *LookupEventsRequest) SetMaxResults(v string) *LookupEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *LookupEventsRequest) SetNextToken(v string) *LookupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *LookupEventsRequest) SetStartTime(v string) *LookupEventsRequest {
	s.StartTime = &v
	return s
}

func (s *LookupEventsRequest) Validate() error {
	return dara.Validate(s)
}

type LookupEventsRequestLookupAttribute struct {
	// The key of the query condition. Valid values:
	//
	// 	- ServiceName: the name of a specific Alibaba Cloud service.
	//
	// 	- EventName: the name of a specific event.
	//
	// 	- User: the name of the RAM user who calls a specific operation.
	//
	// 	- EventId: the ID of a specific event.
	//
	// 	- ResourceType: the type of resources.
	//
	// 	- ResourceName: the name of a specific resource.
	//
	// 	- EventRW: the read/write type of events.
	//
	// 	- EventAccessKeyId: the AccessKey ID used in events.
	//
	// > You can use only one query condition for each query.
	//
	// example:
	//
	// ServiceName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query condition. Valid values:
	//
	// 	- When the LookupAttribute.N.Key parameter is set to ServiceName, you can set this parameter to a value such as `Ecs`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to EventName, you can set this parameter to a value such as `ConsoleSignin`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to User, you can set this parameter to a value such as `Alice`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to EventId, you can set this parameter to a value such as `B702AFA3-FD4B-40E3-88E4-C0752FAA****`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to ResourceType, you can set this parameter to a value such as `ACS::ECS::Instance`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to ResourceName, you can set this parameter to a value such as `i-bp14664y88udkt45****`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to EventRW, you can set this parameter to `Read` or `Write`.
	//
	// 	- When the LookupAttribute.N.Key parameter is set to EventAccessKeyId, you can set this parameter to a value such as `LTAI****************`.
	//
	// example:
	//
	// Ecs
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LookupEventsRequestLookupAttribute) String() string {
	return dara.Prettify(s)
}

func (s LookupEventsRequestLookupAttribute) GoString() string {
	return s.String()
}

func (s *LookupEventsRequestLookupAttribute) GetKey() *string {
	return s.Key
}

func (s *LookupEventsRequestLookupAttribute) GetValue() *string {
	return s.Value
}

func (s *LookupEventsRequestLookupAttribute) SetKey(v string) *LookupEventsRequestLookupAttribute {
	s.Key = &v
	return s
}

func (s *LookupEventsRequestLookupAttribute) SetValue(v string) *LookupEventsRequestLookupAttribute {
	s.Value = &v
	return s
}

func (s *LookupEventsRequestLookupAttribute) Validate() error {
	return dara.Validate(s)
}
