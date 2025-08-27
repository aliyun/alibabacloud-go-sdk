// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifiedTimeGreaterOrEqualThan(v string) *UserQueryRequest
	GetModifiedTimeGreaterOrEqualThan() *string
	SetPageSize(v int32) *UserQueryRequest
	GetPageSize() *int32
	SetPageToken(v string) *UserQueryRequest
	GetPageToken() *string
	SetThirdPartJobNo(v string) *UserQueryRequest
	GetThirdPartJobNo() *string
}

type UserQueryRequest struct {
	// example:
	//
	// 2021-06-02 00:00:00
	ModifiedTimeGreaterOrEqualThan *string `json:"modified_time_greater_or_equal_than,omitempty" xml:"modified_time_greater_or_equal_than,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 9YN+jxa7PcxbNUTISeKjEw==
	PageToken *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
	// example:
	//
	// 12345
	ThirdPartJobNo *string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
}

func (s UserQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s UserQueryRequest) GoString() string {
	return s.String()
}

func (s *UserQueryRequest) GetModifiedTimeGreaterOrEqualThan() *string {
	return s.ModifiedTimeGreaterOrEqualThan
}

func (s *UserQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UserQueryRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *UserQueryRequest) GetThirdPartJobNo() *string {
	return s.ThirdPartJobNo
}

func (s *UserQueryRequest) SetModifiedTimeGreaterOrEqualThan(v string) *UserQueryRequest {
	s.ModifiedTimeGreaterOrEqualThan = &v
	return s
}

func (s *UserQueryRequest) SetPageSize(v int32) *UserQueryRequest {
	s.PageSize = &v
	return s
}

func (s *UserQueryRequest) SetPageToken(v string) *UserQueryRequest {
	s.PageToken = &v
	return s
}

func (s *UserQueryRequest) SetThirdPartJobNo(v string) *UserQueryRequest {
	s.ThirdPartJobNo = &v
	return s
}

func (s *UserQueryRequest) Validate() error {
	return dara.Validate(s)
}
