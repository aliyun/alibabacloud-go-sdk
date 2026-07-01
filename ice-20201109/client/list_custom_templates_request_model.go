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
	// The name of the template.
	//
	// example:
	//
	// test-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order of the results. Valid values:
	//
	// - `CreationTime:Desc`: Sorts the results by Creation Time in descending order.
	//
	// - `CreationTime:Asc`: Sorts the results by Creation Time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number of the results to return.
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
	// The subtype of the template. This parameter applies only when `Type` is set to 1, 2, 7, or 9.
	//
	// - Transcoding Template subtypes:
	//
	//   - 1: Normal (`Normal`)
	//
	//   - 2: Audio-only (`AudioTranscode`)
	//
	//   - 3: Remuxing (`Remux`)
	//
	//   - 4: Narrowband HD 1.0 (`NarrowBandV1`)
	//
	//   - 5: Narrowband HD 2.0 (`NarrowBandV2`)
	//
	// - Screenshot Template subtypes:
	//
	//   - 1: Normal (`Normal`)
	//
	//   - 2: Sprite Image (`Sprite`)
	//
	//   - 3: WebVTT (`WebVtt`)
	//
	// - AI Content Moderation subtypes:
	//
	//   - 1: Video moderation (`Video`)
	//
	//   - 2: Audio moderation (`Audio`)
	//
	//   - 3: Image moderation (`Image`)
	//
	// - AI-powered Object Removal subtypes:
	//
	//   - 1: Logo Removal (`VideoDelogo`)
	//
	//   - 2: Text Removal (`VideoDetext`)
	//
	// example:
	//
	// 2
	Subtype *string `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the template. Valid values:
	//
	// - 1: Transcoding Template
	//
	// - 2: Screenshot Template
	//
	// - 3: Animated GIF Template
	//
	// - 4: Image Watermark Template
	//
	// - 5: Text Watermark Template
	//
	// - 6: Subtitle Template
	//
	// - 7: AI Content Moderation
	//
	// - 8: AI-powered Smart Cover
	//
	// - 9: AI-powered Object Removal
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
