// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvServiceMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvServiceMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvServiceMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvServiceMonitorsResponseBody) *ListEnvServiceMonitorsResponse
	GetBody() *ListEnvServiceMonitorsResponseBody
}

type ListEnvServiceMonitorsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvServiceMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvServiceMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvServiceMonitorsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvServiceMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvServiceMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvServiceMonitorsResponse) GetBody() *ListEnvServiceMonitorsResponseBody {
	return s.Body
}

func (s *ListEnvServiceMonitorsResponse) SetHeaders(v map[string]*string) *ListEnvServiceMonitorsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvServiceMonitorsResponse) SetStatusCode(v int32) *ListEnvServiceMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvServiceMonitorsResponse) SetBody(v *ListEnvServiceMonitorsResponseBody) *ListEnvServiceMonitorsResponse {
	s.Body = v
	return s
}

func (s *ListEnvServiceMonitorsResponse) Validate() error {
	return dara.Validate(s)
}
