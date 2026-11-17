// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationTemplateResponseBody) *UpdateApplicationTemplateResponse
	GetBody() *UpdateApplicationTemplateResponseBody
}

type UpdateApplicationTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationTemplateResponse) GetBody() *UpdateApplicationTemplateResponseBody {
	return s.Body
}

func (s *UpdateApplicationTemplateResponse) SetHeaders(v map[string]*string) *UpdateApplicationTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationTemplateResponse) SetStatusCode(v int32) *UpdateApplicationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationTemplateResponse) SetBody(v *UpdateApplicationTemplateResponseBody) *UpdateApplicationTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
