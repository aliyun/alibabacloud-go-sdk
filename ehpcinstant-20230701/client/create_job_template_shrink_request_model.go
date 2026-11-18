// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateConfigShrink(v string) *CreateJobTemplateShrinkRequest
	GetJobTemplateConfigShrink() *string
	SetJobTemplateDescription(v string) *CreateJobTemplateShrinkRequest
	GetJobTemplateDescription() *string
	SetJobTemplateName(v string) *CreateJobTemplateShrinkRequest
	GetJobTemplateName() *string
}

type CreateJobTemplateShrinkRequest struct {
	// The configuration of the template.
	JobTemplateConfigShrink *string `json:"JobTemplateConfig,omitempty" xml:"JobTemplateConfig,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-template
	JobTemplateName *string `json:"JobTemplateName,omitempty" xml:"JobTemplateName,omitempty"`
}

func (s CreateJobTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateShrinkRequest) GetJobTemplateConfigShrink() *string {
	return s.JobTemplateConfigShrink
}

func (s *CreateJobTemplateShrinkRequest) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *CreateJobTemplateShrinkRequest) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *CreateJobTemplateShrinkRequest) SetJobTemplateConfigShrink(v string) *CreateJobTemplateShrinkRequest {
	s.JobTemplateConfigShrink = &v
	return s
}

func (s *CreateJobTemplateShrinkRequest) SetJobTemplateDescription(v string) *CreateJobTemplateShrinkRequest {
	s.JobTemplateDescription = &v
	return s
}

func (s *CreateJobTemplateShrinkRequest) SetJobTemplateName(v string) *CreateJobTemplateShrinkRequest {
	s.JobTemplateName = &v
	return s
}

func (s *CreateJobTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
