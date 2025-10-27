// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectDashboardResponseBody) *GetFileProtectDashboardResponse
	GetBody() *GetFileProtectDashboardResponseBody
}

type GetFileProtectDashboardResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectDashboardResponse) GetBody() *GetFileProtectDashboardResponseBody {
	return s.Body
}

func (s *GetFileProtectDashboardResponse) SetHeaders(v map[string]*string) *GetFileProtectDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectDashboardResponse) SetStatusCode(v int32) *GetFileProtectDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectDashboardResponse) SetBody(v *GetFileProtectDashboardResponseBody) *GetFileProtectDashboardResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectDashboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
