// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSandboxTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSandboxTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSandboxTemplateResponseBody) *DeleteSandboxTemplateResponse
	GetBody() *DeleteSandboxTemplateResponseBody
}

type DeleteSandboxTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSandboxTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSandboxTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSandboxTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSandboxTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSandboxTemplateResponse) GetBody() *DeleteSandboxTemplateResponseBody {
	return s.Body
}

func (s *DeleteSandboxTemplateResponse) SetHeaders(v map[string]*string) *DeleteSandboxTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSandboxTemplateResponse) SetStatusCode(v int32) *DeleteSandboxTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSandboxTemplateResponse) SetBody(v *DeleteSandboxTemplateResponseBody) *DeleteSandboxTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteSandboxTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
