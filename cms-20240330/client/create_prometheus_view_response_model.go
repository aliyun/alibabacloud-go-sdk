// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrometheusViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrometheusViewResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrometheusViewResponseBody) *CreatePrometheusViewResponse
	GetBody() *CreatePrometheusViewResponseBody
}

type CreatePrometheusViewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusViewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusViewResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrometheusViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrometheusViewResponse) GetBody() *CreatePrometheusViewResponseBody {
	return s.Body
}

func (s *CreatePrometheusViewResponse) SetHeaders(v map[string]*string) *CreatePrometheusViewResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusViewResponse) SetStatusCode(v int32) *CreatePrometheusViewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusViewResponse) SetBody(v *CreatePrometheusViewResponseBody) *CreatePrometheusViewResponse {
	s.Body = v
	return s
}

func (s *CreatePrometheusViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
