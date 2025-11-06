// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DashboardListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DashboardListResponse
	GetStatusCode() *int32
	SetBody(v *DashboardListResponseBody) *DashboardListResponse
	GetBody() *DashboardListResponseBody
}

type DashboardListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DashboardListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DashboardListResponse) String() string {
	return dara.Prettify(s)
}

func (s DashboardListResponse) GoString() string {
	return s.String()
}

func (s *DashboardListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DashboardListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DashboardListResponse) GetBody() *DashboardListResponseBody {
	return s.Body
}

func (s *DashboardListResponse) SetHeaders(v map[string]*string) *DashboardListResponse {
	s.Headers = v
	return s
}

func (s *DashboardListResponse) SetStatusCode(v int32) *DashboardListResponse {
	s.StatusCode = &v
	return s
}

func (s *DashboardListResponse) SetBody(v *DashboardListResponseBody) *DashboardListResponse {
	s.Body = v
	return s
}

func (s *DashboardListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
