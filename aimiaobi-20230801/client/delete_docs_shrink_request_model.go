// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIdsShrink(v string) *DeleteDocsShrinkRequest
	GetDocIdsShrink() *string
	SetWorkspaceId(v string) *DeleteDocsShrinkRequest
	GetWorkspaceId() *string
}

type DeleteDocsShrinkRequest struct {
	// This parameter is required.
	DocIdsShrink *string `json:"DocIds,omitempty" xml:"DocIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDocsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocsShrinkRequest) GetDocIdsShrink() *string {
	return s.DocIdsShrink
}

func (s *DeleteDocsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDocsShrinkRequest) SetDocIdsShrink(v string) *DeleteDocsShrinkRequest {
	s.DocIdsShrink = &v
	return s
}

func (s *DeleteDocsShrinkRequest) SetWorkspaceId(v string) *DeleteDocsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDocsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
