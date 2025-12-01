// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIncrementBackupSetDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIncrementBackupSetDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIncrementBackupSetDownloadResponse
	GetStatusCode() *int32
	SetBody(v *CreateIncrementBackupSetDownloadResponseBody) *CreateIncrementBackupSetDownloadResponse
	GetBody() *CreateIncrementBackupSetDownloadResponseBody
}

type CreateIncrementBackupSetDownloadResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIncrementBackupSetDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIncrementBackupSetDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIncrementBackupSetDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIncrementBackupSetDownloadResponse) GetBody() *CreateIncrementBackupSetDownloadResponseBody {
	return s.Body
}

func (s *CreateIncrementBackupSetDownloadResponse) SetHeaders(v map[string]*string) *CreateIncrementBackupSetDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponse) SetStatusCode(v int32) *CreateIncrementBackupSetDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponse) SetBody(v *CreateIncrementBackupSetDownloadResponseBody) *CreateIncrementBackupSetDownloadResponse {
	s.Body = v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
