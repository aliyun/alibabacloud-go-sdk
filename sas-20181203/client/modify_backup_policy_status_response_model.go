// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupPolicyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupPolicyStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupPolicyStatusResponseBody) *ModifyBackupPolicyStatusResponse
	GetBody() *ModifyBackupPolicyStatusResponseBody
}

type ModifyBackupPolicyStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPolicyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPolicyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupPolicyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupPolicyStatusResponse) GetBody() *ModifyBackupPolicyStatusResponseBody {
	return s.Body
}

func (s *ModifyBackupPolicyStatusResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyStatusResponse) SetStatusCode(v int32) *ModifyBackupPolicyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyStatusResponse) SetBody(v *ModifyBackupPolicyStatusResponseBody) *ModifyBackupPolicyStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupPolicyStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
