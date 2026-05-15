// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeatInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSeatInformationRequest
	GetInstanceId() *string
	SetCurrentPage(v int32) *GetSeatInformationRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSeatInformationRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSeatInformationRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSeatInformationRequest
	GetExistDepartmentGrouping() *bool
	SetPageSize(v int32) *GetSeatInformationRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSeatInformationRequest
	GetStartDate() *int64
}

type GetSeatInformationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32   `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	DepIds      []*int64 `json:"depIds,omitempty" xml:"depIds,omitempty" type:"Repeated"`
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

func (s GetSeatInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSeatInformationRequest) GoString() string {
	return s.String()
}

func (s *GetSeatInformationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSeatInformationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSeatInformationRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSeatInformationRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSeatInformationRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSeatInformationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSeatInformationRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSeatInformationRequest) SetInstanceId(v string) *GetSeatInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSeatInformationRequest) SetCurrentPage(v int32) *GetSeatInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSeatInformationRequest) SetDepIds(v []*int64) *GetSeatInformationRequest {
	s.DepIds = v
	return s
}

func (s *GetSeatInformationRequest) SetEndDate(v int64) *GetSeatInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetSeatInformationRequest) SetExistDepartmentGrouping(v bool) *GetSeatInformationRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSeatInformationRequest) SetPageSize(v int32) *GetSeatInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetSeatInformationRequest) SetStartDate(v int64) *GetSeatInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetSeatInformationRequest) Validate() error {
	return dara.Validate(s)
}
