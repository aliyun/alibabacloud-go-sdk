// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListSystemTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *ListSystemTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSystemTemplatesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListSystemTemplatesRequest
	GetStatus() *string
	SetSubtype(v string) *ListSystemTemplatesRequest
	GetSubtype() *string
	SetTemplateId(v string) *ListSystemTemplatesRequest
	GetTemplateId() *string
	SetType(v string) *ListSystemTemplatesRequest
	GetType() *string
}

type ListSystemTemplatesRequest struct {
	// The template name.
	//
	// example:
	//
	// SampleTemplate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20 Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The template state. Valid values: Normal, Invisible, and All.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype ID of the template.
	//
	// example:
	//
	// 1
	Subtype *string `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template type. Separate multiple types with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSystemTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListSystemTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSystemTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemTemplatesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSystemTemplatesRequest) GetSubtype() *string {
	return s.Subtype
}

func (s *ListSystemTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListSystemTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListSystemTemplatesRequest) SetName(v string) *ListSystemTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetPageNumber(v int32) *ListSystemTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetPageSize(v int32) *ListSystemTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetStatus(v string) *ListSystemTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetSubtype(v string) *ListSystemTemplatesRequest {
	s.Subtype = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetTemplateId(v string) *ListSystemTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListSystemTemplatesRequest) SetType(v string) *ListSystemTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListSystemTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
