// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSupabaseProjectPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSupabaseProjectPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSupabaseProjectPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetSupabaseProjectPasswordResponseBody) *ResetSupabaseProjectPasswordResponse
	GetBody() *ResetSupabaseProjectPasswordResponseBody
}

type ResetSupabaseProjectPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSupabaseProjectPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSupabaseProjectPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSupabaseProjectPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetSupabaseProjectPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSupabaseProjectPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSupabaseProjectPasswordResponse) GetBody() *ResetSupabaseProjectPasswordResponseBody {
	return s.Body
}

func (s *ResetSupabaseProjectPasswordResponse) SetHeaders(v map[string]*string) *ResetSupabaseProjectPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetSupabaseProjectPasswordResponse) SetStatusCode(v int32) *ResetSupabaseProjectPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSupabaseProjectPasswordResponse) SetBody(v *ResetSupabaseProjectPasswordResponseBody) *ResetSupabaseProjectPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetSupabaseProjectPasswordResponse) Validate() error {
	return dara.Validate(s)
}
