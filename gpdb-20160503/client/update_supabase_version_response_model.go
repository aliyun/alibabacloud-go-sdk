// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupabaseVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSupabaseVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSupabaseVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSupabaseVersionResponseBody) *UpdateSupabaseVersionResponse
	GetBody() *UpdateSupabaseVersionResponseBody
}

type UpdateSupabaseVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSupabaseVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSupabaseVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupabaseVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSupabaseVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSupabaseVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSupabaseVersionResponse) GetBody() *UpdateSupabaseVersionResponseBody {
	return s.Body
}

func (s *UpdateSupabaseVersionResponse) SetHeaders(v map[string]*string) *UpdateSupabaseVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSupabaseVersionResponse) SetStatusCode(v int32) *UpdateSupabaseVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSupabaseVersionResponse) SetBody(v *UpdateSupabaseVersionResponseBody) *UpdateSupabaseVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateSupabaseVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
