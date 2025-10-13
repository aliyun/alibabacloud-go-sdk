// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusDashboardsResponseBody) *ListPrometheusDashboardsResponse
	GetBody() *ListPrometheusDashboardsResponseBody
}

type ListPrometheusDashboardsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusDashboardsResponse) GetBody() *ListPrometheusDashboardsResponseBody {
	return s.Body
}

func (s *ListPrometheusDashboardsResponse) SetHeaders(v map[string]*string) *ListPrometheusDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusDashboardsResponse) SetStatusCode(v int32) *ListPrometheusDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusDashboardsResponse) SetBody(v *ListPrometheusDashboardsResponseBody) *ListPrometheusDashboardsResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusDashboardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
