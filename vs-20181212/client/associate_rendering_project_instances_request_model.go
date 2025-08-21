// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRenderingProjectInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *AssociateRenderingProjectInstancesRequest
	GetProjectId() *string
	SetRenderingInstanceIds(v []*string) *AssociateRenderingProjectInstancesRequest
	GetRenderingInstanceIds() []*string
}

type AssociateRenderingProjectInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s AssociateRenderingProjectInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesRequest) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AssociateRenderingProjectInstancesRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *AssociateRenderingProjectInstancesRequest) SetProjectId(v string) *AssociateRenderingProjectInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateRenderingProjectInstancesRequest) SetRenderingInstanceIds(v []*string) *AssociateRenderingProjectInstancesRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *AssociateRenderingProjectInstancesRequest) Validate() error {
	return dara.Validate(s)
}
