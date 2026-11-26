// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceUsageStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSourceUsageStatisticsResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListSourceUsageStatisticsResponseBody
	GetHasMore() *bool
	SetItems(v []*ListSourceUsageStatisticsResponseBodyItems) *ListSourceUsageStatisticsResponseBody
	GetItems() []*ListSourceUsageStatisticsResponseBodyItems
	SetMessage(v string) *ListSourceUsageStatisticsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListSourceUsageStatisticsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListSourceUsageStatisticsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSourceUsageStatisticsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListSourceUsageStatisticsResponseBody
	GetTotal() *int64
}

type ListSourceUsageStatisticsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The user information.
	//
	// example:
	//
	// [{"sourceId":"src_9401","sourceName":"Product documentation","sourceType":"KNOWLEDGE_BASE","retrievalCount":12,"usageCount":5,"lastEventTime":"2026-11-05T08:00:00Z"}]
	Items []*ListSourceUsageStatisticsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of data entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSourceUsageStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSourceUsageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourceUsageStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSourceUsageStatisticsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListSourceUsageStatisticsResponseBody) GetItems() []*ListSourceUsageStatisticsResponseBodyItems {
	return s.Items
}

func (s *ListSourceUsageStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSourceUsageStatisticsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListSourceUsageStatisticsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSourceUsageStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSourceUsageStatisticsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSourceUsageStatisticsResponseBody) SetCode(v string) *ListSourceUsageStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetHasMore(v bool) *ListSourceUsageStatisticsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetItems(v []*ListSourceUsageStatisticsResponseBodyItems) *ListSourceUsageStatisticsResponseBody {
	s.Items = v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetMessage(v string) *ListSourceUsageStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetPage(v int64) *ListSourceUsageStatisticsResponseBody {
	s.Page = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetPageSize(v int64) *ListSourceUsageStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetRequestId(v string) *ListSourceUsageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) SetTotal(v int64) *ListSourceUsageStatisticsResponseBody {
	s.Total = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSourceUsageStatisticsResponseBodyItems struct {
	// The most recent reporting and ingestion time within the filtered range.
	//
	// example:
	//
	// 2024-10-02 11:20:09
	LastEventTime *string `json:"lastEventTime,omitempty" xml:"lastEventTime,omitempty"`
	// The number of deduplicated sources.
	//
	// example:
	//
	// 10
	RetrievalCount *int64 `json:"retrievalCount,omitempty" xml:"retrievalCount,omitempty"`
	// The data source ID, which is unique within the tenant.
	//
	// example:
	//
	// 8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// test123
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The skill source type.
	//
	// example:
	//
	// BUILTIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The number of used sources.
	//
	// example:
	//
	// 8
	UsageCount *int64 `json:"usageCount,omitempty" xml:"usageCount,omitempty"`
}

func (s ListSourceUsageStatisticsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSourceUsageStatisticsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetLastEventTime() *string {
	return s.LastEventTime
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetRetrievalCount() *int64 {
	return s.RetrievalCount
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetSourceId() *string {
	return s.SourceId
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetSourceName() *string {
	return s.SourceName
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSourceUsageStatisticsResponseBodyItems) GetUsageCount() *int64 {
	return s.UsageCount
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetLastEventTime(v string) *ListSourceUsageStatisticsResponseBodyItems {
	s.LastEventTime = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetRetrievalCount(v int64) *ListSourceUsageStatisticsResponseBodyItems {
	s.RetrievalCount = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetSourceId(v string) *ListSourceUsageStatisticsResponseBodyItems {
	s.SourceId = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetSourceName(v string) *ListSourceUsageStatisticsResponseBodyItems {
	s.SourceName = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetSourceType(v string) *ListSourceUsageStatisticsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) SetUsageCount(v int64) *ListSourceUsageStatisticsResponseBodyItems {
	s.UsageCount = &v
	return s
}

func (s *ListSourceUsageStatisticsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
