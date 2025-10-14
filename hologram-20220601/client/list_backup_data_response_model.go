// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBackupDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBackupDataResponse
	GetStatusCode() *int32
	SetBody(v *ListBackupDataResponseBody) *ListBackupDataResponse
	GetBody() *ListBackupDataResponseBody
}

type ListBackupDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBackupDataResponse) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBackupDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBackupDataResponse) GetBody() *ListBackupDataResponseBody {
	return s.Body
}

func (s *ListBackupDataResponse) SetHeaders(v map[string]*string) *ListBackupDataResponse {
	s.Headers = v
	return s
}

func (s *ListBackupDataResponse) SetStatusCode(v int32) *ListBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupDataResponse) SetBody(v *ListBackupDataResponseBody) *ListBackupDataResponse {
	s.Body = v
	return s
}

func (s *ListBackupDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
