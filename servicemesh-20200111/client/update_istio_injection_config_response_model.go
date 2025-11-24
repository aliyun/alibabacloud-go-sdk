// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioInjectionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIstioInjectionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIstioInjectionConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIstioInjectionConfigResponseBody) *UpdateIstioInjectionConfigResponse
	GetBody() *UpdateIstioInjectionConfigResponseBody
}

type UpdateIstioInjectionConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIstioInjectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIstioInjectionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioInjectionConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIstioInjectionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIstioInjectionConfigResponse) GetBody() *UpdateIstioInjectionConfigResponseBody {
	return s.Body
}

func (s *UpdateIstioInjectionConfigResponse) SetHeaders(v map[string]*string) *UpdateIstioInjectionConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioInjectionConfigResponse) SetStatusCode(v int32) *UpdateIstioInjectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioInjectionConfigResponse) SetBody(v *UpdateIstioInjectionConfigResponseBody) *UpdateIstioInjectionConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateIstioInjectionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
