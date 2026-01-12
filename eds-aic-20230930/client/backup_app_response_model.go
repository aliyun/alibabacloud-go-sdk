// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackupAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackupAppResponse
	GetStatusCode() *int32
	SetBody(v *BackupAppResponseBody) *BackupAppResponse
	GetBody() *BackupAppResponseBody
}

type BackupAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackupAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackupAppResponse) String() string {
	return dara.Prettify(s)
}

func (s BackupAppResponse) GoString() string {
	return s.String()
}

func (s *BackupAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackupAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackupAppResponse) GetBody() *BackupAppResponseBody {
	return s.Body
}

func (s *BackupAppResponse) SetHeaders(v map[string]*string) *BackupAppResponse {
	s.Headers = v
	return s
}

func (s *BackupAppResponse) SetStatusCode(v int32) *BackupAppResponse {
	s.StatusCode = &v
	return s
}

func (s *BackupAppResponse) SetBody(v *BackupAppResponseBody) *BackupAppResponse {
	s.Body = v
	return s
}

func (s *BackupAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
