// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRayDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRayDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetRayDashboardResponseBody) *GetRayDashboardResponse
	GetBody() *GetRayDashboardResponseBody
}

type GetRayDashboardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRayDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetRayDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRayDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRayDashboardResponse) GetBody() *GetRayDashboardResponseBody {
	return s.Body
}

func (s *GetRayDashboardResponse) SetHeaders(v map[string]*string) *GetRayDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetRayDashboardResponse) SetStatusCode(v int32) *GetRayDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayDashboardResponse) SetBody(v *GetRayDashboardResponseBody) *GetRayDashboardResponse {
	s.Body = v
	return s
}

func (s *GetRayDashboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
