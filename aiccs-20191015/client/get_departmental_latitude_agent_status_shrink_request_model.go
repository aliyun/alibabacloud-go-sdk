// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepartmentalLatitudeAgentStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetCurrentPage() *int64
	SetDepIdsShrink(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetPageSize() *int64
	SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest
	GetStartDate() *int64
}

type GetDepartmentalLatitudeAgentStatusShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage  *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetDepIdsShrink(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
