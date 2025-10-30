// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectDashboardAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupabaseProjectDashboardAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupabaseProjectDashboardAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetSupabaseProjectDashboardAccountResponseBody) *GetSupabaseProjectDashboardAccountResponse
	GetBody() *GetSupabaseProjectDashboardAccountResponseBody
}

type GetSupabaseProjectDashboardAccountResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupabaseProjectDashboardAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupabaseProjectDashboardAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectDashboardAccountResponse) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectDashboardAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupabaseProjectDashboardAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupabaseProjectDashboardAccountResponse) GetBody() *GetSupabaseProjectDashboardAccountResponseBody {
	return s.Body
}

func (s *GetSupabaseProjectDashboardAccountResponse) SetHeaders(v map[string]*string) *GetSupabaseProjectDashboardAccountResponse {
	s.Headers = v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponse) SetStatusCode(v int32) *GetSupabaseProjectDashboardAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponse) SetBody(v *GetSupabaseProjectDashboardAccountResponseBody) *GetSupabaseProjectDashboardAccountResponse {
	s.Body = v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
