// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectManagerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgEntities(v []*AddProjectManagerRequestOrgEntities) *AddProjectManagerRequest
	GetOrgEntities() []*AddProjectManagerRequestOrgEntities
	SetOutProjectId(v string) *AddProjectManagerRequest
	GetOutProjectId() *string
	SetProjectId(v int64) *AddProjectManagerRequest
	GetProjectId() *int64
}

type AddProjectManagerRequest struct {
	OrgEntities  []*AddProjectManagerRequestOrgEntities `json:"org_entities,omitempty" xml:"org_entities,omitempty" type:"Repeated"`
	OutProjectId *string                                `json:"out_project_id,omitempty" xml:"out_project_id,omitempty"`
	ProjectId    *int64                                 `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

func (s AddProjectManagerRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerRequest) GoString() string {
	return s.String()
}

func (s *AddProjectManagerRequest) GetOrgEntities() []*AddProjectManagerRequestOrgEntities {
	return s.OrgEntities
}

func (s *AddProjectManagerRequest) GetOutProjectId() *string {
	return s.OutProjectId
}

func (s *AddProjectManagerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AddProjectManagerRequest) SetOrgEntities(v []*AddProjectManagerRequestOrgEntities) *AddProjectManagerRequest {
	s.OrgEntities = v
	return s
}

func (s *AddProjectManagerRequest) SetOutProjectId(v string) *AddProjectManagerRequest {
	s.OutProjectId = &v
	return s
}

func (s *AddProjectManagerRequest) SetProjectId(v int64) *AddProjectManagerRequest {
	s.ProjectId = &v
	return s
}

func (s *AddProjectManagerRequest) Validate() error {
	if s.OrgEntities != nil {
		for _, item := range s.OrgEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddProjectManagerRequestOrgEntities struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s AddProjectManagerRequestOrgEntities) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerRequestOrgEntities) GoString() string {
	return s.String()
}

func (s *AddProjectManagerRequestOrgEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *AddProjectManagerRequestOrgEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *AddProjectManagerRequestOrgEntities) SetEntityId(v string) *AddProjectManagerRequestOrgEntities {
	s.EntityId = &v
	return s
}

func (s *AddProjectManagerRequestOrgEntities) SetEntityType(v string) *AddProjectManagerRequestOrgEntities {
	s.EntityType = &v
	return s
}

func (s *AddProjectManagerRequestOrgEntities) Validate() error {
	return dara.Validate(s)
}
