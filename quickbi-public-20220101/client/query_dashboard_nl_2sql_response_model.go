// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDashboardNl2sqlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDashboardNl2sqlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDashboardNl2sqlResponse
	GetStatusCode() *int32
	SetBody(v *QueryDashboardNl2sqlResponseBody) *QueryDashboardNl2sqlResponse
	GetBody() *QueryDashboardNl2sqlResponseBody
}

type QueryDashboardNl2sqlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDashboardNl2sqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDashboardNl2sqlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDashboardNl2sqlResponse) GoString() string {
	return s.String()
}

func (s *QueryDashboardNl2sqlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDashboardNl2sqlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDashboardNl2sqlResponse) GetBody() *QueryDashboardNl2sqlResponseBody {
	return s.Body
}

func (s *QueryDashboardNl2sqlResponse) SetHeaders(v map[string]*string) *QueryDashboardNl2sqlResponse {
	s.Headers = v
	return s
}

func (s *QueryDashboardNl2sqlResponse) SetStatusCode(v int32) *QueryDashboardNl2sqlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDashboardNl2sqlResponse) SetBody(v *QueryDashboardNl2sqlResponseBody) *QueryDashboardNl2sqlResponse {
	s.Body = v
	return s
}

func (s *QueryDashboardNl2sqlResponse) Validate() error {
	return dara.Validate(s)
}
