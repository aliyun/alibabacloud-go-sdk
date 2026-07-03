// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseDashboardPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseDashboardPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseDashboardPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseDashboardPasswordResponseBody) *ModifySupabaseDashboardPasswordResponse
	GetBody() *ModifySupabaseDashboardPasswordResponseBody
}

type ModifySupabaseDashboardPasswordResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseDashboardPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseDashboardPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseDashboardPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseDashboardPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseDashboardPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseDashboardPasswordResponse) GetBody() *ModifySupabaseDashboardPasswordResponseBody {
	return s.Body
}

func (s *ModifySupabaseDashboardPasswordResponse) SetHeaders(v map[string]*string) *ModifySupabaseDashboardPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponse) SetStatusCode(v int32) *ModifySupabaseDashboardPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponse) SetBody(v *ModifySupabaseDashboardPasswordResponseBody) *ModifySupabaseDashboardPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
