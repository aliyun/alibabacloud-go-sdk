// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackupFileResponse
	GetStatusCode() *int32
	SetBody(v *BackupFileResponseBody) *BackupFileResponse
	GetBody() *BackupFileResponseBody
}

type BackupFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s BackupFileResponse) GoString() string {
	return s.String()
}

func (s *BackupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackupFileResponse) GetBody() *BackupFileResponseBody {
	return s.Body
}

func (s *BackupFileResponse) SetHeaders(v map[string]*string) *BackupFileResponse {
	s.Headers = v
	return s
}

func (s *BackupFileResponse) SetStatusCode(v int32) *BackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *BackupFileResponse) SetBody(v *BackupFileResponseBody) *BackupFileResponse {
	s.Body = v
	return s
}

func (s *BackupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
