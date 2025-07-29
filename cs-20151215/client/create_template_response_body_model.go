// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *CreateTemplateResponseBody
	GetTemplateId() *string
}

type CreateTemplateResponseBody struct {
	// The ID of the orchestration template.
	//
	// example:
	//
	// ba1fe77b-b01e-4640-b77e-8f1b80e3e3cf
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
