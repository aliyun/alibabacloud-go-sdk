// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAdvancedBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAdvancedBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAdvancedBackupPolicyResponseBody) *ModifyAdvancedBackupPolicyResponse
	GetBody() *ModifyAdvancedBackupPolicyResponseBody
}

type ModifyAdvancedBackupPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvancedBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvancedBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAdvancedBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAdvancedBackupPolicyResponse) GetBody() *ModifyAdvancedBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyAdvancedBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyAdvancedBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvancedBackupPolicyResponse) SetStatusCode(v int32) *ModifyAdvancedBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyResponse) SetBody(v *ModifyAdvancedBackupPolicyResponseBody) *ModifyAdvancedBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAdvancedBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
