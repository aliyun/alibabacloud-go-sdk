// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateConfigShrink(v string) *UpdateJobTemplateShrinkRequest
	GetJobTemplateConfigShrink() *string
	SetJobTemplateDescription(v string) *UpdateJobTemplateShrinkRequest
	GetJobTemplateDescription() *string
	SetJobTemplateId(v string) *UpdateJobTemplateShrinkRequest
	GetJobTemplateId() *string
	SetJobTemplateName(v string) *UpdateJobTemplateShrinkRequest
	GetJobTemplateName() *string
}

type UpdateJobTemplateShrinkRequest struct {
	// The template configuration.
	JobTemplateConfigShrink *string `json:"JobTemplateConfig,omitempty" xml:"JobTemplateConfig,omitempty"`
	// The template description.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The job template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	JobTemplateName *string `json:"JobTemplateName,omitempty" xml:"JobTemplateName,omitempty"`
}

func (s UpdateJobTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateShrinkRequest) GetJobTemplateConfigShrink() *string {
	return s.JobTemplateConfigShrink
}

func (s *UpdateJobTemplateShrinkRequest) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *UpdateJobTemplateShrinkRequest) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *UpdateJobTemplateShrinkRequest) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *UpdateJobTemplateShrinkRequest) SetJobTemplateConfigShrink(v string) *UpdateJobTemplateShrinkRequest {
	s.JobTemplateConfigShrink = &v
	return s
}

func (s *UpdateJobTemplateShrinkRequest) SetJobTemplateDescription(v string) *UpdateJobTemplateShrinkRequest {
	s.JobTemplateDescription = &v
	return s
}

func (s *UpdateJobTemplateShrinkRequest) SetJobTemplateId(v string) *UpdateJobTemplateShrinkRequest {
	s.JobTemplateId = &v
	return s
}

func (s *UpdateJobTemplateShrinkRequest) SetJobTemplateName(v string) *UpdateJobTemplateShrinkRequest {
	s.JobTemplateName = &v
	return s
}

func (s *UpdateJobTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
