// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupPolicyResponseBody) *DeleteBackupPolicyResponse
	GetBody() *DeleteBackupPolicyResponseBody
}

type DeleteBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupPolicyResponse) GetBody() *DeleteBackupPolicyResponseBody {
	return s.Body
}

func (s *DeleteBackupPolicyResponse) SetHeaders(v map[string]*string) *DeleteBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupPolicyResponse) SetStatusCode(v int32) *DeleteBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupPolicyResponse) SetBody(v *DeleteBackupPolicyResponseBody) *DeleteBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
