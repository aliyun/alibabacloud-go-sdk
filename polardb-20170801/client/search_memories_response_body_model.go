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
	Page     *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of results.
	Results    []*SearchMemoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Total      *string                              `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPages *string                              `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// 2025-09-26T08:25:44Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The memory ID.
	//
	// example:
	//
	// 423
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The memory content.
	//
	// example:
	//
	// xxx
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The agent to which the memory belongs.
	//
	// example:
	//
	// agent1
	MemoryAgentId *string `json:"MemoryAgentId,omitempty" xml:"MemoryAgentId,omitempty"`
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
	// {\\"tags\\":\\"Issue_date_2023-11-30,VD_现行有效\\"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The score.
	//
	// example:
	//
	// 12
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The update time.
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

func (s *SearchMemoriesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *SearchMemoriesResponseBodyResults) GetMemory() *string {
	return s.Memory
}

func (s *SearchMemoriesResponseBodyResults) GetMemoryAgentId() *string {
	return s.MemoryAgentId
}

func (s *SearchMemoriesResponseBodyResults) GetMemoryUserId() *string {
	return s.MemoryUserId
}

func (s *SearchMemoriesResponseBodyResults) GetMetadata() *string {
	return s.Metadata
}

func (s *SearchMemoriesResponseBodyResults) GetScore() *string {
	return s.Score
}

func (s *SearchMemoriesResponseBodyResults) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchMemoriesResponseBodyResults) SetCreateTime(v string) *SearchMemoriesResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetId(v string) *SearchMemoriesResponseBodyResults {
	s.Id = &v
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

func (s *SearchMemoriesResponseBodyResults) SetMemoryUserId(v string) *SearchMemoriesResponseBodyResults {
	s.MemoryUserId = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetMetadata(v string) *SearchMemoriesResponseBodyResults {
	s.Metadata = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetScore(v string) *SearchMemoriesResponseBodyResults {
	s.Score = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) SetUpdateTime(v string) *SearchMemoriesResponseBodyResults {
	s.UpdateTime = &v
	return s
}

func (s *SearchMemoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
