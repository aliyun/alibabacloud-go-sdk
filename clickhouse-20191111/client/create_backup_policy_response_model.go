// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackupPolicyResponseBody) *CreateBackupPolicyResponse
	GetBody() *CreateBackupPolicyResponseBody
}

type CreateBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackupPolicyResponse) GetBody() *CreateBackupPolicyResponseBody {
	return s.Body
}

func (s *CreateBackupPolicyResponse) SetHeaders(v map[string]*string) *CreateBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPolicyResponse) SetStatusCode(v int32) *CreateBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupPolicyResponse) SetBody(v *CreateBackupPolicyResponseBody) *CreateBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
