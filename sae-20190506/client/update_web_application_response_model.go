// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationBody) *UpdateWebApplicationResponse
	GetBody() *WebApplicationBody
}

type UpdateWebApplicationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebApplicationResponse) GetBody() *WebApplicationBody {
	return s.Body
}

func (s *UpdateWebApplicationResponse) SetHeaders(v map[string]*string) *UpdateWebApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebApplicationResponse) SetStatusCode(v int32) *UpdateWebApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebApplicationResponse) SetBody(v *WebApplicationBody) *UpdateWebApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
