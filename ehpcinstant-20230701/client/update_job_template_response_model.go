// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobTemplateResponseBody) *UpdateJobTemplateResponse
	GetBody() *UpdateJobTemplateResponseBody
}

type UpdateJobTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobTemplateResponse) GetBody() *UpdateJobTemplateResponseBody {
	return s.Body
}

func (s *UpdateJobTemplateResponse) SetHeaders(v map[string]*string) *UpdateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobTemplateResponse) SetStatusCode(v int32) *UpdateJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobTemplateResponse) SetBody(v *UpdateJobTemplateResponseBody) *UpdateJobTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
