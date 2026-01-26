// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *AddPrometheusGlobalViewResponseBody) *AddPrometheusGlobalViewResponse
	GetBody() *AddPrometheusGlobalViewResponseBody
}

type AddPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrometheusGlobalViewResponse) GetBody() *AddPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *AddPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AddPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusGlobalViewResponse) SetStatusCode(v int32) *AddPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusGlobalViewResponse) SetBody(v *AddPrometheusGlobalViewResponseBody) *AddPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *AddPrometheusGlobalViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
