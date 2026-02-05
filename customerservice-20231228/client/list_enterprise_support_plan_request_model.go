// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListEnterpriseSupportPlanRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListEnterpriseSupportPlanRequest
	GetPageSize() *int32
}

type ListEnterpriseSupportPlanRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnterpriseSupportPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListEnterpriseSupportPlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseSupportPlanRequest) SetPageNum(v int32) *ListEnterpriseSupportPlanRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanRequest) SetPageSize(v int32) *ListEnterpriseSupportPlanRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseSupportPlanRequest) Validate() error {
	return dara.Validate(s)
}
