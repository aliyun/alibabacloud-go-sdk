// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentDraftsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListDeploymentDraftsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentDraftsRequest
	GetPageSize() *int32
}

type ListDeploymentDraftsRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDeploymentDraftsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentDraftsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentDraftsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentDraftsRequest) SetPageIndex(v int32) *ListDeploymentDraftsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentDraftsRequest) SetPageSize(v int32) *ListDeploymentDraftsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentDraftsRequest) Validate() error {
	return dara.Validate(s)
}
