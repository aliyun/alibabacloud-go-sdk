// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDataProcessJobTemplatesResponseBodyItems) *DescribeDataProcessJobTemplatesResponseBody
	GetItems() []*DescribeDataProcessJobTemplatesResponseBodyItems
	SetPageNumber(v int32) *DescribeDataProcessJobTemplatesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDataProcessJobTemplatesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDataProcessJobTemplatesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDataProcessJobTemplatesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDataProcessJobTemplatesResponseBody struct {
	// The list of data processing task templates.
	Items []*DescribeDataProcessJobTemplatesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of templates returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of templates that meet the specified conditions.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDataProcessJobTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobTemplatesResponseBody) GetItems() []*DescribeDataProcessJobTemplatesResponseBodyItems {
	return s.Items
}

func (s *DescribeDataProcessJobTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataProcessJobTemplatesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDataProcessJobTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataProcessJobTemplatesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDataProcessJobTemplatesResponseBody) SetItems(v []*DescribeDataProcessJobTemplatesResponseBodyItems) *DescribeDataProcessJobTemplatesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBody) SetPageNumber(v int32) *DescribeDataProcessJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBody) SetPageRecordCount(v int32) *DescribeDataProcessJobTemplatesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBody) SetRequestId(v string) *DescribeDataProcessJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBody) SetTotalRecordCount(v int32) *DescribeDataProcessJobTemplatesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataProcessJobTemplatesResponseBodyItems struct {
	// Specifies whether it is required to fetch the summary or video synopsis.
	//
	// example:
	//
	// true
	Captioning *bool `json:"Captioning,omitempty" xml:"Captioning,omitempty"`
	// The prompt used when extracting captions.
	//
	// example:
	//
	// Summarize the image content
	CaptioningPrompt *string `json:"CaptioningPrompt,omitempty" xml:"CaptioningPrompt,omitempty"`
	// The time when the task template was created, in RFC 3339 format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Specifies whether to convert content into vectors.
	//
	// example:
	//
	// true
	Embedding *bool `json:"Embedding,omitempty" xml:"Embedding,omitempty"`
	// The output dimension of the vector.
	//
	// example:
	//
	// 768
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// The task template status.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The JSON configuration string of the template.
	//
	// example:
	//
	// {"ETL_TEMPLATE_STR":"config"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The task template ID.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The task template name.
	//
	// example:
	//
	// template1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The node templatetype. Valid value: image.
	//
	// example:
	//
	// image
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDataProcessJobTemplatesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobTemplatesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetCaptioning() *bool {
	return s.Captioning
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetCaptioningPrompt() *string {
	return s.CaptioningPrompt
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetEmbedding() *bool {
	return s.Embedding
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetCaptioning(v bool) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.Captioning = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetCaptioningPrompt(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.CaptioningPrompt = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetCreationTime(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetEmbedding(v bool) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.Embedding = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetEmbeddingDimension(v int32) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.EmbeddingDimension = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetStatus(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetTemplateConfig(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.TemplateConfig = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetTemplateId(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetTemplateName(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.TemplateName = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) SetTemplateType(v string) *DescribeDataProcessJobTemplatesResponseBodyItems {
	s.TemplateType = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
