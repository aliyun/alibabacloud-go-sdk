// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUniBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUniBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateUniBackupPolicyResponseBody) *CreateUniBackupPolicyResponse
	GetBody() *CreateUniBackupPolicyResponseBody
}

type CreateUniBackupPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUniBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUniBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUniBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateUniBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUniBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUniBackupPolicyResponse) GetBody() *CreateUniBackupPolicyResponseBody {
	return s.Body
}

func (s *CreateUniBackupPolicyResponse) SetHeaders(v map[string]*string) *CreateUniBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateUniBackupPolicyResponse) SetStatusCode(v int32) *CreateUniBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUniBackupPolicyResponse) SetBody(v *CreateUniBackupPolicyResponseBody) *CreateUniBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateUniBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
