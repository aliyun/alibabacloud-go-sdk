// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateExperimentRequest
	GetAccessibility() *string
	SetDescription(v string) *CreateExperimentRequest
	GetDescription() *string
	SetFolderId(v string) *CreateExperimentRequest
	GetFolderId() *string
	SetName(v string) *CreateExperimentRequest
	GetName() *string
	SetOptions(v string) *CreateExperimentRequest
	GetOptions() *string
	SetSource(v string) *CreateExperimentRequest
	GetSource() *string
	SetTemplateId(v string) *CreateExperimentRequest
	GetTemplateId() *string
	SetWorkspaceId(v string) *CreateExperimentRequest
	GetWorkspaceId() *string
}

type CreateExperimentRequest struct {
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
	// folder-xfdafx093xxfd
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"mlflow":{"experimentId":"exp-1"}}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// template-xze5df2scrxxz
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 84972
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateExperimentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExperimentRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateExperimentRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperimentRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateExperimentRequest) GetSource() *string {
	return s.Source
}

func (s *CreateExperimentRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateExperimentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateExperimentRequest) SetAccessibility(v string) *CreateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentRequest) SetDescription(v string) *CreateExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentRequest) SetFolderId(v string) *CreateExperimentRequest {
	s.FolderId = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetOptions(v string) *CreateExperimentRequest {
	s.Options = &v
	return s
}

func (s *CreateExperimentRequest) SetSource(v string) *CreateExperimentRequest {
	s.Source = &v
	return s
}

func (s *CreateExperimentRequest) SetTemplateId(v string) *CreateExperimentRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateExperimentRequest) SetWorkspaceId(v string) *CreateExperimentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateExperimentRequest) Validate() error {
	return dara.Validate(s)
}
