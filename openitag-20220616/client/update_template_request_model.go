// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *TemplateDTO) *UpdateTemplateRequest
	GetBody() *TemplateDTO
}

type UpdateTemplateRequest struct {
	// template
	Body *TemplateDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetBody() *TemplateDTO {
	return s.Body
}

func (s *UpdateTemplateRequest) SetBody(v *TemplateDTO) *UpdateTemplateRequest {
	s.Body = v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
