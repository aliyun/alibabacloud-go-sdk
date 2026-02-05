// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanSimpleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListEnterpriseSupportPlanSimpleRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListEnterpriseSupportPlanSimpleRequest
	GetPageSize() *int32
}

type ListEnterpriseSupportPlanSimpleRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListEnterpriseSupportPlanSimpleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseSupportPlanSimpleRequest) SetPageNum(v int32) *ListEnterpriseSupportPlanSimpleRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleRequest) SetPageSize(v int32) *ListEnterpriseSupportPlanSimpleRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleRequest) Validate() error {
	return dara.Validate(s)
}
