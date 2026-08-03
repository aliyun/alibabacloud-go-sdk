// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupInsightEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *LookupInsightEventsRequest
	GetEndTime() *string
	SetLookupAttribute(v []*LookupInsightEventsRequestLookupAttribute) *LookupInsightEventsRequest
	GetLookupAttribute() []*LookupInsightEventsRequestLookupAttribute
	SetMaxResults(v string) *LookupInsightEventsRequest
	GetMaxResults() *string
	SetNextToken(v string) *LookupInsightEventsRequest
	GetNextToken() *string
	SetStartTime(v string) *LookupInsightEventsRequest
	GetStartTime() *string
}

type LookupInsightEventsRequest struct {
	// The end of the time range to query. The default value is the current time.
	//
	// Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2026-01-07T07:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// An array of fliter conditions.
	//
	// > - You can specify one or two query conditions. For more information, see [Limitations](https://help.aliyun.com/document_detail/3011147.html).
	LookupAttribute []*LookupInsightEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// - Valid values: 1 to 50.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// - You do not need to specify this parameter for the first request.
	//
	// - You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// VjE6dLbnNpVmbz06****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query. The default value is seven days before the current time.
	//
	// Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// > - The maximum time range that can be queried is 93 days. If the specified time range is longer than 93 days, only events from the last 93 days are returned.
	//
	// example:
	//
	// 2026-01-07T04:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupInsightEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s LookupInsightEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupInsightEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *LookupInsightEventsRequest) GetLookupAttribute() []*LookupInsightEventsRequestLookupAttribute {
	return s.LookupAttribute
}

func (s *LookupInsightEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *LookupInsightEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupInsightEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *LookupInsightEventsRequest) SetEndTime(v string) *LookupInsightEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupInsightEventsRequest) SetLookupAttribute(v []*LookupInsightEventsRequestLookupAttribute) *LookupInsightEventsRequest {
	s.LookupAttribute = v
	return s
}

func (s *LookupInsightEventsRequest) SetMaxResults(v string) *LookupInsightEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *LookupInsightEventsRequest) SetNextToken(v string) *LookupInsightEventsRequest {
	s.NextToken = &v
	return s
}

func (s *LookupInsightEventsRequest) SetStartTime(v string) *LookupInsightEventsRequest {
	s.StartTime = &v
	return s
}

func (s *LookupInsightEventsRequest) Validate() error {
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

type LookupInsightEventsRequestLookupAttribute struct {
	// The attribute key. For more information about valid values, see [How do I configure the LookupAttribute parameter when calling LookupInsightEvents?](https://help.aliyun.com/document_detail/3011147.html)
	//
	// example:
	//
	// InsightType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The attribute value. For more information about valid values, see [How do I configure the LookupAttribute parameter when calling LookupInsightEvents?](https://help.aliyun.com/document_detail/3011147.html)
	//
	// example:
	//
	// IpInsight
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LookupInsightEventsRequestLookupAttribute) String() string {
	return dara.Prettify(s)
}

func (s LookupInsightEventsRequestLookupAttribute) GoString() string {
	return s.String()
}

func (s *LookupInsightEventsRequestLookupAttribute) GetKey() *string {
	return s.Key
}

func (s *LookupInsightEventsRequestLookupAttribute) GetValue() *string {
	return s.Value
}

func (s *LookupInsightEventsRequestLookupAttribute) SetKey(v string) *LookupInsightEventsRequestLookupAttribute {
	s.Key = &v
	return s
}

func (s *LookupInsightEventsRequestLookupAttribute) SetValue(v string) *LookupInsightEventsRequestLookupAttribute {
	s.Value = &v
	return s
}

func (s *LookupInsightEventsRequestLookupAttribute) Validate() error {
	return dara.Validate(s)
}
