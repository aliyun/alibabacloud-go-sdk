// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectManagerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgEntitiesShrink(v string) *RemoveProjectManagerShrinkRequest
	GetOrgEntitiesShrink() *string
	SetOutProjectId(v string) *RemoveProjectManagerShrinkRequest
	GetOutProjectId() *string
	SetProjectId(v int64) *RemoveProjectManagerShrinkRequest
	GetProjectId() *int64
	SetRemoveAll(v bool) *RemoveProjectManagerShrinkRequest
	GetRemoveAll() *bool
}

type RemoveProjectManagerShrinkRequest struct {
	OrgEntitiesShrink *string `json:"org_entities,omitempty" xml:"org_entities,omitempty"`
	// example:
	//
	// projectabc
	OutProjectId *string `json:"out_project_id,omitempty" xml:"out_project_id,omitempty"`
	// example:
	//
	// 123
	ProjectId *int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// example:
	//
	// false
	RemoveAll *bool `json:"remove_all,omitempty" xml:"remove_all,omitempty"`
}

func (s RemoveProjectManagerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerShrinkRequest) GetOrgEntitiesShrink() *string {
	return s.OrgEntitiesShrink
}

func (s *RemoveProjectManagerShrinkRequest) GetOutProjectId() *string {
	return s.OutProjectId
}

func (s *RemoveProjectManagerShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RemoveProjectManagerShrinkRequest) GetRemoveAll() *bool {
	return s.RemoveAll
}

func (s *RemoveProjectManagerShrinkRequest) SetOrgEntitiesShrink(v string) *RemoveProjectManagerShrinkRequest {
	s.OrgEntitiesShrink = &v
	return s
}

func (s *RemoveProjectManagerShrinkRequest) SetOutProjectId(v string) *RemoveProjectManagerShrinkRequest {
	s.OutProjectId = &v
	return s
}

func (s *RemoveProjectManagerShrinkRequest) SetProjectId(v int64) *RemoveProjectManagerShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *RemoveProjectManagerShrinkRequest) SetRemoveAll(v bool) *RemoveProjectManagerShrinkRequest {
	s.RemoveAll = &v
	return s
}

func (s *RemoveProjectManagerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
