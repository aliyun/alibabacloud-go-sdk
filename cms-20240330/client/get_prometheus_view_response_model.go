// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusViewResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusViewResponseBody) *GetPrometheusViewResponse
	GetBody() *GetPrometheusViewResponseBody
}

type GetPrometheusViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusViewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusViewResponse) GetBody() *GetPrometheusViewResponseBody {
	return s.Body
}

func (s *GetPrometheusViewResponse) SetHeaders(v map[string]*string) *GetPrometheusViewResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusViewResponse) SetStatusCode(v int32) *GetPrometheusViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusViewResponse) SetBody(v *GetPrometheusViewResponseBody) *GetPrometheusViewResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
