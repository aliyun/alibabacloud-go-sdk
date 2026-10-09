// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseBackupJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupabaseBackupJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupabaseBackupJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupabaseBackupJobsResponseBody) *ListSupabaseBackupJobsResponse
	GetBody() *ListSupabaseBackupJobsResponseBody
}

type ListSupabaseBackupJobsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupabaseBackupJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupabaseBackupJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseBackupJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSupabaseBackupJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupabaseBackupJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupabaseBackupJobsResponse) GetBody() *ListSupabaseBackupJobsResponseBody {
	return s.Body
}

func (s *ListSupabaseBackupJobsResponse) SetHeaders(v map[string]*string) *ListSupabaseBackupJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSupabaseBackupJobsResponse) SetStatusCode(v int32) *ListSupabaseBackupJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupabaseBackupJobsResponse) SetBody(v *ListSupabaseBackupJobsResponseBody) *ListSupabaseBackupJobsResponse {
	s.Body = v
	return s
}

func (s *ListSupabaseBackupJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
