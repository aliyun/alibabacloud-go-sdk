// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallManagedPrometheusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallManagedPrometheusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallManagedPrometheusResponse
	GetStatusCode() *int32
	SetBody(v *UninstallManagedPrometheusResponseBody) *UninstallManagedPrometheusResponse
	GetBody() *UninstallManagedPrometheusResponseBody
}

type UninstallManagedPrometheusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallManagedPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallManagedPrometheusResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallManagedPrometheusResponse) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallManagedPrometheusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallManagedPrometheusResponse) GetBody() *UninstallManagedPrometheusResponseBody {
	return s.Body
}

func (s *UninstallManagedPrometheusResponse) SetHeaders(v map[string]*string) *UninstallManagedPrometheusResponse {
	s.Headers = v
	return s
}

func (s *UninstallManagedPrometheusResponse) SetStatusCode(v int32) *UninstallManagedPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallManagedPrometheusResponse) SetBody(v *UninstallManagedPrometheusResponseBody) *UninstallManagedPrometheusResponse {
	s.Body = v
	return s
}

func (s *UninstallManagedPrometheusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
