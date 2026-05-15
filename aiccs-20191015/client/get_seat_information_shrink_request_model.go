// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeatInformationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSeatInformationShrinkRequest
	GetInstanceId() *string
	SetCurrentPage(v int32) *GetSeatInformationShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSeatInformationShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSeatInformationShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSeatInformationShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetPageSize(v int32) *GetSeatInformationShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSeatInformationShrinkRequest
	GetStartDate() *int64
}

type GetSeatInformationShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	CurrentPage  *int32  `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	DepIdsShrink *string `json:"depIds,omitempty" xml:"depIds,omitempty"`
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"existDepartmentGrouping,omitempty" xml:"existDepartmentGrouping,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s GetSeatInformationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSeatInformationShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSeatInformationShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSeatInformationShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSeatInformationShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSeatInformationShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSeatInformationShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSeatInformationShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSeatInformationShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSeatInformationShrinkRequest) SetInstanceId(v string) *GetSeatInformationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetCurrentPage(v int32) *GetSeatInformationShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetDepIdsShrink(v string) *GetSeatInformationShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetEndDate(v int64) *GetSeatInformationShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSeatInformationShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetPageSize(v int32) *GetSeatInformationShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) SetStartDate(v int64) *GetSeatInformationShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSeatInformationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
