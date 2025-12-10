// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateExperimentMetaRequest
	GetAccessibility() *string
	SetDescription(v string) *UpdateExperimentMetaRequest
	GetDescription() *string
	SetFolderId(v string) *UpdateExperimentMetaRequest
	GetFolderId() *string
	SetName(v string) *UpdateExperimentMetaRequest
	GetName() *string
	SetOptions(v string) *UpdateExperimentMetaRequest
	GetOptions() *string
}

type UpdateExperimentMetaRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// Pipeline draft description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// folder-xfd782efd08wex
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// Pipeline draft name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"mlflow":{"experimentId":"exp-1"}}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateExperimentMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateExperimentMetaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExperimentMetaRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateExperimentMetaRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExperimentMetaRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateExperimentMetaRequest) SetAccessibility(v string) *UpdateExperimentMetaRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetDescription(v string) *UpdateExperimentMetaRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetFolderId(v string) *UpdateExperimentMetaRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetName(v string) *UpdateExperimentMetaRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetOptions(v string) *UpdateExperimentMetaRequest {
	s.Options = &v
	return s
}

func (s *UpdateExperimentMetaRequest) Validate() error {
	return dara.Validate(s)
}
