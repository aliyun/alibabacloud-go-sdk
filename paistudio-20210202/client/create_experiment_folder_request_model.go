// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateExperimentFolderRequest
	GetAccessibility() *string
	SetName(v string) *CreateExperimentFolderRequest
	GetName() *string
	SetParentFolderId(v string) *CreateExperimentFolderRequest
	GetParentFolderId() *string
	SetSource(v string) *CreateExperimentFolderRequest
	GetSource() *string
	SetWorkspaceId(v string) *CreateExperimentFolderRequest
	GetWorkspaceId() *string
}

type CreateExperimentFolderRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// folder-xxfdjhfxdfad
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 45699
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateExperimentFolderRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperimentFolderRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateExperimentFolderRequest) GetSource() *string {
	return s.Source
}

func (s *CreateExperimentFolderRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateExperimentFolderRequest) SetAccessibility(v string) *CreateExperimentFolderRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetName(v string) *CreateExperimentFolderRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetParentFolderId(v string) *CreateExperimentFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetSource(v string) *CreateExperimentFolderRequest {
	s.Source = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetWorkspaceId(v string) *CreateExperimentFolderRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateExperimentFolderRequest) Validate() error {
	return dara.Validate(s)
}
