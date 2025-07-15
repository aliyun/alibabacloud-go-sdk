// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditTermsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdListShrink(v string) *DeleteAuditTermsShrinkRequest
	GetIdListShrink() *string
	SetWorkspaceId(v string) *DeleteAuditTermsShrinkRequest
	GetWorkspaceId() *string
}

type DeleteAuditTermsShrinkRequest struct {
	IdListShrink *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteAuditTermsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditTermsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuditTermsShrinkRequest) GetIdListShrink() *string {
	return s.IdListShrink
}

func (s *DeleteAuditTermsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteAuditTermsShrinkRequest) SetIdListShrink(v string) *DeleteAuditTermsShrinkRequest {
	s.IdListShrink = &v
	return s
}

func (s *DeleteAuditTermsShrinkRequest) SetWorkspaceId(v string) *DeleteAuditTermsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteAuditTermsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
