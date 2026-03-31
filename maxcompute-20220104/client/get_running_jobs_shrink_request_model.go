// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetRunningJobsShrinkRequest
	GetFrom() *int64
	SetJobOwnerListShrink(v string) *GetRunningJobsShrinkRequest
	GetJobOwnerListShrink() *string
	SetPageNumber(v int64) *GetRunningJobsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetRunningJobsShrinkRequest
	GetPageSize() *int64
	SetQuotaNicknameListShrink(v string) *GetRunningJobsShrinkRequest
	GetQuotaNicknameListShrink() *string
	SetTo(v int64) *GetRunningJobsShrinkRequest
	GetTo() *int64
}

type GetRunningJobsShrinkRequest struct {
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
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
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
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
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

func (s GetRunningJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsShrinkRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetRunningJobsShrinkRequest) GetJobOwnerListShrink() *string {
	return s.JobOwnerListShrink
}

func (s *GetRunningJobsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetRunningJobsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRunningJobsShrinkRequest) GetQuotaNicknameListShrink() *string {
	return s.QuotaNicknameListShrink
}

func (s *GetRunningJobsShrinkRequest) GetTo() *int64 {
	return s.To
}

func (s *GetRunningJobsShrinkRequest) SetFrom(v int64) *GetRunningJobsShrinkRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetJobOwnerListShrink(v string) *GetRunningJobsShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageNumber(v int64) *GetRunningJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageSize(v int64) *GetRunningJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetQuotaNicknameListShrink(v string) *GetRunningJobsShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetTo(v int64) *GetRunningJobsShrinkRequest {
	s.To = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
