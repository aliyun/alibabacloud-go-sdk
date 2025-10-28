// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigTemplateResponseBody) *UpdateConfigTemplateResponse
	GetBody() *UpdateConfigTemplateResponseBody
}

type UpdateConfigTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigTemplateResponse) GetBody() *UpdateConfigTemplateResponseBody {
	return s.Body
}

func (s *UpdateConfigTemplateResponse) SetHeaders(v map[string]*string) *UpdateConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigTemplateResponse) SetStatusCode(v int32) *UpdateConfigTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigTemplateResponse) SetBody(v *UpdateConfigTemplateResponseBody) *UpdateConfigTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
