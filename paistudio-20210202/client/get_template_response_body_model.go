// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetTemplateResponseBody
	GetContent() *string
	SetDescription(v string) *GetTemplateResponseBody
	GetDescription() *string
	SetDetail(v string) *GetTemplateResponseBody
	GetDetail() *string
	SetDocLink(v string) *GetTemplateResponseBody
	GetDocLink() *string
	SetImageLink(v string) *GetTemplateResponseBody
	GetImageLink() *string
	SetLabels(v []map[string]interface{}) *GetTemplateResponseBody
	GetLabels() []map[string]interface{}
	SetName(v string) *GetTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetSourceId(v string) *GetTemplateResponseBody
	GetSourceId() *string
	SetSourceType(v string) *GetTemplateResponseBody
	GetSourceType() *string
	SetTemplateId(v string) *GetTemplateResponseBody
	GetTemplateId() *string
	SetTemplateType(v string) *GetTemplateResponseBody
	GetTemplateType() *string
}

type GetTemplateResponseBody struct {
	// example:
	//
	// {     "metadata": {       "name": "实验名称",       "id": "pai_exp_xxxdfafafasfa",       "desc": "实验描述",     },     "nodes": [     ],     "edges": [     ],     "globalParams": [     ],     "globalSettings":[     ]  }
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// URL://xxx
	DocLink *string `json:"DocLink,omitempty" xml:"DocLink,omitempty"`
	// example:
	//
	// URL://xxx
	ImageLink  *string                  `json:"ImageLink,omitempty" xml:"ImageLink,omitempty"`
	Labels     []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name       *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId  *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId   *string                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType *string                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// template-12345
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *GetTemplateResponseBody) GetDocLink() *string {
	return s.DocLink
}

func (s *GetTemplateResponseBody) GetImageLink() *string {
	return s.ImageLink
}

func (s *GetTemplateResponseBody) GetLabels() []map[string]interface{} {
	return s.Labels
}

func (s *GetTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetTemplateResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetTemplateResponseBody) SetContent(v string) *GetTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetDetail(v string) *GetTemplateResponseBody {
	s.Detail = &v
	return s
}

func (s *GetTemplateResponseBody) SetDocLink(v string) *GetTemplateResponseBody {
	s.DocLink = &v
	return s
}

func (s *GetTemplateResponseBody) SetImageLink(v string) *GetTemplateResponseBody {
	s.ImageLink = &v
	return s
}

func (s *GetTemplateResponseBody) SetLabels(v []map[string]interface{}) *GetTemplateResponseBody {
	s.Labels = v
	return s
}

func (s *GetTemplateResponseBody) SetName(v string) *GetTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSourceId(v string) *GetTemplateResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSourceType(v string) *GetTemplateResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateType(v string) *GetTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
