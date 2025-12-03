// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestWorkflowExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderBy(v string) *ListChangeRequestWorkflowExecutionsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListChangeRequestWorkflowExecutionsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListChangeRequestWorkflowExecutionsRequest
	GetPage() *int64
	SetPerPage(v int64) *ListChangeRequestWorkflowExecutionsRequest
	GetPerPage() *int64
	SetReleaseStageSn(v string) *ListChangeRequestWorkflowExecutionsRequest
	GetReleaseStageSn() *string
	SetReleaseWorkflowSn(v string) *ListChangeRequestWorkflowExecutionsRequest
	GetReleaseWorkflowSn() *string
	SetSort(v string) *ListChangeRequestWorkflowExecutionsRequest
	GetSort() *string
}

type ListChangeRequestWorkflowExecutionsRequest struct {
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PerPage *int64 `json:"perPage,omitempty" xml:"perPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e173f3c11db5445eb426ca33c92207c8
	ReleaseStageSn *string `json:"releaseStageSn,omitempty" xml:"releaseStageSn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ce51b31b996246ecaf874736838360b2
	ReleaseWorkflowSn *string `json:"releaseWorkflowSn,omitempty" xml:"releaseWorkflowSn,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListChangeRequestWorkflowExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestWorkflowExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetPerPage() *int64 {
	return s.PerPage
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetReleaseStageSn() *string {
	return s.ReleaseStageSn
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetReleaseWorkflowSn() *string {
	return s.ReleaseWorkflowSn
}

func (s *ListChangeRequestWorkflowExecutionsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetOrderBy(v string) *ListChangeRequestWorkflowExecutionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetOrganizationId(v string) *ListChangeRequestWorkflowExecutionsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetPage(v int64) *ListChangeRequestWorkflowExecutionsRequest {
	s.Page = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetPerPage(v int64) *ListChangeRequestWorkflowExecutionsRequest {
	s.PerPage = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetReleaseStageSn(v string) *ListChangeRequestWorkflowExecutionsRequest {
	s.ReleaseStageSn = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetReleaseWorkflowSn(v string) *ListChangeRequestWorkflowExecutionsRequest {
	s.ReleaseWorkflowSn = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) SetSort(v string) *ListChangeRequestWorkflowExecutionsRequest {
	s.Sort = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
