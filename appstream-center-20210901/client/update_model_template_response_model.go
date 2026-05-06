// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelTemplateResponseBody) *UpdateModelTemplateResponse
	GetBody() *UpdateModelTemplateResponseBody
}

type UpdateModelTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelTemplateResponse) GetBody() *UpdateModelTemplateResponseBody {
	return s.Body
}

func (s *UpdateModelTemplateResponse) SetHeaders(v map[string]*string) *UpdateModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelTemplateResponse) SetStatusCode(v int32) *UpdateModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelTemplateResponse) SetBody(v *UpdateModelTemplateResponseBody) *UpdateModelTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
