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
	// Valid values for transcoding templates:
	//
	// 	- 1 (Normal): regular template.
	//
	// 	- 2 (AudioTranscode): audio transcoding template.
	//
	// 	- 3 (Remux): container format conversion template.
	//
	// 	- 4 (NarrowBandV1): Narrowband HD 1.0 template.
	//
	// 	- 5 (NarrowBandV2): Narrowband HD 2.0 template.
	//
	// Valid values for snapshot templates:
	//
	// 	- 1 (Normal): regular template.
	//
	// 	- 2 (Sprite): sprite template.
	//
	// 	- 3 (WebVtt): WebVTT template.
	//
	// Valid values for AI-assisted content moderation templates:
	//
	// 	- 1 (Video): video moderation template.
	//
	// 	- 2 (Audio): audio moderation template.
	//
	// 	- 3 (Image): image moderation template.
	//
	// Valid values for AI-assisted intelligent erasure templates.
	//
	// 	- 1 (VideoDelogo): logo erasure template.
	//
	// 	- 2 (VideoDetext): subtitle erasure template.
	//
	// example:
	//
	// 1
	Subtype *int32 `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template configurations. For more information, see [Template parameters](https://help.aliyun.com/document_detail/448291.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Container":{"Format":"flv"},"Video":{},"Audio":{}}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template type. Valid values:
	//
	// 	- 1: transcoding template.
	//
	// 	- 2: snapshot template.
	//
	// 	- 3: animated image template.
	//
	// 	- 4\\. image watermark template.
	//
	// 	- 5: text watermark template.
	//
	// 	- 6: subtitle template.
	//
	// 	- 7: AI-assisted content moderation template.
	//
	// 	- 8: AI-assisted intelligent thumbnail template.
	//
	// 	- 9: AI-assisted intelligent erasure template.
	//
	// 	- 10: AI-assisted media fingerprint analysis template.
	//
	// 	- 11: AI-assisted smart tagging template.
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
