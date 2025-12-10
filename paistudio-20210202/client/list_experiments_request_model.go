// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListExperimentsRequest
	GetCreator() *string
	SetExperimentId(v string) *ListExperimentsRequest
	GetExperimentId() *string
	SetName(v string) *ListExperimentsRequest
	GetName() *string
	SetOrder(v string) *ListExperimentsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListExperimentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExperimentsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListExperimentsRequest
	GetSortBy() *string
	SetSource(v string) *ListExperimentsRequest
	GetSource() *string
	SetWorkspaceId(v string) *ListExperimentsRequest
	GetWorkspaceId() *string
}

type ListExperimentsRequest struct {
	// example:
	//
	// 13266******376250
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// draft-rbvg5wzljzjhc9ks92
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreate
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34875
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListExperimentsRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListExperimentsRequest) GetName() *string {
	return s.Name
}

func (s *ListExperimentsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListExperimentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExperimentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListExperimentsRequest) GetSource() *string {
	return s.Source
}

func (s *ListExperimentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListExperimentsRequest) SetCreator(v string) *ListExperimentsRequest {
	s.Creator = &v
	return s
}

func (s *ListExperimentsRequest) SetExperimentId(v string) *ListExperimentsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsRequest) SetName(v string) *ListExperimentsRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentsRequest) SetOrder(v string) *ListExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentsRequest) SetPageNumber(v int32) *ListExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentsRequest) SetPageSize(v int32) *ListExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentsRequest) SetSortBy(v string) *ListExperimentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentsRequest) SetSource(v string) *ListExperimentsRequest {
	s.Source = &v
	return s
}

func (s *ListExperimentsRequest) SetWorkspaceId(v string) *ListExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListExperimentsRequest) Validate() error {
	return dara.Validate(s)
}
