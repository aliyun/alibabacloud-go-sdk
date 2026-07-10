// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseUserResponseBodyData) *DescribeLangfuseUserResponseBody
	GetData() *DescribeLangfuseUserResponseBodyData
	SetRequestId(v string) *DescribeLangfuseUserResponseBody
	GetRequestId() *string
}

type DescribeLangfuseUserResponseBody struct {
	// The returned result.
	Data *DescribeLangfuseUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserResponseBody) GetData() *DescribeLangfuseUserResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseUserResponseBody) SetData(v *DescribeLangfuseUserResponseBodyData) *DescribeLangfuseUserResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseUserResponseBody) SetRequestId(v string) *DescribeLangfuseUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseUserResponseBodyData struct {
	// The time when the user was created.
	//
	// example:
	//
	// 2026-06-01T10:03:05Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The role information of the user.
	Memberships []*DescribeLangfuseUserResponseBodyDataMemberships `json:"Memberships,omitempty" xml:"Memberships,omitempty" type:"Repeated"`
	// The username.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the user was last updated.
	//
	// example:
	//
	// 2026-06-01T10:03:05Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeLangfuseUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLangfuseUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseUserResponseBodyData) GetMemberships() []*DescribeLangfuseUserResponseBodyDataMemberships {
	return s.Memberships
}

func (s *DescribeLangfuseUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseUserResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeLangfuseUserResponseBodyData) SetCreatedAt(v string) *DescribeLangfuseUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyData) SetEmail(v string) *DescribeLangfuseUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyData) SetMemberships(v []*DescribeLangfuseUserResponseBodyDataMemberships) *DescribeLangfuseUserResponseBodyData {
	s.Memberships = v
	return s
}

func (s *DescribeLangfuseUserResponseBodyData) SetName(v string) *DescribeLangfuseUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyData) SetUpdatedAt(v string) *DescribeLangfuseUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyData) Validate() error {
	if s.Memberships != nil {
		for _, item := range s.Memberships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseUserResponseBodyDataMemberships struct {
	// The role of the user within the organization.
	//
	// example:
	//
	// ADMIN
	OrgRole *string `json:"OrgRole,omitempty" xml:"OrgRole,omitempty"`
	// The Langfuse organization ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The Langfuse organization name.
	//
	// example:
	//
	// org_name
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// The list of Langfuse projects.
	Projects []*DescribeLangfuseUserResponseBodyDataMembershipsProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s DescribeLangfuseUserResponseBodyDataMemberships) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserResponseBodyDataMemberships) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) GetOrgRole() *string {
	return s.OrgRole
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) GetOrganizationName() *string {
	return s.OrganizationName
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) GetProjects() []*DescribeLangfuseUserResponseBodyDataMembershipsProjects {
	return s.Projects
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) SetOrgRole(v string) *DescribeLangfuseUserResponseBodyDataMemberships {
	s.OrgRole = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) SetOrganizationId(v string) *DescribeLangfuseUserResponseBodyDataMemberships {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) SetOrganizationName(v string) *DescribeLangfuseUserResponseBodyDataMemberships {
	s.OrganizationName = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) SetProjects(v []*DescribeLangfuseUserResponseBodyDataMembershipsProjects) *DescribeLangfuseUserResponseBodyDataMemberships {
	s.Projects = v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMemberships) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseUserResponseBodyDataMembershipsProjects struct {
	// The Langfuse project ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The Langfuse project name.
	//
	// example:
	//
	// project_name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The role of the user within the project.
	//
	// example:
	//
	// VIEWER
	ProjectRole *string `json:"ProjectRole,omitempty" xml:"ProjectRole,omitempty"`
}

func (s DescribeLangfuseUserResponseBodyDataMembershipsProjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserResponseBodyDataMembershipsProjects) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) GetProjectRole() *string {
	return s.ProjectRole
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) SetProjectId(v string) *DescribeLangfuseUserResponseBodyDataMembershipsProjects {
	s.ProjectId = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) SetProjectName(v string) *DescribeLangfuseUserResponseBodyDataMembershipsProjects {
	s.ProjectName = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) SetProjectRole(v string) *DescribeLangfuseUserResponseBodyDataMembershipsProjects {
	s.ProjectRole = &v
	return s
}

func (s *DescribeLangfuseUserResponseBodyDataMembershipsProjects) Validate() error {
	return dara.Validate(s)
}
