// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentSkillMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *DeleteDataAgentSkillMetaRequest
	GetSkillId() *string
	SetWorkspaceId(v string) *DeleteDataAgentSkillMetaRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentSkillMetaRequest struct {
	// The skill ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentSkillMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentSkillMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentSkillMetaRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteDataAgentSkillMetaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentSkillMetaRequest) SetSkillId(v string) *DeleteDataAgentSkillMetaRequest {
	s.SkillId = &v
	return s
}

func (s *DeleteDataAgentSkillMetaRequest) SetWorkspaceId(v string) *DeleteDataAgentSkillMetaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentSkillMetaRequest) Validate() error {
	return dara.Validate(s)
}
