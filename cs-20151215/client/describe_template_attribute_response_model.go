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
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplateAttributeResponseBody struct {
	// 编排模板ID，每次变更都会有一个模板ID。
	//
	// example:
	//
	// 72d20cf8-a533-4ea9-a10d-e7630d3d****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 编排模板的访问权限，取值：
	//
	// - `private`：私有。
	//
	// - `public`：公共。
	//
	// - `shared`：可共享。
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// 编排模板名称。
	//
	// example:
	//
	// web
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 编排模板YAML内容。
	//
	// example:
	//
	// apiVersion: V1\\n***
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 模板类型。
	//
	// - 当取值为kubernetes时将在控制台的编排模板页面展示该模板。
	//
	// - 该参数不填写或者取值为其他值时，控制台的编排模板页面将不会展示该模板。
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// 编排模板描述信息。
	//
	// example:
	//
	// test template
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 部署模板的标签。
	//
	// example:
	//
	// sa
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 编排模板唯一ID，不随模板更新而改变。
	//
	// example:
	//
	// 874ec485-e7e6-4373-8a3b-47bde8ae****
	TemplateWithHistId *string `json:"template_with_hist_id,omitempty" xml:"template_with_hist_id,omitempty"`
	// 编排模板创建时间。
	//
	// example:
	//
	// 2025-04-25T16:56:33+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 编排模板更新时间。
	//
	// example:
	//
	// 2025-04-25T16:56:33+08:00
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
