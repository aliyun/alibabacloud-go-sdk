// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptioning(v bool) *CreateDataProcessJobTemplateRequest
	GetCaptioning() *bool
	SetCaptioningPrompt(v string) *CreateDataProcessJobTemplateRequest
	GetCaptioningPrompt() *string
	SetEmbedding(v bool) *CreateDataProcessJobTemplateRequest
	GetEmbedding() *bool
	SetEmbeddingDimension(v int32) *CreateDataProcessJobTemplateRequest
	GetEmbeddingDimension() *int32
	SetRegionId(v string) *CreateDataProcessJobTemplateRequest
	GetRegionId() *string
	SetTemplateConfig(v string) *CreateDataProcessJobTemplateRequest
	GetTemplateConfig() *string
	SetTemplateName(v string) *CreateDataProcessJobTemplateRequest
	GetTemplateName() *string
	SetTemplateType(v string) *CreateDataProcessJobTemplateRequest
	GetTemplateType() *string
}

type CreateDataProcessJobTemplateRequest struct {
	// Specifies whether to fetch the image summary or video synopsis. When Settings is true, CaptioningPrompt must be provided.
	Captioning *bool `json:"Captioning,omitempty" xml:"Captioning,omitempty"`
	// The prompt used for extracting captions. This parameter takes effect only when Captioning is set to true.
	//
	// example:
	//
	// Summarize the image content
	CaptioningPrompt *string `json:"CaptioningPrompt,omitempty" xml:"CaptioningPrompt,omitempty"`
	// Specifies whether to convert the content into vectors.
	Embedding *bool `json:"Embedding,omitempty" xml:"Embedding,omitempty"`
	// The output dimension of the vector. Valid values when Embedding is set to true: 512, 768, 1024, and 2048.
	//
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The JSON string for additional template configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ETL_TEMPLATE_STR":"config"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The name of the task template.
	//
	// example:
	//
	// template1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The templatetype of the node template. Valid value: image.
	//
	// This parameter is required.
	//
	// example:
	//
	// image
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateDataProcessJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobTemplateRequest) GetCaptioning() *bool {
	return s.Captioning
}

func (s *CreateDataProcessJobTemplateRequest) GetCaptioningPrompt() *string {
	return s.CaptioningPrompt
}

func (s *CreateDataProcessJobTemplateRequest) GetEmbedding() *bool {
	return s.Embedding
}

func (s *CreateDataProcessJobTemplateRequest) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *CreateDataProcessJobTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataProcessJobTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *CreateDataProcessJobTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateDataProcessJobTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateDataProcessJobTemplateRequest) SetCaptioning(v bool) *CreateDataProcessJobTemplateRequest {
	s.Captioning = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetCaptioningPrompt(v string) *CreateDataProcessJobTemplateRequest {
	s.CaptioningPrompt = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetEmbedding(v bool) *CreateDataProcessJobTemplateRequest {
	s.Embedding = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetEmbeddingDimension(v int32) *CreateDataProcessJobTemplateRequest {
	s.EmbeddingDimension = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetRegionId(v string) *CreateDataProcessJobTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetTemplateConfig(v string) *CreateDataProcessJobTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetTemplateName(v string) *CreateDataProcessJobTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) SetTemplateType(v string) *CreateDataProcessJobTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateDataProcessJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
