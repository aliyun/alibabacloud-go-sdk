// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetBackupResponse
	GetStatusCode() *int32
	SetBody(v *ResetBackupResponseBody) *ResetBackupResponse
	GetBody() *ResetBackupResponseBody
}

type ResetBackupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetBackupResponse) GoString() string {
	return s.String()
}

func (s *ResetBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetBackupResponse) GetBody() *ResetBackupResponseBody {
	return s.Body
}

func (s *ResetBackupResponse) SetHeaders(v map[string]*string) *ResetBackupResponse {
	s.Headers = v
	return s
}

func (s *ResetBackupResponse) SetStatusCode(v int32) *ResetBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetBackupResponse) SetBody(v *ResetBackupResponseBody) *ResetBackupResponse {
	s.Body = v
	return s
}

func (s *ResetBackupResponse) Validate() error {
	return dara.Validate(s)
}
