// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLogBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLogBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLogBackupPolicyResponseBody) *ModifyLogBackupPolicyResponse
	GetBody() *ModifyLogBackupPolicyResponseBody
}

type ModifyLogBackupPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLogBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLogBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLogBackupPolicyResponse) GetBody() *ModifyLogBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyLogBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetStatusCode(v int32) *ModifyLogBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetBody(v *ModifyLogBackupPolicyResponseBody) *ModifyLogBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyLogBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
