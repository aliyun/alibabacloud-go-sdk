// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectManagerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgEntities(v []*RemoveProjectManagerRequestOrgEntities) *RemoveProjectManagerRequest
	GetOrgEntities() []*RemoveProjectManagerRequestOrgEntities
	SetOutProjectId(v string) *RemoveProjectManagerRequest
	GetOutProjectId() *string
	SetProjectId(v int64) *RemoveProjectManagerRequest
	GetProjectId() *int64
	SetRemoveAll(v bool) *RemoveProjectManagerRequest
	GetRemoveAll() *bool
}

type RemoveProjectManagerRequest struct {
	OrgEntities  []*RemoveProjectManagerRequestOrgEntities `json:"org_entities,omitempty" xml:"org_entities,omitempty" type:"Repeated"`
	OutProjectId *string                                   `json:"out_project_id,omitempty" xml:"out_project_id,omitempty"`
	ProjectId    *int64                                    `json:"project_id,omitempty" xml:"project_id,omitempty"`
	RemoveAll    *bool                                     `json:"remove_all,omitempty" xml:"remove_all,omitempty"`
}

func (s RemoveProjectManagerRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerRequest) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerRequest) GetOrgEntities() []*RemoveProjectManagerRequestOrgEntities {
	return s.OrgEntities
}

func (s *RemoveProjectManagerRequest) GetOutProjectId() *string {
	return s.OutProjectId
}

func (s *RemoveProjectManagerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RemoveProjectManagerRequest) GetRemoveAll() *bool {
	return s.RemoveAll
}

func (s *RemoveProjectManagerRequest) SetOrgEntities(v []*RemoveProjectManagerRequestOrgEntities) *RemoveProjectManagerRequest {
	s.OrgEntities = v
	return s
}

func (s *RemoveProjectManagerRequest) SetOutProjectId(v string) *RemoveProjectManagerRequest {
	s.OutProjectId = &v
	return s
}

func (s *RemoveProjectManagerRequest) SetProjectId(v int64) *RemoveProjectManagerRequest {
	s.ProjectId = &v
	return s
}

func (s *RemoveProjectManagerRequest) SetRemoveAll(v bool) *RemoveProjectManagerRequest {
	s.RemoveAll = &v
	return s
}

func (s *RemoveProjectManagerRequest) Validate() error {
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

type RemoveProjectManagerRequestOrgEntities struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s RemoveProjectManagerRequestOrgEntities) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerRequestOrgEntities) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerRequestOrgEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *RemoveProjectManagerRequestOrgEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *RemoveProjectManagerRequestOrgEntities) SetEntityId(v string) *RemoveProjectManagerRequestOrgEntities {
	s.EntityId = &v
	return s
}

func (s *RemoveProjectManagerRequestOrgEntities) SetEntityType(v string) *RemoveProjectManagerRequestOrgEntities {
	s.EntityType = &v
	return s
}

func (s *RemoveProjectManagerRequestOrgEntities) Validate() error {
	return dara.Validate(s)
}
