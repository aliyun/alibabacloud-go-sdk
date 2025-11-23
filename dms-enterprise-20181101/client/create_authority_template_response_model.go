// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAuthorityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAuthorityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAuthorityTemplateResponseBody) *CreateAuthorityTemplateResponse
	GetBody() *CreateAuthorityTemplateResponseBody
}

type CreateAuthorityTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthorityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthorityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorityTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAuthorityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAuthorityTemplateResponse) GetBody() *CreateAuthorityTemplateResponseBody {
	return s.Body
}

func (s *CreateAuthorityTemplateResponse) SetHeaders(v map[string]*string) *CreateAuthorityTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorityTemplateResponse) SetStatusCode(v int32) *CreateAuthorityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorityTemplateResponse) SetBody(v *CreateAuthorityTemplateResponseBody) *CreateAuthorityTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAuthorityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
