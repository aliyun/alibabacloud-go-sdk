// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarClawSkillTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarClawSkillTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarClawSkillTemplateResponseBody) *CreatePolarClawSkillTemplateResponse
	GetBody() *CreatePolarClawSkillTemplateResponseBody
}

type CreatePolarClawSkillTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarClawSkillTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarClawSkillTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarClawSkillTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarClawSkillTemplateResponse) GetBody() *CreatePolarClawSkillTemplateResponseBody {
	return s.Body
}

func (s *CreatePolarClawSkillTemplateResponse) SetHeaders(v map[string]*string) *CreatePolarClawSkillTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarClawSkillTemplateResponse) SetStatusCode(v int32) *CreatePolarClawSkillTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarClawSkillTemplateResponse) SetBody(v *CreatePolarClawSkillTemplateResponseBody) *CreatePolarClawSkillTemplateResponse {
	s.Body = v
	return s
}

func (s *CreatePolarClawSkillTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
