// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *TemplateDTO) *CreateTemplateRequest
	GetBody() *TemplateDTO
}

type CreateTemplateRequest struct {
	// Template information.
	//
	// This parameter is required.
	Body *TemplateDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetBody() *TemplateDTO {
	return s.Body
}

func (s *CreateTemplateRequest) SetBody(v *TemplateDTO) *CreateTemplateRequest {
	s.Body = v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
