// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSystemTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetSystemTemplateRequest
	GetTemplateId() *string
}

type GetSystemTemplateRequest struct {
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// S00000001-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetSystemTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSystemTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetSystemTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetSystemTemplateRequest) SetTemplateId(v string) *GetSystemTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetSystemTemplateRequest) Validate() error {
	return dara.Validate(s)
}
