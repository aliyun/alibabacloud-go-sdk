// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullBackupSetDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFullBackupSetDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFullBackupSetDownloadResponse
	GetStatusCode() *int32
	SetBody(v *CreateFullBackupSetDownloadResponseBody) *CreateFullBackupSetDownloadResponse
	GetBody() *CreateFullBackupSetDownloadResponseBody
}

type CreateFullBackupSetDownloadResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFullBackupSetDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFullBackupSetDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFullBackupSetDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFullBackupSetDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFullBackupSetDownloadResponse) GetBody() *CreateFullBackupSetDownloadResponseBody {
	return s.Body
}

func (s *CreateFullBackupSetDownloadResponse) SetHeaders(v map[string]*string) *CreateFullBackupSetDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateFullBackupSetDownloadResponse) SetStatusCode(v int32) *CreateFullBackupSetDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponse) SetBody(v *CreateFullBackupSetDownloadResponseBody) *CreateFullBackupSetDownloadResponse {
	s.Body = v
	return s
}

func (s *CreateFullBackupSetDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
