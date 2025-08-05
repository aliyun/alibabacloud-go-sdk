// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetRunningJobsRequest
	GetFrom() *int64
	SetJobOwnerList(v []*string) *GetRunningJobsRequest
	GetJobOwnerList() []*string
	SetPageNumber(v int64) *GetRunningJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetRunningJobsRequest
	GetPageSize() *int64
	SetQuotaNicknameList(v []*string) *GetRunningJobsRequest
	GetQuotaNicknameList() []*string
	SetTo(v int64) *GetRunningJobsRequest
	GetTo() *int64
}

type GetRunningJobsRequest struct {
	// The time when the query starts. This parameter specifies the time when a job is submitted.
	//
	// 	- The time range that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the range. If the value of **from*	- is the same as the value of **to**, the time range is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683785928
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
	// The time when the query ends. This parameter specifies the time when a job is submitted.
	//
	// 	- The time interval that is specified by the **from*	- and **to*	- request parameters is a closed interval. The start time and end time are included in the interval. If the value of **from*	- is the same as the value of **to**, the interval is invalid, and a null value is returned.
	//
	// 	- The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683612946
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetRunningJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetRunningJobsRequest) GetJobOwnerList() []*string {
	return s.JobOwnerList
}

func (s *GetRunningJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetRunningJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRunningJobsRequest) GetQuotaNicknameList() []*string {
	return s.QuotaNicknameList
}

func (s *GetRunningJobsRequest) GetTo() *int64 {
	return s.To
}

func (s *GetRunningJobsRequest) SetFrom(v int64) *GetRunningJobsRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsRequest) SetJobOwnerList(v []*string) *GetRunningJobsRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetRunningJobsRequest) SetPageNumber(v int64) *GetRunningJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsRequest) SetPageSize(v int64) *GetRunningJobsRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsRequest) SetQuotaNicknameList(v []*string) *GetRunningJobsRequest {
	s.QuotaNicknameList = v
	return s
}

func (s *GetRunningJobsRequest) SetTo(v int64) *GetRunningJobsRequest {
	s.To = &v
	return s
}

func (s *GetRunningJobsRequest) Validate() error {
	return dara.Validate(s)
}
