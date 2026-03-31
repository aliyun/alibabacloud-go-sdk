// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetJobResourceUsageRequest
	GetDate() *string
	SetJobOwnerList(v []*string) *GetJobResourceUsageRequest
	GetJobOwnerList() []*string
	SetPageNumber(v int64) *GetJobResourceUsageRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetJobResourceUsageRequest
	GetPageSize() *int64
	SetQuotaNicknameList(v []*string) *GetJobResourceUsageRequest
	GetQuotaNicknameList() []*string
}

type GetJobResourceUsageRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-15
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
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
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
}

func (s GetJobResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageRequest) GetDate() *string {
	return s.Date
}

func (s *GetJobResourceUsageRequest) GetJobOwnerList() []*string {
	return s.JobOwnerList
}

func (s *GetJobResourceUsageRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetJobResourceUsageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetJobResourceUsageRequest) GetQuotaNicknameList() []*string {
	return s.QuotaNicknameList
}

func (s *GetJobResourceUsageRequest) SetDate(v string) *GetJobResourceUsageRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetJobOwnerList(v []*string) *GetJobResourceUsageRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageNumber(v int64) *GetJobResourceUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageSize(v int64) *GetJobResourceUsageRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetQuotaNicknameList(v []*string) *GetJobResourceUsageRequest {
	s.QuotaNicknameList = v
	return s
}

func (s *GetJobResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
