// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUniBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUniBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUniBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUniBackupPolicyResponseBody) *DeleteUniBackupPolicyResponse
	GetBody() *DeleteUniBackupPolicyResponseBody
}

type DeleteUniBackupPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUniBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUniBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUniBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteUniBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUniBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUniBackupPolicyResponse) GetBody() *DeleteUniBackupPolicyResponseBody {
	return s.Body
}

func (s *DeleteUniBackupPolicyResponse) SetHeaders(v map[string]*string) *DeleteUniBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteUniBackupPolicyResponse) SetStatusCode(v int32) *DeleteUniBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUniBackupPolicyResponse) SetBody(v *DeleteUniBackupPolicyResponseBody) *DeleteUniBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteUniBackupPolicyResponse) Validate() error {
	return dara.Validate(s)
}
