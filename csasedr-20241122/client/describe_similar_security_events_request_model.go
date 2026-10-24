// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarSecurityEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSimilarSecurityEventsRequest
	GetCurrentPage() *int32
	SetMaxResults(v int32) *DescribeSimilarSecurityEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSimilarSecurityEventsRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribeSimilarSecurityEventsRequest
	GetPageSize() *int32
	SetTaskId(v int64) *DescribeSimilarSecurityEventsRequest
	GetTaskId() *int64
}

type DescribeSimilarSecurityEventsRequest struct {
	// The page number of the returned results. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries to return in a single query. This parameter is reserved for compatibility. The backend currently uses CurrentPage and PageSize for pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the query. You can obtain the token from the result of the previous request.
	//
	// example:
	//
	// ****==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Settings for the number of alerting events displayed on each page in a paged query. Default value: **20**, which indicates that 20 alerting events are displayed per paging.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the query task. You can call the [CreateSimilarSecurityEventsQueryTask](~~CreateSimilarSecurityEventsQueryTask~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1689135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSimilarSecurityEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSimilarSecurityEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSimilarSecurityEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeSimilarSecurityEventsRequest) SetCurrentPage(v int32) *DescribeSimilarSecurityEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetMaxResults(v int32) *DescribeSimilarSecurityEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetNextToken(v string) *DescribeSimilarSecurityEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetPageSize(v int32) *DescribeSimilarSecurityEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetTaskId(v int64) *DescribeSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}
