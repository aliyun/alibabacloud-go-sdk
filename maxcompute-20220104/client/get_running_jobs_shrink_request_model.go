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
	// This parameter is required.
	From                    *int64  `json:"from,omitempty" xml:"from,omitempty"`
	JobOwnerListShrink      *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	PageNumber              *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize                *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
	// This parameter is required.
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
