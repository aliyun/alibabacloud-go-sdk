// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusViewResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusViewResponseBody) *UpdatePrometheusViewResponse
	GetBody() *UpdatePrometheusViewResponseBody
}

type UpdatePrometheusViewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusViewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusViewResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusViewResponse) GetBody() *UpdatePrometheusViewResponseBody {
	return s.Body
}

func (s *UpdatePrometheusViewResponse) SetHeaders(v map[string]*string) *UpdatePrometheusViewResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusViewResponse) SetStatusCode(v int32) *UpdatePrometheusViewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusViewResponse) SetBody(v *UpdatePrometheusViewResponseBody) *UpdatePrometheusViewResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
