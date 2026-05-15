// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *DeleteDataAgentKnowledgeBaseRequest
	GetDMSUnit() *string
	SetKbUuid(v string) *DeleteDataAgentKnowledgeBaseRequest
	GetKbUuid() *string
	SetWorkspaceId(v string) *DeleteDataAgentKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentKnowledgeBaseRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-HZ-rtl5lwx********q32d3ux
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2dl4opo5vbh*******frxfsmw
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentKnowledgeBaseRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DeleteDataAgentKnowledgeBaseRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DeleteDataAgentKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentKnowledgeBaseRequest) SetDMSUnit(v string) *DeleteDataAgentKnowledgeBaseRequest {
	s.DMSUnit = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseRequest) SetKbUuid(v string) *DeleteDataAgentKnowledgeBaseRequest {
	s.KbUuid = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseRequest) SetWorkspaceId(v string) *DeleteDataAgentKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
