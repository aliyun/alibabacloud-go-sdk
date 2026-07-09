// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCmsWorkspace(v bool) *DeleteAgentSpaceRequest
	GetDeleteCmsWorkspace() *bool
	SetDeleteMseNamespace(v bool) *DeleteAgentSpaceRequest
	GetDeleteMseNamespace() *bool
	SetDeleteSlsProject(v bool) *DeleteAgentSpaceRequest
	GetDeleteSlsProject() *bool
}

type DeleteAgentSpaceRequest struct {
	// Specifies whether to delete the associated Hybrid Cloud Monitoring workspace.
	//
	// example:
	//
	// false
	DeleteCmsWorkspace *bool `json:"deleteCmsWorkspace,omitempty" xml:"deleteCmsWorkspace,omitempty"`
	// Specifies whether to delete the associated MSE namespace.
	//
	// example:
	//
	// false
	DeleteMseNamespace *bool `json:"deleteMseNamespace,omitempty" xml:"deleteMseNamespace,omitempty"`
	// Specifies whether to delete the associated SLS project.
	//
	// example:
	//
	// false
	DeleteSlsProject *bool `json:"deleteSlsProject,omitempty" xml:"deleteSlsProject,omitempty"`
}

func (s DeleteAgentSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpaceRequest) GetDeleteCmsWorkspace() *bool {
	return s.DeleteCmsWorkspace
}

func (s *DeleteAgentSpaceRequest) GetDeleteMseNamespace() *bool {
	return s.DeleteMseNamespace
}

func (s *DeleteAgentSpaceRequest) GetDeleteSlsProject() *bool {
	return s.DeleteSlsProject
}

func (s *DeleteAgentSpaceRequest) SetDeleteCmsWorkspace(v bool) *DeleteAgentSpaceRequest {
	s.DeleteCmsWorkspace = &v
	return s
}

func (s *DeleteAgentSpaceRequest) SetDeleteMseNamespace(v bool) *DeleteAgentSpaceRequest {
	s.DeleteMseNamespace = &v
	return s
}

func (s *DeleteAgentSpaceRequest) SetDeleteSlsProject(v bool) *DeleteAgentSpaceRequest {
	s.DeleteSlsProject = &v
	return s
}

func (s *DeleteAgentSpaceRequest) Validate() error {
	return dara.Validate(s)
}
