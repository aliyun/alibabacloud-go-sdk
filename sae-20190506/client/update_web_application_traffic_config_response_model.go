// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationTrafficConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebApplicationTrafficConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebApplicationTrafficConfigResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationTrafficConfigBody) *UpdateWebApplicationTrafficConfigResponse
	GetBody() *WebApplicationTrafficConfigBody
}

type UpdateWebApplicationTrafficConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationTrafficConfigBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationTrafficConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationTrafficConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationTrafficConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebApplicationTrafficConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebApplicationTrafficConfigResponse) GetBody() *WebApplicationTrafficConfigBody {
	return s.Body
}

func (s *UpdateWebApplicationTrafficConfigResponse) SetHeaders(v map[string]*string) *UpdateWebApplicationTrafficConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebApplicationTrafficConfigResponse) SetStatusCode(v int32) *UpdateWebApplicationTrafficConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebApplicationTrafficConfigResponse) SetBody(v *WebApplicationTrafficConfigBody) *UpdateWebApplicationTrafficConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationTrafficConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
