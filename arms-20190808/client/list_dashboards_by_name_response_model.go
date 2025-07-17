// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDashboardsByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDashboardsByNameResponse
	GetStatusCode() *int32
	SetBody(v *ListDashboardsByNameResponseBody) *ListDashboardsByNameResponse
	GetBody() *ListDashboardsByNameResponseBody
}

type ListDashboardsByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDashboardsByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDashboardsByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsByNameResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDashboardsByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDashboardsByNameResponse) GetBody() *ListDashboardsByNameResponseBody {
	return s.Body
}

func (s *ListDashboardsByNameResponse) SetHeaders(v map[string]*string) *ListDashboardsByNameResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardsByNameResponse) SetStatusCode(v int32) *ListDashboardsByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardsByNameResponse) SetBody(v *ListDashboardsByNameResponseBody) *ListDashboardsByNameResponse {
	s.Body = v
	return s
}

func (s *ListDashboardsByNameResponse) Validate() error {
	return dara.Validate(s)
}
