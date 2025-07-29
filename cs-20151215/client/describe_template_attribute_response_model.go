// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateAttributeResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeTemplateAttributeResponseBody) *DescribeTemplateAttributeResponse
	GetBody() []*DescribeTemplateAttributeResponseBody
}

type DescribeTemplateAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeTemplateAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateAttributeResponse) GetBody() []*DescribeTemplateAttributeResponseBody {
	return s.Body
}

func (s *DescribeTemplateAttributeResponse) SetHeaders(v map[string]*string) *DescribeTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateAttributeResponse) SetStatusCode(v int32) *DescribeTemplateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateAttributeResponse) SetBody(v []*DescribeTemplateAttributeResponseBody) *DescribeTemplateAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateAttributeResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplateAttributeResponseBody struct {
	// The ID of the template. When you update a template, a new template ID is generated.
	//
	// example:
	//
	// 72d20cf8-a533-4ea9-a10d-e7630d3d****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The access control policy of the template.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// web
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The template content in the YAML format.
	//
	// example:
	//
	// apiVersion: V1\\n***
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The type of template. The value can be a custom value.
	//
	// 	- If the parameter is set to `kubernetes`, the template is displayed on the Templates page in the console.
	//
	// 	- If the parameter is set to `compose`, the template is displayed on the Container Service - Swarm page in the console. Container Service for Swarm is deprecated.
	//
	// 	- If the value of the parameter is not `kubernetes`, the template is not displayed on the Templates page in the console. We recommend that you set the parameter to `kubernetes`.
	//
	// Default value: `kubernetes`.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// test template
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The label of the template.
	//
	// example:
	//
	// sa
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The unique ID of the template. The value remains unchanged after the template is updated.
	//
	// example:
	//
	// 874ec485-e7e6-4373-8a3b-47bde8ae****
	TemplateWithHistId *string `json:"template_with_hist_id,omitempty" xml:"template_with_hist_id,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2020-09-16T19:21:29+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The time when the template was updated.
	//
	// example:
	//
	// 2020-09-16T19:21:29+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeTemplateAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeTemplateAttributeResponseBody) GetAcl() *string {
	return s.Acl
}

func (s *DescribeTemplateAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeTemplateAttributeResponseBody) GetTemplate() *string {
	return s.Template
}

func (s *DescribeTemplateAttributeResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplateAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplateAttributeResponseBody) GetTags() *string {
	return s.Tags
}

func (s *DescribeTemplateAttributeResponseBody) GetTemplateWithHistId() *string {
	return s.TemplateWithHistId
}

func (s *DescribeTemplateAttributeResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeTemplateAttributeResponseBody) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeTemplateAttributeResponseBody) SetId(v string) *DescribeTemplateAttributeResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetAcl(v string) *DescribeTemplateAttributeResponseBody {
	s.Acl = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetName(v string) *DescribeTemplateAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetTemplate(v string) *DescribeTemplateAttributeResponseBody {
	s.Template = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetTemplateType(v string) *DescribeTemplateAttributeResponseBody {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetDescription(v string) *DescribeTemplateAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetTags(v string) *DescribeTemplateAttributeResponseBody {
	s.Tags = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetTemplateWithHistId(v string) *DescribeTemplateAttributeResponseBody {
	s.TemplateWithHistId = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetCreated(v string) *DescribeTemplateAttributeResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) SetUpdated(v string) *DescribeTemplateAttributeResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeTemplateAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
