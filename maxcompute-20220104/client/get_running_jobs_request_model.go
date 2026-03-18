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
	// This parameter is required.
	From              *int64    `json:"from,omitempty" xml:"from,omitempty"`
	JobOwnerList      []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	PageNumber        *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize          *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
	// This parameter is required.
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
