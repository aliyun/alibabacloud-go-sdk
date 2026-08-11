// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateCustomTemplateRequest
	GetName() *string
	SetSubtype(v int32) *CreateCustomTemplateRequest
	GetSubtype() *int32
	SetTemplateConfig(v string) *CreateCustomTemplateRequest
	GetTemplateConfig() *string
	SetType(v int32) *CreateCustomTemplateRequest
	GetType() *int32
}

type CreateCustomTemplateRequest struct {
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template subtype.
	//
	// Transcoding template subtypes:
	//
	// - 1: normal transcoding template (Normal)
	//
	// - 2: audio transcoding template (AudioTranscode)
	//
	// - 3: container format conversion (Remux)
	//
	// - 4: Narrowband HD 1.0 (NarrowBandV1)
	//
	// - 5: Narrowband HD 2.0 (NarrowBandV2)
	//
	// Snapshot template subtypes:
	//
	// - 1: normal snapshot/static snapshot (Normal)
	//
	// - 2: sprite snapshot (Sprite)
	//
	// - 3: WebVTT snapshot (WebVtt)
	//
	// AI review template subtypes:
	//
	// - 1: video review (Video)
	//
	// - 2: audio review (Audio)
	//
	// - 3: image review (Image)
	//
	// AI intelligent erasure template subtypes:
	//
	// - 1: logo erasure (VideoDelogo)
	//
	// - 2: subtitle erasure (VideoDetext)
	//
	// example:
	//
	// 1
	Subtype *int32 `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template configuration. For detailed metric description, see [Template parameters](https://help.aliyun.com/document_detail/448291.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Container":{"Format":"flv"},"Video":{},"Audio":{}}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The templatetype. Valid values:
	//
	// - 1: transcoding template
	//
	// - 2: snapshot template
	//
	// - 3: animated image template
	//
	// - 4: image watermark template
	//
	// - 5: text watermark template
	//
	// - 6: subtitle template
	//
	// - 7: AI intelligent review
	//
	// - 8: AI intelligent cover
	//
	// - 9: AI intelligent erasure
	//
	// - 10: AI intelligent DNA template
	//
	// - 11: AI intelligent label template
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomTemplateRequest) GetSubtype() *int32 {
	return s.Subtype
}

func (s *CreateCustomTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *CreateCustomTemplateRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateCustomTemplateRequest) SetName(v string) *CreateCustomTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetSubtype(v int32) *CreateCustomTemplateRequest {
	s.Subtype = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetTemplateConfig(v string) *CreateCustomTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetType(v int32) *CreateCustomTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
