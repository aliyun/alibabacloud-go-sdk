// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *ListDashboardsResponseBody) *ListDashboardsResponse
	GetBody() *ListDashboardsResponseBody
}

type ListDashboardsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDashboardsResponse) GetBody() *ListDashboardsResponseBody {
	return s.Body
}

func (s *ListDashboardsResponse) SetHeaders(v map[string]*string) *ListDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardsResponse) SetStatusCode(v int32) *ListDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardsResponse) SetBody(v *ListDashboardsResponseBody) *ListDashboardsResponse {
	s.Body = v
	return s
}

func (s *ListDashboardsResponse) Validate() error {
	return dara.Validate(s)
}
