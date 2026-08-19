// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlerRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListCrawlerRunsRequest
	GetId() *int64
	SetPageNumber(v int32) *ListCrawlerRunsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrawlerRunsRequest
	GetPageSize() *int32
	SetStartTimeFrom(v int64) *ListCrawlerRunsRequest
	GetStartTimeFrom() *int64
	SetStartTimeTo(v int64) *ListCrawlerRunsRequest
	GetStartTimeTo() *int64
	SetStatus(v string) *ListCrawlerRunsRequest
	GetStatus() *string
}

type ListCrawlerRunsRequest struct {
	// The ID of the metadata crawler. You can call ListCrawlers to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The lower bound of the run start time, in millisecond-level UNIX timestamp. The value must be within the last 30 days. If not specified, the default value is 30 days before the current time.
	//
	// example:
	//
	// 1710239005403
	StartTimeFrom *int64 `json:"StartTimeFrom,omitempty" xml:"StartTimeFrom,omitempty"`
	// The upper bound of the run start time, in millisecond-level UNIX timestamp. The value must be within the last 30 days. If not specified, the default value is the current time.
	//
	// example:
	//
	// 1710325405403
	StartTimeTo *int64 `json:"StartTimeTo,omitempty" xml:"StartTimeTo,omitempty"`
	// The run status. Valid values: WAITING, RUNNING, SUCCESS, ERROR, SHUTDOWN.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCrawlerRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerRunsRequest) GoString() string {
	return s.String()
}

func (s *ListCrawlerRunsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListCrawlerRunsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrawlerRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrawlerRunsRequest) GetStartTimeFrom() *int64 {
	return s.StartTimeFrom
}

func (s *ListCrawlerRunsRequest) GetStartTimeTo() *int64 {
	return s.StartTimeTo
}

func (s *ListCrawlerRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCrawlerRunsRequest) SetId(v int64) *ListCrawlerRunsRequest {
	s.Id = &v
	return s
}

func (s *ListCrawlerRunsRequest) SetPageNumber(v int32) *ListCrawlerRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrawlerRunsRequest) SetPageSize(v int32) *ListCrawlerRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrawlerRunsRequest) SetStartTimeFrom(v int64) *ListCrawlerRunsRequest {
	s.StartTimeFrom = &v
	return s
}

func (s *ListCrawlerRunsRequest) SetStartTimeTo(v int64) *ListCrawlerRunsRequest {
	s.StartTimeTo = &v
	return s
}

func (s *ListCrawlerRunsRequest) SetStatus(v string) *ListCrawlerRunsRequest {
	s.Status = &v
	return s
}

func (s *ListCrawlerRunsRequest) Validate() error {
	return dara.Validate(s)
}
