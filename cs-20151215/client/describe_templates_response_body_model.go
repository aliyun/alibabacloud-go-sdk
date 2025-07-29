// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeTemplatesResponseBodyPageInfo) *DescribeTemplatesResponseBody
	GetPageInfo() *DescribeTemplatesResponseBodyPageInfo
	SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody
	GetTemplates() []*DescribeTemplatesResponseBodyTemplates
}

type DescribeTemplatesResponseBody struct {
	// The pagination information.
	PageInfo *DescribeTemplatesResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// The list of returned templates.
	Templates []*DescribeTemplatesResponseBodyTemplates `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) GetPageInfo() *DescribeTemplatesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeTemplatesResponseBody) GetTemplates() []*DescribeTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeTemplatesResponseBody) SetPageInfo(v *DescribeTemplatesResponseBodyPageInfo) *DescribeTemplatesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplatesResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 20
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeTemplatesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyPageInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeTemplatesResponseBodyPageInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTemplatesResponseBodyPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageNumber(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageSize(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetTotalCount(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplatesResponseBodyTemplates struct {
	// The access control policy of the template. Valid values:
	//
	// 	- `private`: The template is private.
	//
	// 	- `public`: The template is public.
	//
	// 	- `shared`: The template can be shared.
	//
	// Default value: `private`.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2020-06-10T16:30:16+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// a web server
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 874ec485-e7e6-4373-8a3b-47bde8ae789f
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// webserver
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The label of the template. By default, the value is the name of the template.
	//
	// example:
	//
	// kubernetes
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The template content in the YAML format.
	//
	// example:
	//
	// apiVersion: apps/v1\\nkind: Deployment\\nmetadata:\\n  name: nginx-deployment-basic\\n  labels:\\n    app: nginx\\nspec:\\n  replicas: 2\\n  selector:\\n    matchLabels:\\n      app: nginx\\n  template:\\n    metadata:\\n      labels:\\n        app: nginx\\n    spec:\\n      containers:\\n      - name: nginx\\n        image: busybox:latest\\n        ports:\\n        - containerPort: 80
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The type of template. This parameter can be set to a custom value.
	//
	// 	- If the parameter is set to `kubernetes`, the template is displayed on the Templates page in the console.
	//
	// 	- If the parameter is set to `compose`, the template is displayed on the Container Service - Swarm page in the console. However, Container Service for Swarm is deprecated.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// The ID of the parent template. The value of `template_with_hist_id` is the same for each template version. This allows you to manage different template versions.
	//
	// example:
	//
	// ad81d115-7c8b-47e7-a222-9c28d7f6e588
	TemplateWithHistId *string `json:"template_with_hist_id,omitempty" xml:"template_with_hist_id,omitempty"`
	// The time when the template was updated.
	//
	// example:
	//
	// 2020-06-10T16:30:16+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplates) GetAcl() *string {
	return s.Acl
}

func (s *DescribeTemplatesResponseBodyTemplates) GetCreated() *string {
	return s.Created
}

func (s *DescribeTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplatesResponseBodyTemplates) GetId() *string {
	return s.Id
}

func (s *DescribeTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTags() *string {
	return s.Tags
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTemplate() *string {
	return s.Template
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatesResponseBodyTemplates) GetTemplateWithHistId() *string {
	return s.TemplateWithHistId
}

func (s *DescribeTemplatesResponseBodyTemplates) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeTemplatesResponseBodyTemplates) SetAcl(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Acl = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCreated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Created = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetName(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTags(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Tags = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplate(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Template = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplateWithHistId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.TemplateWithHistId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetUpdated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Updated = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
