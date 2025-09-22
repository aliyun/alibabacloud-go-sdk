// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetDashboardResponseBody) *GetDashboardResponse
	GetBody() *GetDashboardResponseBody
}

type GetDashboardResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDashboardResponse) GetBody() *GetDashboardResponseBody {
	return s.Body
}

func (s *GetDashboardResponse) SetHeaders(v map[string]*string) *GetDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetDashboardResponse) SetStatusCode(v int32) *GetDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDashboardResponse) SetBody(v *GetDashboardResponseBody) *GetDashboardResponse {
	s.Body = v
	return s
}

func (s *GetDashboardResponse) Validate() error {
	return dara.Validate(s)
}
