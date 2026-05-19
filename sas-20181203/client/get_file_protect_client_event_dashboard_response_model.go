// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectClientEventDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectClientEventDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectClientEventDashboardResponseBody) *GetFileProtectClientEventDashboardResponse
	GetBody() *GetFileProtectClientEventDashboardResponseBody
}

type GetFileProtectClientEventDashboardResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectClientEventDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectClientEventDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectClientEventDashboardResponse) GetBody() *GetFileProtectClientEventDashboardResponseBody {
	return s.Body
}

func (s *GetFileProtectClientEventDashboardResponse) SetHeaders(v map[string]*string) *GetFileProtectClientEventDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponse) SetStatusCode(v int32) *GetFileProtectClientEventDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponse) SetBody(v *GetFileProtectClientEventDashboardResponseBody) *GetFileProtectClientEventDashboardResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
