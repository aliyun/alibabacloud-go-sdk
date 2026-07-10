// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifiedTimeGreaterOrEqualThan(v string) *BatchQueryDepartmentRequest
	GetModifiedTimeGreaterOrEqualThan() *string
	SetOutDeptId(v string) *BatchQueryDepartmentRequest
	GetOutDeptId() *string
	SetPageSize(v int32) *BatchQueryDepartmentRequest
	GetPageSize() *int32
	SetPageToken(v string) *BatchQueryDepartmentRequest
	GetPageToken() *string
}

type BatchQueryDepartmentRequest struct {
	ModifiedTimeGreaterOrEqualThan *string `json:"modified_time_greater_or_equal_than,omitempty" xml:"modified_time_greater_or_equal_than,omitempty"`
	OutDeptId                      *string `json:"out_dept_id,omitempty" xml:"out_dept_id,omitempty"`
	PageSize                       *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PageToken                      *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
}

func (s BatchQueryDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentRequest) GetModifiedTimeGreaterOrEqualThan() *string {
	return s.ModifiedTimeGreaterOrEqualThan
}

func (s *BatchQueryDepartmentRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *BatchQueryDepartmentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BatchQueryDepartmentRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *BatchQueryDepartmentRequest) SetModifiedTimeGreaterOrEqualThan(v string) *BatchQueryDepartmentRequest {
	s.ModifiedTimeGreaterOrEqualThan = &v
	return s
}

func (s *BatchQueryDepartmentRequest) SetOutDeptId(v string) *BatchQueryDepartmentRequest {
	s.OutDeptId = &v
	return s
}

func (s *BatchQueryDepartmentRequest) SetPageSize(v int32) *BatchQueryDepartmentRequest {
	s.PageSize = &v
	return s
}

func (s *BatchQueryDepartmentRequest) SetPageToken(v string) *BatchQueryDepartmentRequest {
	s.PageToken = &v
	return s
}

func (s *BatchQueryDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
