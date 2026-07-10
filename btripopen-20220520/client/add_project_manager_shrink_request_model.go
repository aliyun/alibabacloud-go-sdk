// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectManagerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgEntitiesShrink(v string) *AddProjectManagerShrinkRequest
	GetOrgEntitiesShrink() *string
	SetOutProjectId(v string) *AddProjectManagerShrinkRequest
	GetOutProjectId() *string
	SetProjectId(v int64) *AddProjectManagerShrinkRequest
	GetProjectId() *int64
}

type AddProjectManagerShrinkRequest struct {
	OrgEntitiesShrink *string `json:"org_entities,omitempty" xml:"org_entities,omitempty"`
	OutProjectId      *string `json:"out_project_id,omitempty" xml:"out_project_id,omitempty"`
	ProjectId         *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

func (s AddProjectManagerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddProjectManagerShrinkRequest) GetOrgEntitiesShrink() *string {
	return s.OrgEntitiesShrink
}

func (s *AddProjectManagerShrinkRequest) GetOutProjectId() *string {
	return s.OutProjectId
}

func (s *AddProjectManagerShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AddProjectManagerShrinkRequest) SetOrgEntitiesShrink(v string) *AddProjectManagerShrinkRequest {
	s.OrgEntitiesShrink = &v
	return s
}

func (s *AddProjectManagerShrinkRequest) SetOutProjectId(v string) *AddProjectManagerShrinkRequest {
	s.OutProjectId = &v
	return s
}

func (s *AddProjectManagerShrinkRequest) SetProjectId(v int64) *AddProjectManagerShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *AddProjectManagerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
