// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateRenderingProjectInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DisassociateRenderingProjectInstancesShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceIdsShrink(v string) *DisassociateRenderingProjectInstancesShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type DisassociateRenderingProjectInstancesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s DisassociateRenderingProjectInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DisassociateRenderingProjectInstancesShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *DisassociateRenderingProjectInstancesShrinkRequest) SetProjectId(v string) *DisassociateRenderingProjectInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesShrinkRequest) SetRenderingInstanceIdsShrink(v string) *DisassociateRenderingProjectInstancesShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
