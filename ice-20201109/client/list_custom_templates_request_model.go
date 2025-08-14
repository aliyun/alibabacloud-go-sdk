// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListCustomTemplatesRequest
	GetName() *string
	SetOrderBy(v string) *ListCustomTemplatesRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListCustomTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomTemplatesRequest
	GetPageSize() *int32
	SetSubtype(v string) *ListCustomTemplatesRequest
	GetSubtype() *string
	SetTemplateId(v string) *ListCustomTemplatesRequest
	GetTemplateId() *string
	SetType(v string) *ListCustomTemplatesRequest
	GetType() *string
}

type ListCustomTemplatesRequest struct {
	// The template name.
	//
	// example:
	//
	// test-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- CreateTimeDesc: sorted by creation time in descending order.
	//
	// 	- CreateTimeAsc: sorted by creation time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The subtype ID of the template.
	//
	// 	- Valid values for transcoding templates:
	//
	//     	- 1 (Normal): regular template.
	//
	//     	- 2 (AudioTranscode): audio transcoding template.
	//
	//     	- 3 (Remux): container format conversion template.
	//
	//     	- 4 (NarrowBandV1): Narrowband HD 1.0 template.
	//
	//     	- 5 (NarrowBandV2): Narrowband HD 2.0 template.
	//
	// 	- Valid values for snapshot templates:
	//
	//     	- 1 (Normal): regular template.
	//
	//     	- 2 (Sprite): sprite template.
	//
	//     	- 3 (WebVtt): WebVTT template.
	//
	// 	- Valid values for AI-assisted content moderation templates:
	//
	//     	- 1 (Video): video moderation template.
	//
	//     	- 2 (Audio): audio moderation template.
	//
	//     	- 3 (Image): image moderation template.
	//
	// 	- Valid values for AI-assisted intelligent erasure templates:
	//
	//     	- 1 (VideoDelogo): logo erasure template.
	//
	//     	- 2 (VideoDetext): subtitle erasure template.
	//
	// example:
	//
	// 2
	Subtype *string `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListCustomTemplatesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListCustomTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomTemplatesRequest) GetSubtype() *string {
	return s.Subtype
}

func (s *ListCustomTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCustomTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListCustomTemplatesRequest) SetName(v string) *ListCustomTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetOrderBy(v string) *ListCustomTemplatesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetPageNumber(v int32) *ListCustomTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetPageSize(v int32) *ListCustomTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetSubtype(v string) *ListCustomTemplatesRequest {
	s.Subtype = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetTemplateId(v string) *ListCustomTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListCustomTemplatesRequest) SetType(v string) *ListCustomTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListCustomTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
