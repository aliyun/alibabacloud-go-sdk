// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationBackupPolicyResponseBody) *ModifyApplicationBackupPolicyResponse
	GetBody() *ModifyApplicationBackupPolicyResponseBody
}

type ModifyApplicationBackupPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationBackupPolicyResponse) GetBody() *ModifyApplicationBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyApplicationBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyApplicationBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationBackupPolicyResponse) SetStatusCode(v int32) *ModifyApplicationBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationBackupPolicyResponse) SetBody(v *ModifyApplicationBackupPolicyResponseBody) *ModifyApplicationBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
