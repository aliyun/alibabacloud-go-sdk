// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PutWorkspaceRequest
	GetDescription() *string
	SetDisplayName(v string) *PutWorkspaceRequest
	GetDisplayName() *string
	SetResourceGroupId(v string) *PutWorkspaceRequest
	GetResourceGroupId() *string
	SetSlsProject(v string) *PutWorkspaceRequest
	GetSlsProject() *string
	SetTags(v []*PutWorkspaceRequestTags) *PutWorkspaceRequest
	GetTags() []*PutWorkspaceRequestTags
}

type PutWorkspaceRequest struct {
	// The description of the workspace.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The ID of the resource group specified when the workspace is created.
	//
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// The tags attached to the workspace when it is created.
	Tags []*PutWorkspaceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s PutWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *PutWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *PutWorkspaceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PutWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *PutWorkspaceRequest) GetSlsProject() *string {
	return s.SlsProject
}

func (s *PutWorkspaceRequest) GetTags() []*PutWorkspaceRequestTags {
	return s.Tags
}

func (s *PutWorkspaceRequest) SetDescription(v string) *PutWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *PutWorkspaceRequest) SetDisplayName(v string) *PutWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *PutWorkspaceRequest) SetResourceGroupId(v string) *PutWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *PutWorkspaceRequest) SetSlsProject(v string) *PutWorkspaceRequest {
	s.SlsProject = &v
	return s
}

func (s *PutWorkspaceRequest) SetTags(v []*PutWorkspaceRequestTags) *PutWorkspaceRequest {
	s.Tags = v
	return s
}

func (s *PutWorkspaceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutWorkspaceRequestTags struct {
	// The key of the tag.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PutWorkspaceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s PutWorkspaceRequestTags) GoString() string {
	return s.String()
}

func (s *PutWorkspaceRequestTags) GetKey() *string {
	return s.Key
}

func (s *PutWorkspaceRequestTags) GetValue() *string {
	return s.Value
}

func (s *PutWorkspaceRequestTags) SetKey(v string) *PutWorkspaceRequestTags {
	s.Key = &v
	return s
}

func (s *PutWorkspaceRequestTags) SetValue(v string) *PutWorkspaceRequestTags {
	s.Value = &v
	return s
}

func (s *PutWorkspaceRequestTags) Validate() error {
	return dara.Validate(s)
}
