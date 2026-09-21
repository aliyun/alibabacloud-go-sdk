// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateCacheInput interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *CreateTemplateCacheInput
	GetTeamID() *string
	SetTemplateID(v string) *CreateTemplateCacheInput
	GetTemplateID() *string
}

type CreateTemplateCacheInput struct {
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// example:
	//
	// us7dxqaezw5uu7aa2cm5
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s CreateTemplateCacheInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateCacheInput) GoString() string {
	return s.String()
}

func (s *CreateTemplateCacheInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateTemplateCacheInput) GetTemplateID() *string {
	return s.TemplateID
}

func (s *CreateTemplateCacheInput) SetTeamID(v string) *CreateTemplateCacheInput {
	s.TeamID = &v
	return s
}

func (s *CreateTemplateCacheInput) SetTemplateID(v string) *CreateTemplateCacheInput {
	s.TemplateID = &v
	return s
}

func (s *CreateTemplateCacheInput) Validate() error {
	return dara.Validate(s)
}
