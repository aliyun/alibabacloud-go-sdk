// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseBackupPolicyResponseBody) *ModifySupabaseBackupPolicyResponse
	GetBody() *ModifySupabaseBackupPolicyResponseBody
}

type ModifySupabaseBackupPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseBackupPolicyResponse) GetBody() *ModifySupabaseBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifySupabaseBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifySupabaseBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseBackupPolicyResponse) SetStatusCode(v int32) *ModifySupabaseBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseBackupPolicyResponse) SetBody(v *ModifySupabaseBackupPolicyResponseBody) *ModifySupabaseBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
