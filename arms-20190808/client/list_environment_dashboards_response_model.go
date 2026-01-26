// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentDashboardsResponseBody) *ListEnvironmentDashboardsResponse
	GetBody() *ListEnvironmentDashboardsResponseBody
}

type ListEnvironmentDashboardsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentDashboardsResponse) GetBody() *ListEnvironmentDashboardsResponseBody {
	return s.Body
}

func (s *ListEnvironmentDashboardsResponse) SetHeaders(v map[string]*string) *ListEnvironmentDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentDashboardsResponse) SetStatusCode(v int32) *ListEnvironmentDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentDashboardsResponse) SetBody(v *ListEnvironmentDashboardsResponseBody) *ListEnvironmentDashboardsResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentDashboardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
