// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedEndAt(v string) *CarApplyQueryRequest
	GetCreatedEndAt() *string
	SetCreatedStartAt(v string) *CarApplyQueryRequest
	GetCreatedStartAt() *string
	SetPageNumber(v int32) *CarApplyQueryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *CarApplyQueryRequest
	GetPageSize() *int32
	SetThirdPartApplyId(v string) *CarApplyQueryRequest
	GetThirdPartApplyId() *string
	SetUserId(v string) *CarApplyQueryRequest
	GetUserId() *string
}

type CarApplyQueryRequest struct {
	// example:
	//
	// 2021-03-18 20:26:56
	CreatedEndAt *string `json:"created_end_at,omitempty" xml:"created_end_at,omitempty"`
	// example:
	//
	// 2021-03-18 20:26:56
	CreatedStartAt *string `json:"created_start_at,omitempty" xml:"created_start_at,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// apply1
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// userid
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *CarApplyQueryRequest) GetCreatedEndAt() *string {
	return s.CreatedEndAt
}

func (s *CarApplyQueryRequest) GetCreatedStartAt() *string {
	return s.CreatedStartAt
}

func (s *CarApplyQueryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CarApplyQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CarApplyQueryRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *CarApplyQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyQueryRequest) SetCreatedEndAt(v string) *CarApplyQueryRequest {
	s.CreatedEndAt = &v
	return s
}

func (s *CarApplyQueryRequest) SetCreatedStartAt(v string) *CarApplyQueryRequest {
	s.CreatedStartAt = &v
	return s
}

func (s *CarApplyQueryRequest) SetPageNumber(v int32) *CarApplyQueryRequest {
	s.PageNumber = &v
	return s
}

func (s *CarApplyQueryRequest) SetPageSize(v int32) *CarApplyQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarApplyQueryRequest) SetThirdPartApplyId(v string) *CarApplyQueryRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyQueryRequest) SetUserId(v string) *CarApplyQueryRequest {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
