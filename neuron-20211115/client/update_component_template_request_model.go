// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ComponentTemplateUpdateCmd) *UpdateComponentTemplateRequest
	GetBody() *ComponentTemplateUpdateCmd
}

type UpdateComponentTemplateRequest struct {
	// This parameter is required.
	Body *ComponentTemplateUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentTemplateRequest) GetBody() *ComponentTemplateUpdateCmd {
	return s.Body
}

func (s *UpdateComponentTemplateRequest) SetBody(v *ComponentTemplateUpdateCmd) *UpdateComponentTemplateRequest {
	s.Body = v
	return s
}

func (s *UpdateComponentTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
