// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebApplicationScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebApplicationScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationScalingConfigBody) *UpdateWebApplicationScalingConfigResponse
	GetBody() *WebApplicationScalingConfigBody
}

type UpdateWebApplicationScalingConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationScalingConfigBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebApplicationScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebApplicationScalingConfigResponse) GetBody() *WebApplicationScalingConfigBody {
	return s.Body
}

func (s *UpdateWebApplicationScalingConfigResponse) SetHeaders(v map[string]*string) *UpdateWebApplicationScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebApplicationScalingConfigResponse) SetStatusCode(v int32) *UpdateWebApplicationScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebApplicationScalingConfigResponse) SetBody(v *WebApplicationScalingConfigBody) *UpdateWebApplicationScalingConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationScalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
