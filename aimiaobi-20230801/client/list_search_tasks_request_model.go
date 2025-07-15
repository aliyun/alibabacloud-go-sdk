// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueTypes(v []*int32) *ListSearchTasksRequest
	GetDialogueTypes() []*int32
	SetPageNumber(v int32) *ListSearchTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTasksRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListSearchTasksRequest
	GetWorkspaceId() *string
}

type ListSearchTasksRequest struct {
	// example:
	//
	// 24
	DialogueTypes []*int32 `json:"DialogueTypes,omitempty" xml:"DialogueTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSearchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSearchTasksRequest) GetDialogueTypes() []*int32 {
	return s.DialogueTypes
}

func (s *ListSearchTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSearchTasksRequest) SetDialogueTypes(v []*int32) *ListSearchTasksRequest {
	s.DialogueTypes = v
	return s
}

func (s *ListSearchTasksRequest) SetPageNumber(v int32) *ListSearchTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTasksRequest) SetPageSize(v int32) *ListSearchTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchTasksRequest) SetWorkspaceId(v string) *ListSearchTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSearchTasksRequest) Validate() error {
	return dara.Validate(s)
}
