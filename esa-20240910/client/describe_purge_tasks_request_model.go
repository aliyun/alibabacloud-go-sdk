// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurgeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribePurgeTasksRequest
	GetContent() *string
	SetEndTime(v string) *DescribePurgeTasksRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribePurgeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePurgeTasksRequest
	GetPageSize() *int32
	SetSiteId(v int64) *DescribePurgeTasksRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribePurgeTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribePurgeTasksRequest
	GetStatus() *string
	SetType(v string) *DescribePurgeTasksRequest
	GetType() *string
}

type DescribePurgeTasksRequest struct {
	// The content to purge. Exact match is supported.
	//
	// example:
	//
	// http://a.com/1.jpg?b=1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2022-11-18T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Valid values: 1 to 100000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-11-16T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Refreshing**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type. Valid values:
	//
	// 	- **file*	- (default): purges the cache by file.
	//
	// 	- **cachetag**: purges the cache by cache tag.
	//
	// 	- **directory**: purges the cache by directory.
	//
	// 	- **ignoreParams**: purges the cache by URL with specified parameters ignored.
	//
	// 	- **hostname**: purges the cache by hostname.
	//
	// 	- **purgeall**: purges all cache.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePurgeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurgeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksRequest) GetContent() *string {
	return s.Content
}

func (s *DescribePurgeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePurgeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePurgeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePurgeTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribePurgeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePurgeTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribePurgeTasksRequest) GetType() *string {
	return s.Type
}

func (s *DescribePurgeTasksRequest) SetContent(v string) *DescribePurgeTasksRequest {
	s.Content = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetEndTime(v string) *DescribePurgeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageNumber(v int32) *DescribePurgeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageSize(v int32) *DescribePurgeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetSiteId(v int64) *DescribePurgeTasksRequest {
	s.SiteId = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStartTime(v string) *DescribePurgeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStatus(v string) *DescribePurgeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetType(v string) *DescribePurgeTasksRequest {
	s.Type = &v
	return s
}

func (s *DescribePurgeTasksRequest) Validate() error {
	return dara.Validate(s)
}
