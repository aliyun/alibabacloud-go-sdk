// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebFilterTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebFilterTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebFilterTemplateResponseBody) *CreateWebFilterTemplateResponse
	GetBody() *CreateWebFilterTemplateResponseBody
}

type CreateWebFilterTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebFilterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebFilterTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateWebFilterTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebFilterTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebFilterTemplateResponse) GetBody() *CreateWebFilterTemplateResponseBody {
	return s.Body
}

func (s *CreateWebFilterTemplateResponse) SetHeaders(v map[string]*string) *CreateWebFilterTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateWebFilterTemplateResponse) SetStatusCode(v int32) *CreateWebFilterTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebFilterTemplateResponse) SetBody(v *CreateWebFilterTemplateResponseBody) *CreateWebFilterTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateWebFilterTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
