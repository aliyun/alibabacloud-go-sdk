// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSandboxTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSandboxTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateSandboxTemplateResponseBody) *CreateSandboxTemplateResponse
	GetBody() *CreateSandboxTemplateResponseBody
}

type CreateSandboxTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSandboxTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSandboxTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSandboxTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSandboxTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSandboxTemplateResponse) GetBody() *CreateSandboxTemplateResponseBody {
	return s.Body
}

func (s *CreateSandboxTemplateResponse) SetHeaders(v map[string]*string) *CreateSandboxTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSandboxTemplateResponse) SetStatusCode(v int32) *CreateSandboxTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSandboxTemplateResponse) SetBody(v *CreateSandboxTemplateResponseBody) *CreateSandboxTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateSandboxTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
