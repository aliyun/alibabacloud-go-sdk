// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebFilterTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebFilterTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWebFilterTemplateResponseBody) *UpdateWebFilterTemplateResponse
	GetBody() *UpdateWebFilterTemplateResponseBody
}

type UpdateWebFilterTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebFilterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebFilterTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebFilterTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebFilterTemplateResponse) GetBody() *UpdateWebFilterTemplateResponseBody {
	return s.Body
}

func (s *UpdateWebFilterTemplateResponse) SetHeaders(v map[string]*string) *UpdateWebFilterTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebFilterTemplateResponse) SetStatusCode(v int32) *UpdateWebFilterTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebFilterTemplateResponse) SetBody(v *UpdateWebFilterTemplateResponseBody) *UpdateWebFilterTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateWebFilterTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
