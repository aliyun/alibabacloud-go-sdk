// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetBackupPolicyResponseBody) *SetBackupPolicyResponse
	GetBody() *SetBackupPolicyResponseBody
}

type SetBackupPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetBackupPolicyResponse) GetBody() *SetBackupPolicyResponseBody {
	return s.Body
}

func (s *SetBackupPolicyResponse) SetHeaders(v map[string]*string) *SetBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetBackupPolicyResponse) SetStatusCode(v int32) *SetBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupPolicyResponse) SetBody(v *SetBackupPolicyResponseBody) *SetBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *SetBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
