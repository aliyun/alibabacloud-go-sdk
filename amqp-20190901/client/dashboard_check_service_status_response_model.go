// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardCheckServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DashboardCheckServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DashboardCheckServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DashboardCheckServiceStatusResponseBody) *DashboardCheckServiceStatusResponse
	GetBody() *DashboardCheckServiceStatusResponseBody
}

type DashboardCheckServiceStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DashboardCheckServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DashboardCheckServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DashboardCheckServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DashboardCheckServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DashboardCheckServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DashboardCheckServiceStatusResponse) GetBody() *DashboardCheckServiceStatusResponseBody {
	return s.Body
}

func (s *DashboardCheckServiceStatusResponse) SetHeaders(v map[string]*string) *DashboardCheckServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DashboardCheckServiceStatusResponse) SetStatusCode(v int32) *DashboardCheckServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DashboardCheckServiceStatusResponse) SetBody(v *DashboardCheckServiceStatusResponseBody) *DashboardCheckServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DashboardCheckServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
