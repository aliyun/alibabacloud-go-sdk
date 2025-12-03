// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseBackupResponse
	GetStatusCode() *int32
	SetBody(v *CloseBackupResponseBody) *CloseBackupResponse
	GetBody() *CloseBackupResponseBody
}

type CloseBackupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseBackupResponse) GoString() string {
	return s.String()
}

func (s *CloseBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseBackupResponse) GetBody() *CloseBackupResponseBody {
	return s.Body
}

func (s *CloseBackupResponse) SetHeaders(v map[string]*string) *CloseBackupResponse {
	s.Headers = v
	return s
}

func (s *CloseBackupResponse) SetStatusCode(v int32) *CloseBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseBackupResponse) SetBody(v *CloseBackupResponseBody) *CloseBackupResponse {
	s.Body = v
	return s
}

func (s *CloseBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
