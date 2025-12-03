// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePluginTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v *DescribePluginTemplatesResponseBodyTemplates) *DescribePluginTemplatesResponseBody
	GetTemplates() *DescribePluginTemplatesResponseBodyTemplates
}

type DescribePluginTemplatesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The templates.
	Templates *DescribePluginTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
}

func (s DescribePluginTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginTemplatesResponseBody) GetTemplates() *DescribePluginTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribePluginTemplatesResponseBody) SetRequestId(v string) *DescribePluginTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginTemplatesResponseBody) SetTemplates(v *DescribePluginTemplatesResponseBodyTemplates) *DescribePluginTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribePluginTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePluginTemplatesResponseBodyTemplates struct {
	Template []*DescribePluginTemplatesResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s DescribePluginTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBodyTemplates) GetTemplate() []*DescribePluginTemplatesResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *DescribePluginTemplatesResponseBodyTemplates) SetTemplate(v []*DescribePluginTemplatesResponseBodyTemplatesTemplate) *DescribePluginTemplatesResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplates) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePluginTemplatesResponseBodyTemplatesTemplate struct {
	// The description.
	//
	// example:
	//
	// balabala
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document anchor point.
	//
	// example:
	//
	// anchor
	DocumentAnchor *string `json:"DocumentAnchor,omitempty" xml:"DocumentAnchor,omitempty"`
	// The ID of the document.
	//
	// example:
	//
	// 41079
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The sample.
	//
	// example:
	//
	// 1
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The title of the plug-in template title.
	//
	// example:
	//
	// template title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePluginTemplatesResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginTemplatesResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) GetDocumentAnchor() *string {
	return s.DocumentAnchor
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) GetDocumentId() *string {
	return s.DocumentId
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) GetSample() *string {
	return s.Sample
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) GetTitle() *string {
	return s.Title
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDescription(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Description = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDocumentAnchor(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.DocumentAnchor = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDocumentId(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.DocumentId = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetSample(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Sample = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetTitle(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Title = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) Validate() error {
	return dara.Validate(s)
}
