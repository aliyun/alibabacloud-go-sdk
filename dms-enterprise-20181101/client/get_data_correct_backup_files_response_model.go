// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectBackupFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCorrectBackupFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCorrectBackupFilesResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCorrectBackupFilesResponseBody) *GetDataCorrectBackupFilesResponse
	GetBody() *GetDataCorrectBackupFilesResponseBody
}

type GetDataCorrectBackupFilesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCorrectBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCorrectBackupFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCorrectBackupFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCorrectBackupFilesResponse) GetBody() *GetDataCorrectBackupFilesResponseBody {
	return s.Body
}

func (s *GetDataCorrectBackupFilesResponse) SetHeaders(v map[string]*string) *GetDataCorrectBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectBackupFilesResponse) SetStatusCode(v int32) *GetDataCorrectBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponse) SetBody(v *GetDataCorrectBackupFilesResponseBody) *GetDataCorrectBackupFilesResponse {
	s.Body = v
	return s
}

func (s *GetDataCorrectBackupFilesResponse) Validate() error {
	return dara.Validate(s)
}
