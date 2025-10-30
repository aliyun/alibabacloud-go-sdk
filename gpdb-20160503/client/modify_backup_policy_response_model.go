// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse
	GetBody() *ModifyBackupPolicyResponseBody
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupPolicyResponse) GetBody() *ModifyBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
