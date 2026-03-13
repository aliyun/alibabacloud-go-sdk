// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListWorkspacesRequest
	GetName() *string
	SetPageNumber(v string) *ListWorkspacesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListWorkspacesRequest
	GetPageSize() *string
	SetResourceGroupId(v string) *ListWorkspacesRequest
	GetResourceGroupId() *string
}

type ListWorkspacesRequest struct {
	// example:
	//
	// intl_synonym_module
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// rg-aek25sodlatnioq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListWorkspacesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListWorkspacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v string) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v string) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetResourceGroupId(v string) *ListWorkspacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
