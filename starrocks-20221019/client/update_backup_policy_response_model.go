// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBackupPolicyResponseBody) *UpdateBackupPolicyResponse
	GetBody() *UpdateBackupPolicyResponseBody
}

type UpdateBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBackupPolicyResponse) GetBody() *UpdateBackupPolicyResponseBody {
	return s.Body
}

func (s *UpdateBackupPolicyResponse) SetHeaders(v map[string]*string) *UpdateBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupPolicyResponse) SetStatusCode(v int32) *UpdateBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupPolicyResponse) SetBody(v *UpdateBackupPolicyResponseBody) *UpdateBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
