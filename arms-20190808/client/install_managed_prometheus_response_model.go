// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallManagedPrometheusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallManagedPrometheusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallManagedPrometheusResponse
	GetStatusCode() *int32
	SetBody(v *InstallManagedPrometheusResponseBody) *InstallManagedPrometheusResponse
	GetBody() *InstallManagedPrometheusResponseBody
}

type InstallManagedPrometheusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallManagedPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallManagedPrometheusResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallManagedPrometheusResponse) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallManagedPrometheusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallManagedPrometheusResponse) GetBody() *InstallManagedPrometheusResponseBody {
	return s.Body
}

func (s *InstallManagedPrometheusResponse) SetHeaders(v map[string]*string) *InstallManagedPrometheusResponse {
	s.Headers = v
	return s
}

func (s *InstallManagedPrometheusResponse) SetStatusCode(v int32) *InstallManagedPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallManagedPrometheusResponse) SetBody(v *InstallManagedPrometheusResponseBody) *InstallManagedPrometheusResponse {
	s.Body = v
	return s
}

func (s *InstallManagedPrometheusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
