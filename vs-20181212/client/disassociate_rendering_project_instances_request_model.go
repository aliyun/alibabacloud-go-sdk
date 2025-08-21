// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateRenderingProjectInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DisassociateRenderingProjectInstancesRequest
	GetProjectId() *string
	SetRenderingInstanceIds(v []*string) *DisassociateRenderingProjectInstancesRequest
	GetRenderingInstanceIds() []*string
}

type DisassociateRenderingProjectInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s DisassociateRenderingProjectInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesRequest) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DisassociateRenderingProjectInstancesRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *DisassociateRenderingProjectInstancesRequest) SetProjectId(v string) *DisassociateRenderingProjectInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesRequest) SetRenderingInstanceIds(v []*string) *DisassociateRenderingProjectInstancesRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *DisassociateRenderingProjectInstancesRequest) Validate() error {
	return dara.Validate(s)
}
