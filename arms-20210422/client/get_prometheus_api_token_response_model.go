// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusApiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusApiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusApiTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusApiTokenResponseBody) *GetPrometheusApiTokenResponse
	GetBody() *GetPrometheusApiTokenResponseBody
}

type GetPrometheusApiTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusApiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusApiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusApiTokenResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusApiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusApiTokenResponse) GetBody() *GetPrometheusApiTokenResponseBody {
	return s.Body
}

func (s *GetPrometheusApiTokenResponse) SetHeaders(v map[string]*string) *GetPrometheusApiTokenResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetStatusCode(v int32) *GetPrometheusApiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetBody(v *GetPrometheusApiTokenResponseBody) *GetPrometheusApiTokenResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusApiTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
