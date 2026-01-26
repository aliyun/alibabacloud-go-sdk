// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusInstanceResponseBody) *UpdatePrometheusInstanceResponse
	GetBody() *UpdatePrometheusInstanceResponseBody
}

type UpdatePrometheusInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusInstanceResponse) GetBody() *UpdatePrometheusInstanceResponseBody {
	return s.Body
}

func (s *UpdatePrometheusInstanceResponse) SetHeaders(v map[string]*string) *UpdatePrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusInstanceResponse) SetStatusCode(v int32) *UpdatePrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusInstanceResponse) SetBody(v *UpdatePrometheusInstanceResponseBody) *UpdatePrometheusInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
