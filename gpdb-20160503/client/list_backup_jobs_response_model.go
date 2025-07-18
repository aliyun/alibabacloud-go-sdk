// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBackupJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBackupJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListBackupJobsResponseBody) *ListBackupJobsResponse
	GetBody() *ListBackupJobsResponseBody
}

type ListBackupJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBackupJobsResponse) GoString() string {
	return s.String()
}

func (s *ListBackupJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBackupJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBackupJobsResponse) GetBody() *ListBackupJobsResponseBody {
	return s.Body
}

func (s *ListBackupJobsResponse) SetHeaders(v map[string]*string) *ListBackupJobsResponse {
	s.Headers = v
	return s
}

func (s *ListBackupJobsResponse) SetStatusCode(v int32) *ListBackupJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupJobsResponse) SetBody(v *ListBackupJobsResponseBody) *ListBackupJobsResponse {
	s.Body = v
	return s
}

func (s *ListBackupJobsResponse) Validate() error {
	return dara.Validate(s)
}
