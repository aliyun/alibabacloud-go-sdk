// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelProviderTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelProviderTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelProviderTemplateResponseBody) *UpdateModelProviderTemplateResponse
	GetBody() *UpdateModelProviderTemplateResponseBody
}

type UpdateModelProviderTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelProviderTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelProviderTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelProviderTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelProviderTemplateResponse) GetBody() *UpdateModelProviderTemplateResponseBody {
	return s.Body
}

func (s *UpdateModelProviderTemplateResponse) SetHeaders(v map[string]*string) *UpdateModelProviderTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelProviderTemplateResponse) SetStatusCode(v int32) *UpdateModelProviderTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelProviderTemplateResponse) SetBody(v *UpdateModelProviderTemplateResponseBody) *UpdateModelProviderTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateModelProviderTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
