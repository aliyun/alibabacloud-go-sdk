// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplate(v string) *CreateCustomTemplateRequest
	GetCustomTemplate() *string
	SetOwnerId(v int64) *CreateCustomTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCustomTemplateRequest
	GetRegionId() *string
	SetTemplate(v string) *CreateCustomTemplateRequest
	GetTemplate() *string
}

type CreateCustomTemplateRequest struct {
	// The configuration of the template. The value is in the following JSON format: {height:xxx,scale:xxx,gop:xxx,bframes:xxx,cdesc:xxx}. All fields are required. If any field is left empty, the call fails.
	//
	// >  For more information, see **Fields of the CustomTemplate parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// {height:1080,scale:[16:9],gop:60,bframes:30,cdesc:h264}
	CustomTemplate *string `json:"CustomTemplate,omitempty" xml:"CustomTemplate,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template.
	//
	// > Record the template name. The template name is required if you want to use, query, or delete the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestTemplate
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateRequest) GetCustomTemplate() *string {
	return s.CustomTemplate
}

func (s *CreateCustomTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateCustomTemplateRequest) SetCustomTemplate(v string) *CreateCustomTemplateRequest {
	s.CustomTemplate = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetOwnerId(v int64) *CreateCustomTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetRegionId(v string) *CreateCustomTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetTemplate(v string) *CreateCustomTemplateRequest {
	s.Template = &v
	return s
}

func (s *CreateCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
