// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v string) *SearchMemoriesResponseBody
	GetPage() *string
	SetPageSize(v string) *SearchMemoriesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *SearchMemoriesResponseBody
	GetRequestId() *string
	SetResults(v []*SearchMemoriesResponseBodyResults) *SearchMemoriesResponseBody
	GetResults() []*SearchMemoriesResponseBodyResults
	SetTotal(v string) *SearchMemoriesResponseBody
	GetTotal() *string
	SetTotalPages(v string) *SearchMemoriesResponseBody
	GetTotalPages() *string
}

type SearchMemoriesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of results.
	Results []*SearchMemoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of records that meet the specified conditions.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *string `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBody) GetPage() *string {
	return s.Page
}

func (s *SearchMemoriesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *SearchMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMemoriesResponseBody) GetResults() []*SearchMemoriesResponseBodyResults {
	return s.Results
}

func (s *SearchMemoriesResponseBody) GetTotal() *string {
	return s.Total
}

func (s *SearchMemoriesResponseBody) GetTotalPages() *string {
	return s.TotalPages
}

func (s *SearchMemoriesResponseBody) SetPage(v string) *SearchMemoriesResponseBody {
	s.Page = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetPageSize(v string) *SearchMemoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetRequestId(v string) *SearchMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetResults(v []*SearchMemoriesResponseBodyResults) *SearchMemoriesResponseBody {
	s.Results = v
	return s
}

func (s *SearchMemoriesResponseBody) SetTotal(v string) *SearchMemoriesResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetTotalPages(v string) *SearchMemoriesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchMemoriesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMemoriesResponseBodyResults struct {
	// The creation time.
	//
	// The time is in ISO 8601 format and follows the UTC time zone (format: YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2025-09-26T08:25:44Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event information.
	//
	// example:
	//
	// abc
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The memory ID.
	//
	// example:
	//
	// 423
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The first input item in the result list.
	Input []interface{} `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// The memory content.
	//
	// example:
	//
	// The user prefers PostgreSQL
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The agent to which the memory belongs.
	//
	// example:
	//
	// agent1
	MemoryAgentId *string `json:"MemoryAgentId,omitempty" xml:"MemoryAgentId,omitempty"`
	// The first memory ID in the result set.
	//
	// example:
	//
	// abc
	MemoryId *string `json:"MemoryId,omitempty" xml:"MemoryId,omitempty"`
	// The user to whom the memory belongs.
	//
	// example:
	//
	// user1
	MemoryUserId *string `json:"MemoryUserId,omitempty" xml:"MemoryUserId,omitempty"`
	// The metadata.
	//
	// example:
	//
	// {"tags":"Issue_date_2023-11-30,VD_Active"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The first new memory information in the result.
	//
	// example:
	//
	// abc
	NewMemory *string `json:"NewMemory,omitempty" xml:"NewMemory,omitempty"`
	// The first old memory value in the result.
	//
	// example:
	//
	// abc
	OldMemory *string `json:"OldMemory,omitempty" xml:"OldMemory,omitempty"`
	// The score.
	//
	// example:
	//
	// 12
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The number of update records.
	//
	// example:
	//
	// 1
	UpdateCount *int32 `json:"UpdateCount,omitempty" xml:"UpdateCount,omitempty"`
	// The update time.
	//
	// The time is in ISO 8601 format and follows the UTC time zone (format: YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2025-10-16T02:27:33Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s SearchMemoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyResults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMemoriesResponseBodyResults) GetEvent() *string {
	return s.Event
}

func (s *SearchMemoriesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *SearchMemoriesResponseBodyResults) GetInput() []interface{} {
	return s.Input
}

func (s *SearchMemoriesResponseBodyResults) GetMemory() *string {
	return s.Memory
}

func (s *SearchMemoriesResponseBodyResults) GetMemoryAgentId() *string {
	return s.MemoryAgentId
}

func (s *SearchMemoriesResponseBodyResults) GetMemoryId() *string {
	return s.MemoryId
}

func (s *SearchMemoriesResponseBodyResults) GetMemoryUserId() *string {
	return s.MemoryUserId
}

func (s *SearchMemoriesResponseBodyResults) GetMetadata() *string {
	return s.Metadata
}

func (s *SearchMemoriesResponseBodyResults) GetNewMemory() *string {
	return s.NewMemory
}

func (s *SearchMemoriesResponseBodyResults) GetOldMemory() *string {
	return s.OldMemory
}

func (s *SearchMemoriesResponseBodyResults) GetScore() *string {
	return s.Score
}

func (s *SearchMemoriesResponseBodyResults) GetUpdateCount() *int32 {
	return s.UpdateCount
}

func (s *SearchMemoriesResponseBodyResults) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchMemoriesResponseBodyResults) SetCreateTime(v string) *SearchMemoriesResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetEvent(v string) *SearchMemoriesResponseBodyResults {
	s.Event = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetId(v string) *SearchMemoriesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetInput(v []interface{}) *SearchMemoriesResponseBodyResults {
	s.Input = v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMemory(v string) *SearchMemoriesResponseBodyResults {
	s.Memory = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMemoryAgentId(v string) *SearchMemoriesResponseBodyResults {
	s.MemoryAgentId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMemoryId(v string) *SearchMemoriesResponseBodyResults {
	s.MemoryId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMemoryUserId(v string) *SearchMemoriesResponseBodyResults {
	s.MemoryUserId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMetadata(v string) *SearchMemoriesResponseBodyResults {
	s.Metadata = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetNewMemory(v string) *SearchMemoriesResponseBodyResults {
	s.NewMemory = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetOldMemory(v string) *SearchMemoriesResponseBodyResults {
	s.OldMemory = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetScore(v string) *SearchMemoriesResponseBodyResults {
	s.Score = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetUpdateCount(v int32) *SearchMemoriesResponseBodyResults {
	s.UpdateCount = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetUpdateTime(v string) *SearchMemoriesResponseBodyResults {
	s.UpdateTime = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
