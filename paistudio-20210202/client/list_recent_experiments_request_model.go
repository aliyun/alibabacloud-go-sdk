// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentExperimentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListRecentExperimentsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListRecentExperimentsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRecentExperimentsRequest
	GetPageSize() *int64
	SetSource(v string) *ListRecentExperimentsRequest
	GetSource() *string
	SetType(v string) *ListRecentExperimentsRequest
	GetType() *string
	SetWorkspaceId(v string) *ListRecentExperimentsRequest
	GetWorkspaceId() *string
}

type ListRecentExperimentsRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Modified
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRecentExperimentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecentExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRecentExperimentsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRecentExperimentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRecentExperimentsRequest) GetSource() *string {
	return s.Source
}

func (s *ListRecentExperimentsRequest) GetType() *string {
	return s.Type
}

func (s *ListRecentExperimentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRecentExperimentsRequest) SetOrder(v string) *ListRecentExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetPageNumber(v int64) *ListRecentExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetPageSize(v int64) *ListRecentExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetSource(v string) *ListRecentExperimentsRequest {
	s.Source = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetType(v string) *ListRecentExperimentsRequest {
	s.Type = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetWorkspaceId(v string) *ListRecentExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListRecentExperimentsRequest) Validate() error {
	return dara.Validate(s)
}
