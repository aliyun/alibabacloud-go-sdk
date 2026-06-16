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
	DeleteCmsWorkspace *bool `json:"deleteCmsWorkspace,omitempty" xml:"deleteCmsWorkspace,omitempty"`
	DeleteMseNamespace *bool `json:"deleteMseNamespace,omitempty" xml:"deleteMseNamespace,omitempty"`
	DeleteSlsProject   *bool `json:"deleteSlsProject,omitempty" xml:"deleteSlsProject,omitempty"`
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
