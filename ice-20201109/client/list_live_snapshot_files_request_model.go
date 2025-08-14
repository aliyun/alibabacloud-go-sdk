// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLiveSnapshotFilesRequest
	GetEndTime() *string
	SetJobId(v string) *ListLiveSnapshotFilesRequest
	GetJobId() *string
	SetLimit(v int32) *ListLiveSnapshotFilesRequest
	GetLimit() *int32
	SetSortBy(v string) *ListLiveSnapshotFilesRequest
	GetSortBy() *string
	SetStartTime(v string) *ListLiveSnapshotFilesRequest
	GetStartTime() *string
}

type ListLiveSnapshotFilesRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- The maximum time range that can be specified is one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-02T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the snapshot job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of results to return each time. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The sorting order. Default value: asc.
	//
	// Valid values:
	//
	// 	- asc: sorts the query results by creation time in ascending order.
	//
	// 	- desc: sorts the query results by creation time in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-02T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLiveSnapshotFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotFilesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveSnapshotFilesRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListLiveSnapshotFilesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLiveSnapshotFilesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveSnapshotFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveSnapshotFilesRequest) SetEndTime(v string) *ListLiveSnapshotFilesRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveSnapshotFilesRequest) SetJobId(v string) *ListLiveSnapshotFilesRequest {
	s.JobId = &v
	return s
}

func (s *ListLiveSnapshotFilesRequest) SetLimit(v int32) *ListLiveSnapshotFilesRequest {
	s.Limit = &v
	return s
}

func (s *ListLiveSnapshotFilesRequest) SetSortBy(v string) *ListLiveSnapshotFilesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveSnapshotFilesRequest) SetStartTime(v string) *ListLiveSnapshotFilesRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveSnapshotFilesRequest) Validate() error {
	return dara.Validate(s)
}
