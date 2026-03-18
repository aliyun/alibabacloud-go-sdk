// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddBackupPolicyResponseBody) *AddBackupPolicyResponse
	GetBody() *AddBackupPolicyResponseBody
}

type AddBackupPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBackupPolicyResponse) GetBody() *AddBackupPolicyResponseBody {
	return s.Body
}

func (s *AddBackupPolicyResponse) SetHeaders(v map[string]*string) *AddBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddBackupPolicyResponse) SetStatusCode(v int32) *AddBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBackupPolicyResponse) SetBody(v *AddBackupPolicyResponseBody) *AddBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *AddBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
