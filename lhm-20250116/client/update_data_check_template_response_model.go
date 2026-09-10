// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataCheckTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataCheckTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataCheckTemplateResponseBody) *UpdateDataCheckTemplateResponse
	GetBody() *UpdateDataCheckTemplateResponseBody
}

type UpdateDataCheckTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataCheckTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataCheckTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataCheckTemplateResponse) GetBody() *UpdateDataCheckTemplateResponseBody {
	return s.Body
}

func (s *UpdateDataCheckTemplateResponse) SetHeaders(v map[string]*string) *UpdateDataCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataCheckTemplateResponse) SetStatusCode(v int32) *UpdateDataCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataCheckTemplateResponse) SetBody(v *UpdateDataCheckTemplateResponseBody) *UpdateDataCheckTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateDataCheckTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
