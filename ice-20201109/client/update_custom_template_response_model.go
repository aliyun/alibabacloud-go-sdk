// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomTemplateResponseBody) *UpdateCustomTemplateResponse
	GetBody() *UpdateCustomTemplateResponseBody
}

type UpdateCustomTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomTemplateResponse) GetBody() *UpdateCustomTemplateResponseBody {
	return s.Body
}

func (s *UpdateCustomTemplateResponse) SetHeaders(v map[string]*string) *UpdateCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomTemplateResponse) SetStatusCode(v int32) *UpdateCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomTemplateResponse) SetBody(v *UpdateCustomTemplateResponseBody) *UpdateCustomTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
