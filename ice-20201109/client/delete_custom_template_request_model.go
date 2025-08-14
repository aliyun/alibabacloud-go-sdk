// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteCustomTemplateRequest
	GetTemplateId() *string
}

type DeleteCustomTemplateRequest struct {
	// The ID of the custom template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteCustomTemplateRequest) SetTemplateId(v string) *DeleteCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
