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
	// This parameter is required.
	Date                    *string `json:"date,omitempty" xml:"date,omitempty"`
	JobOwnerListShrink      *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	PageNumber              *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize                *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
