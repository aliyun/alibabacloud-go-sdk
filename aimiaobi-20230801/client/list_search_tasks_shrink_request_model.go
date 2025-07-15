// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueTypesShrink(v string) *ListSearchTasksShrinkRequest
	GetDialogueTypesShrink() *string
	SetPageNumber(v int32) *ListSearchTasksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTasksShrinkRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListSearchTasksShrinkRequest
	GetWorkspaceId() *string
}

type ListSearchTasksShrinkRequest struct {
	// example:
	//
	// 24
	DialogueTypesShrink *string `json:"DialogueTypes,omitempty" xml:"DialogueTypes,omitempty"`
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

func (s ListSearchTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSearchTasksShrinkRequest) GetDialogueTypesShrink() *string {
	return s.DialogueTypesShrink
}

func (s *ListSearchTasksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTasksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTasksShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSearchTasksShrinkRequest) SetDialogueTypesShrink(v string) *ListSearchTasksShrinkRequest {
	s.DialogueTypesShrink = &v
	return s
}

func (s *ListSearchTasksShrinkRequest) SetPageNumber(v int32) *ListSearchTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTasksShrinkRequest) SetPageSize(v int32) *ListSearchTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchTasksShrinkRequest) SetWorkspaceId(v string) *ListSearchTasksShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSearchTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
