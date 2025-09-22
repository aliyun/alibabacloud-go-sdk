// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermEditShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TermEditShrinkRequest
	GetAction() *string
	SetExtShrink(v string) *TermEditShrinkRequest
	GetExtShrink() *string
	SetScene(v string) *TermEditShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *TermEditShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TermEditShrinkRequest
	GetTargetLanguage() *string
	SetWorkspaceId(v string) *TermEditShrinkRequest
	GetWorkspaceId() *string
}

type TermEditShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ADD
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s TermEditShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TermEditShrinkRequest) GoString() string {
	return s.String()
}

func (s *TermEditShrinkRequest) GetAction() *string {
	return s.Action
}

func (s *TermEditShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *TermEditShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *TermEditShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TermEditShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TermEditShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *TermEditShrinkRequest) SetAction(v string) *TermEditShrinkRequest {
	s.Action = &v
	return s
}

func (s *TermEditShrinkRequest) SetExtShrink(v string) *TermEditShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *TermEditShrinkRequest) SetScene(v string) *TermEditShrinkRequest {
	s.Scene = &v
	return s
}

func (s *TermEditShrinkRequest) SetSourceLanguage(v string) *TermEditShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TermEditShrinkRequest) SetTargetLanguage(v string) *TermEditShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TermEditShrinkRequest) SetWorkspaceId(v string) *TermEditShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *TermEditShrinkRequest) Validate() error {
	return dara.Validate(s)
}
