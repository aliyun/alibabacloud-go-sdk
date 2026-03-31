// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResourceUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetJobResourceUsageShrinkRequest
	GetDate() *string
	SetJobOwnerListShrink(v string) *GetJobResourceUsageShrinkRequest
	GetJobOwnerListShrink() *string
	SetPageNumber(v int64) *GetJobResourceUsageShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetJobResourceUsageShrinkRequest
	GetPageSize() *int64
	SetQuotaNicknameListShrink(v string) *GetJobResourceUsageShrinkRequest
	GetQuotaNicknameListShrink() *string
}

type GetJobResourceUsageShrinkRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-15
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
}

func (s GetJobResourceUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageShrinkRequest) GetDate() *string {
	return s.Date
}

func (s *GetJobResourceUsageShrinkRequest) GetJobOwnerListShrink() *string {
	return s.JobOwnerListShrink
}

func (s *GetJobResourceUsageShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetJobResourceUsageShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetJobResourceUsageShrinkRequest) GetQuotaNicknameListShrink() *string {
	return s.QuotaNicknameListShrink
}

func (s *GetJobResourceUsageShrinkRequest) SetDate(v string) *GetJobResourceUsageShrinkRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetJobOwnerListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageNumber(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageSize(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetQuotaNicknameListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
