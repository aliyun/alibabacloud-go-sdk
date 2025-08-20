// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTemplateRequest
	GetDescription() *string
	SetIsDraft(v bool) *UpdateTemplateRequest
	GetIsDraft() *bool
	SetRotateStrategy(v string) *UpdateTemplateRequest
	GetRotateStrategy() *string
	SetTemplateBody(v string) *UpdateTemplateRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *UpdateTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateTemplateRequest
	GetTemplateName() *string
	SetTemplateURL(v string) *UpdateTemplateRequest
	GetTemplateURL() *string
	SetValidationOptions(v []*string) *UpdateTemplateRequest
	GetValidationOptions() []*string
}

type UpdateTemplateRequest struct {
	// The description of the template. The maximum length is 256 characters.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to update the Draft (draft) version. Values:
	//
	// - false (default): If template content is provided, a new version is created, and the Draft version is cleared. Otherwise, the current latest version is modified.
	//
	// - true: Modifies the Draft version. The Draft version can only be retrieved via the GetTemplate interface. The ListTemplateVersions interface will not return it. The TemplateVersion parameter in other interfaces cannot specify Draft.
	//
	// example:
	//
	// false
	IsDraft *bool `json:"IsDraft,omitempty" xml:"IsDraft,omitempty"`
	// Template version rotation strategy. Values:
	//
	// - None (default): No rotation. An error occurs when the version limit is reached.
	//
	// - DeleteOldestNonSharedVersionWhenLimitExceeded: Rotates and deletes non-shared template versions.
	//
	// >
	//
	// > - If all versions of the template are shared, they cannot be rotated and deleted.
	//
	// > - The current latest version will not be rotated and deleted.
	//
	// > - Regardless of whether rotation deletion is used, the template version number cannot exceed v65000.
	//
	// example:
	//
	// None
	RotateStrategy *string `json:"RotateStrategy,omitempty" xml:"RotateStrategy,omitempty"`
	// The structure of the template body. The length should be between 1 and 524,288 bytes. If the content is long, it is recommended to use HTTP POST + Body Param to pass the parameter in the request body to avoid request failure due to an overly long URL.
	//
	// > You must and can only specify one of `TemplateBody`, `TemplateURL`, `TemplateId`, or `TemplateScratchId`.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion":"2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. Supports both shared and private templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// The length should not exceed 255 characters (utf-8 encoding), and it must start with a number, letter, or Chinese character. It can include numbers, letters, Chinese characters, hyphens (-), and underscores (_).
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The location of the file containing the template body. The URL must point to a template located on a web server (HTTP or HTTPS) or in an Alibaba Cloud OSS bucket (e.g., oss://ros/template/demo, oss://ros/template/demo?RegionId=cn-hangzhou), with a maximum size of 524,288 bytes.
	//
	// > If the OSS region is not specified, it defaults to the same as the `RegionId` parameter in the request.
	//
	//
	//
	// You can only specify one of `TemplateBody` or `TemplateURL`.
	//
	// The maximum length of the URL is 1,024 bytes.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// Validation options.
	//
	// By default, no options are enabled, and strict validation is performed.
	ValidationOptions []*string `json:"ValidationOptions,omitempty" xml:"ValidationOptions,omitempty" type:"Repeated"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateRequest) GetIsDraft() *bool {
	return s.IsDraft
}

func (s *UpdateTemplateRequest) GetRotateStrategy() *string {
	return s.RotateStrategy
}

func (s *UpdateTemplateRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *UpdateTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateTemplateRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *UpdateTemplateRequest) GetValidationOptions() []*string {
	return s.ValidationOptions
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetIsDraft(v bool) *UpdateTemplateRequest {
	s.IsDraft = &v
	return s
}

func (s *UpdateTemplateRequest) SetRotateStrategy(v string) *UpdateTemplateRequest {
	s.RotateStrategy = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateBody(v string) *UpdateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateURL(v string) *UpdateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateTemplateRequest) SetValidationOptions(v []*string) *UpdateTemplateRequest {
	s.ValidationOptions = v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
