// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackupAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackupAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BackupAndroidInstanceResponseBody) *BackupAndroidInstanceResponse
	GetBody() *BackupAndroidInstanceResponseBody
}

type BackupAndroidInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackupAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackupAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BackupAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *BackupAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackupAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackupAndroidInstanceResponse) GetBody() *BackupAndroidInstanceResponseBody {
	return s.Body
}

func (s *BackupAndroidInstanceResponse) SetHeaders(v map[string]*string) *BackupAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *BackupAndroidInstanceResponse) SetStatusCode(v int32) *BackupAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BackupAndroidInstanceResponse) SetBody(v *BackupAndroidInstanceResponseBody) *BackupAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *BackupAndroidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
