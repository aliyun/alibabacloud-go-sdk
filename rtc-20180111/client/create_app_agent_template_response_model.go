// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAgentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppAgentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppAgentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppAgentTemplateResponseBody) *CreateAppAgentTemplateResponse
	GetBody() *CreateAppAgentTemplateResponseBody
}

type CreateAppAgentTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppAgentTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppAgentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppAgentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppAgentTemplateResponse) GetBody() *CreateAppAgentTemplateResponseBody {
	return s.Body
}

func (s *CreateAppAgentTemplateResponse) SetHeaders(v map[string]*string) *CreateAppAgentTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppAgentTemplateResponse) SetStatusCode(v int32) *CreateAppAgentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppAgentTemplateResponse) SetBody(v *CreateAppAgentTemplateResponseBody) *CreateAppAgentTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAppAgentTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
