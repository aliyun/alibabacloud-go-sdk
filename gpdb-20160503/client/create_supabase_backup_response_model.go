// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSupabaseBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSupabaseBackupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSupabaseBackupResponseBody) *CreateSupabaseBackupResponse
	GetBody() *CreateSupabaseBackupResponseBody
}

type CreateSupabaseBackupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSupabaseBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSupabaseBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateSupabaseBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSupabaseBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSupabaseBackupResponse) GetBody() *CreateSupabaseBackupResponseBody {
	return s.Body
}

func (s *CreateSupabaseBackupResponse) SetHeaders(v map[string]*string) *CreateSupabaseBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateSupabaseBackupResponse) SetStatusCode(v int32) *CreateSupabaseBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSupabaseBackupResponse) SetBody(v *CreateSupabaseBackupResponseBody) *CreateSupabaseBackupResponse {
	s.Body = v
	return s
}

func (s *CreateSupabaseBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
