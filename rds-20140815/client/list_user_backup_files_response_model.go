// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBackupFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserBackupFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserBackupFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserBackupFilesResponseBody) *ListUserBackupFilesResponse
	GetBody() *ListUserBackupFilesResponseBody
}

type ListUserBackupFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserBackupFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *ListUserBackupFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserBackupFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserBackupFilesResponse) GetBody() *ListUserBackupFilesResponseBody {
	return s.Body
}

func (s *ListUserBackupFilesResponse) SetHeaders(v map[string]*string) *ListUserBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *ListUserBackupFilesResponse) SetStatusCode(v int32) *ListUserBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserBackupFilesResponse) SetBody(v *ListUserBackupFilesResponseBody) *ListUserBackupFilesResponse {
	s.Body = v
	return s
}

func (s *ListUserBackupFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
