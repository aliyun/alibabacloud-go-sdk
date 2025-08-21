// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRenderingProjectInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *AssociateRenderingProjectInstancesShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceIdsShrink(v string) *AssociateRenderingProjectInstancesShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type AssociateRenderingProjectInstancesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s AssociateRenderingProjectInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AssociateRenderingProjectInstancesShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *AssociateRenderingProjectInstancesShrinkRequest) SetProjectId(v string) *AssociateRenderingProjectInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateRenderingProjectInstancesShrinkRequest) SetRenderingInstanceIdsShrink(v string) *AssociateRenderingProjectInstancesShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *AssociateRenderingProjectInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
