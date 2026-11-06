// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSkillDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *DescribeDataAgentSkillDownloadRequest
	GetSkillId() *string
	SetWorkspaceId(v string) *DescribeDataAgentSkillDownloadRequest
	GetWorkspaceId() *string
}

type DescribeDataAgentSkillDownloadRequest struct {
	// The skill ID.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 2ktbdovwf6********cznfo
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeDataAgentSkillDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSkillDownloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSkillDownloadRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *DescribeDataAgentSkillDownloadRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDataAgentSkillDownloadRequest) SetSkillId(v string) *DescribeDataAgentSkillDownloadRequest {
	s.SkillId = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadRequest) SetWorkspaceId(v string) *DescribeDataAgentSkillDownloadRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadRequest) Validate() error {
	return dara.Validate(s)
}
