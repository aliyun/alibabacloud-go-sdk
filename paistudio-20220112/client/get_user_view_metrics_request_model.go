// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserViewMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *GetUserViewMetricsRequest
	GetOrder() *string
	SetPageNumber(v string) *GetUserViewMetricsRequest
	GetPageNumber() *string
	SetPageSize(v string) *GetUserViewMetricsRequest
	GetPageSize() *string
	SetSortBy(v string) *GetUserViewMetricsRequest
	GetSortBy() *string
	SetTimeStep(v string) *GetUserViewMetricsRequest
	GetTimeStep() *string
	SetUserId(v string) *GetUserViewMetricsRequest
	GetUserId() *string
	SetWorkspaceId(v string) *GetUserViewMetricsRequest
	GetWorkspaceId() *string
}

type GetUserViewMetricsRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 1h
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserViewMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserViewMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsRequest) GetOrder() *string {
	return s.Order
}

func (s *GetUserViewMetricsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *GetUserViewMetricsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetUserViewMetricsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetUserViewMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *GetUserViewMetricsRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserViewMetricsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUserViewMetricsRequest) SetOrder(v string) *GetUserViewMetricsRequest {
	s.Order = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetPageNumber(v string) *GetUserViewMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetPageSize(v string) *GetUserViewMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetSortBy(v string) *GetUserViewMetricsRequest {
	s.SortBy = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetTimeStep(v string) *GetUserViewMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetUserId(v string) *GetUserViewMetricsRequest {
	s.UserId = &v
	return s
}

func (s *GetUserViewMetricsRequest) SetWorkspaceId(v string) *GetUserViewMetricsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetUserViewMetricsRequest) Validate() error {
	return dara.Validate(s)
}
