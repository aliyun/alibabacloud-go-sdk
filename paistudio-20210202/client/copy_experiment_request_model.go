// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CopyExperimentRequest
	GetAccessibility() *string
	SetDescription(v string) *CopyExperimentRequest
	GetDescription() *string
	SetFolderId(v string) *CopyExperimentRequest
	GetFolderId() *string
	SetName(v string) *CopyExperimentRequest
	GetName() *string
	SetSource(v string) *CopyExperimentRequest
	GetSource() *string
	SetWorkspaceId(v string) *CopyExperimentRequest
	GetWorkspaceId() *string
}

type CopyExperimentRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// Pipeline draft description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// folder-erwx872xuryx
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 84972
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CopyExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyExperimentRequest) GoString() string {
	return s.String()
}

func (s *CopyExperimentRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CopyExperimentRequest) GetDescription() *string {
	return s.Description
}

func (s *CopyExperimentRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *CopyExperimentRequest) GetName() *string {
	return s.Name
}

func (s *CopyExperimentRequest) GetSource() *string {
	return s.Source
}

func (s *CopyExperimentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CopyExperimentRequest) SetAccessibility(v string) *CopyExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CopyExperimentRequest) SetDescription(v string) *CopyExperimentRequest {
	s.Description = &v
	return s
}

func (s *CopyExperimentRequest) SetFolderId(v string) *CopyExperimentRequest {
	s.FolderId = &v
	return s
}

func (s *CopyExperimentRequest) SetName(v string) *CopyExperimentRequest {
	s.Name = &v
	return s
}

func (s *CopyExperimentRequest) SetSource(v string) *CopyExperimentRequest {
	s.Source = &v
	return s
}

func (s *CopyExperimentRequest) SetWorkspaceId(v string) *CopyExperimentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CopyExperimentRequest) Validate() error {
	return dara.Validate(s)
}
