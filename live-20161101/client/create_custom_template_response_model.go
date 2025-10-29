// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomTemplateResponseBody) *CreateCustomTemplateResponse
	GetBody() *CreateCustomTemplateResponseBody
}

type CreateCustomTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomTemplateResponse) GetBody() *CreateCustomTemplateResponseBody {
	return s.Body
}

func (s *CreateCustomTemplateResponse) SetHeaders(v map[string]*string) *CreateCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomTemplateResponse) SetStatusCode(v int32) *CreateCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomTemplateResponse) SetBody(v *CreateCustomTemplateResponseBody) *CreateCustomTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCustomTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
