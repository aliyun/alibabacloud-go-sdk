// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedPrometheusStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedPrometheusStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedPrometheusStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedPrometheusStatusResponseBody) *GetManagedPrometheusStatusResponse
	GetBody() *GetManagedPrometheusStatusResponseBody
}

type GetManagedPrometheusStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedPrometheusStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedPrometheusStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedPrometheusStatusResponse) GoString() string {
	return s.String()
}

func (s *GetManagedPrometheusStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedPrometheusStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedPrometheusStatusResponse) GetBody() *GetManagedPrometheusStatusResponseBody {
	return s.Body
}

func (s *GetManagedPrometheusStatusResponse) SetHeaders(v map[string]*string) *GetManagedPrometheusStatusResponse {
	s.Headers = v
	return s
}

func (s *GetManagedPrometheusStatusResponse) SetStatusCode(v int32) *GetManagedPrometheusStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedPrometheusStatusResponse) SetBody(v *GetManagedPrometheusStatusResponseBody) *GetManagedPrometheusStatusResponse {
	s.Body = v
	return s
}

func (s *GetManagedPrometheusStatusResponse) Validate() error {
	return dara.Validate(s)
}
