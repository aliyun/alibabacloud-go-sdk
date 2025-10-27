// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUniBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUniBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUniBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUniBackupPolicyResponseBody) *ModifyUniBackupPolicyResponse
	GetBody() *ModifyUniBackupPolicyResponseBody
}

type ModifyUniBackupPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUniBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUniBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUniBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyUniBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUniBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUniBackupPolicyResponse) GetBody() *ModifyUniBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyUniBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyUniBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyUniBackupPolicyResponse) SetStatusCode(v int32) *ModifyUniBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUniBackupPolicyResponse) SetBody(v *ModifyUniBackupPolicyResponseBody) *ModifyUniBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyUniBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
