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
	// The order in which events are retrieved. Valid values:
	//
	// - FORWARD: Chronological order.
	//
	// - BACKWARD (default): Reverse chronological order.
	//
	// example:
	//
	// BACKWARD
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > You must specify both `StartTime` and `EndTime`, or leave both unspecified. If you leave them unspecified, the default value of `EndTime` is the current time.
	//
	// example:
	//
	// 2020-10-15T11:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	//
	// > You can specify one or two filter conditions at a time. For more information, see [Limitations](https://help.aliyun.com/document_detail/2920829.html).
	LookupAttribute []*LookupEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	// The maximum number of results to return.<br>Valid values: 1 to 50.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// > You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// > You must specify both `StartTime` and `EndTime`, or leave both unspecified. If you leave them unspecified, the default value of `StartTime` is 7 days before the current time.
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
	if s.LookupAttribute != nil {
		for _, item := range s.LookupAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupEventsRequestLookupAttribute struct {
	// The attribute key. For information about valid values, see [How do I configure the LookupAttribute parameter when calling LookupInsightEvents?](https://help.aliyun.com/document_detail/2920829.html)
	//
	// example:
	//
	// ServiceName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The attribute value. For information about valid values, see [How do I configure the LookupAttribute parameter when calling LookupInsightEvents?](https://help.aliyun.com/document_detail/2920829.html)
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
