// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditableNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListEditableNamespaceRequest
	GetNamespace() *string
	SetPageIndex(v string) *ListEditableNamespaceRequest
	GetPageIndex() *string
	SetPageSize(v string) *ListEditableNamespaceRequest
	GetPageSize() *string
	SetRegionId(v string) *ListEditableNamespaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListEditableNamespaceRequest
	GetWorkspaceId() *string
}

type ListEditableNamespaceRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListEditableNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEditableNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEditableNamespaceRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *ListEditableNamespaceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListEditableNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEditableNamespaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListEditableNamespaceRequest) SetNamespace(v string) *ListEditableNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetPageIndex(v string) *ListEditableNamespaceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetPageSize(v string) *ListEditableNamespaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetRegionId(v string) *ListEditableNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *ListEditableNamespaceRequest) SetWorkspaceId(v string) *ListEditableNamespaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListEditableNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
