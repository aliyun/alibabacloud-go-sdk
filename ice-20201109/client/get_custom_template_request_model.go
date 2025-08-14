// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubtype(v int32) *GetCustomTemplateRequest
	GetSubtype() *int32
	SetTemplateId(v string) *GetCustomTemplateRequest
	GetTemplateId() *string
	SetType(v int32) *GetCustomTemplateRequest
	GetType() *int32
}

type GetCustomTemplateRequest struct {
	// The template subtype.
	//
	// example:
	//
	// 1
	Subtype *int32 `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the template type that is used to query the default template. This parameter is required if TemplateId is not specified.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateRequest) GetSubtype() *int32 {
	return s.Subtype
}

func (s *GetCustomTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetCustomTemplateRequest) GetType() *int32 {
	return s.Type
}

func (s *GetCustomTemplateRequest) SetSubtype(v int32) *GetCustomTemplateRequest {
	s.Subtype = &v
	return s
}

func (s *GetCustomTemplateRequest) SetTemplateId(v string) *GetCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCustomTemplateRequest) SetType(v int32) *GetCustomTemplateRequest {
	s.Type = &v
	return s
}

func (s *GetCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
