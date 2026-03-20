// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBackupEncryptionAuthorizedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckBackupEncryptionAuthorizedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckBackupEncryptionAuthorizedResponse
	GetStatusCode() *int32
	SetBody(v *CheckBackupEncryptionAuthorizedResponseBody) *CheckBackupEncryptionAuthorizedResponse
	GetBody() *CheckBackupEncryptionAuthorizedResponseBody
}

type CheckBackupEncryptionAuthorizedResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBackupEncryptionAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBackupEncryptionAuthorizedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckBackupEncryptionAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *CheckBackupEncryptionAuthorizedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckBackupEncryptionAuthorizedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckBackupEncryptionAuthorizedResponse) GetBody() *CheckBackupEncryptionAuthorizedResponseBody {
	return s.Body
}

func (s *CheckBackupEncryptionAuthorizedResponse) SetHeaders(v map[string]*string) *CheckBackupEncryptionAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponse) SetStatusCode(v int32) *CheckBackupEncryptionAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponse) SetBody(v *CheckBackupEncryptionAuthorizedResponseBody) *CheckBackupEncryptionAuthorizedResponse {
	s.Body = v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
