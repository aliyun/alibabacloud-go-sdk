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
	// The list of templates.
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
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTemplatesResponseBodyPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 20
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The maximum number of entries per page.
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
	// The access permissions for the deployment template. Valid values:
	//
	// - `private`: private.
	//
	// - `public`: public.
	//
	// - `shared`: shared.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The time when the orchestration template was created.
	//
	// example:
	//
	// 2025-04-25T16:56:33+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the orchestration template.
	//
	// example:
	//
	// a web server
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the orchestration template.
	//
	// example:
	//
	// 874ec485-e7e6-4373-8a3b-47bde8******
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the orchestration template.
	//
	// example:
	//
	// webserver
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tag of the orchestration template. If not explicitly specified, the tag defaults to the template name.
	//
	// example:
	//
	// kubernetes
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The template content in YAML format.
	//
	// example:
	//
	// apiVersion: apps/v1\\nkind: Deployment\\nmetadata:\\n  name: nginx-deployment-basic\\n  labels:\\n    app: nginx\\nspec:\\n  replicas: 2\\n  selector:\\n    matchLabels:\\n      app: nginx\\n  template:\\n    metadata:\\n      labels:\\n        app: nginx\\n    spec:\\n      containers:\\n      - name: nginx\\n        image: busybox:latest\\n        ports:\\n        - containerPort: 80
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The templatetype.
	//
	// - If the value is set to kubernetes, the template is displayed on the Orchestration Templates page in the console.
	//
	// - If this parameter is left empty or set to other values, the template is not displayed on the Orchestration Templates page in the console.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// The ID of the parent template associated with the template. This parameter is used to implement template versioning. Different versions of the same template share the same `template_with_hist_id` value.
	//
	// example:
	//
	// ad81d115-7c8b-47e7-a222-9c28d7******
	TemplateWithHistId *string `json:"template_with_hist_id,omitempty" xml:"template_with_hist_id,omitempty"`
	// The time when the orchestration template was last updated.
	//
	// example:
	//
	// 2025-04-25T16:56:33+08:00
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
