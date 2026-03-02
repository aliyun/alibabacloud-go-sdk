// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ComponentTemplateCreateCmd) *CreateComponentTemplateRequest
	GetBody() *ComponentTemplateCreateCmd
}

type CreateComponentTemplateRequest struct {
	// This parameter is required.
	Body *ComponentTemplateCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentTemplateRequest) GetBody() *ComponentTemplateCreateCmd {
	return s.Body
}

func (s *CreateComponentTemplateRequest) SetBody(v *ComponentTemplateCreateCmd) *CreateComponentTemplateRequest {
	s.Body = v
	return s
}

func (s *CreateComponentTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
