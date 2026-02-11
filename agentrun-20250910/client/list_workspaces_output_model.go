// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWorkspacesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkspacesOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListWorkspacesOutput
	GetTotal() *int64
	SetWorkspaces(v []*Workspace) *ListWorkspacesOutput
	GetWorkspaces() []*Workspace
}

type ListWorkspacesOutput struct {
	PageNumber *int32       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64       `json:"total,omitempty" xml:"total,omitempty"`
	Workspaces []*Workspace `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesOutput) GoString() string {
	return s.String()
}

func (s *ListWorkspacesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkspacesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspacesOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkspacesOutput) GetWorkspaces() []*Workspace {
	return s.Workspaces
}

func (s *ListWorkspacesOutput) SetPageNumber(v int32) *ListWorkspacesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesOutput) SetPageSize(v int32) *ListWorkspacesOutput {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesOutput) SetTotal(v int64) *ListWorkspacesOutput {
	s.Total = &v
	return s
}

func (s *ListWorkspacesOutput) SetWorkspaces(v []*Workspace) *ListWorkspacesOutput {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesOutput) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
