// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceCrossBackupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceCrossBackupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceCrossBackupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceCrossBackupPolicyResponseBody) *ModifyInstanceCrossBackupPolicyResponse
	GetBody() *ModifyInstanceCrossBackupPolicyResponseBody
}

type ModifyInstanceCrossBackupPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceCrossBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceCrossBackupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceCrossBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceCrossBackupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceCrossBackupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceCrossBackupPolicyResponse) GetBody() *ModifyInstanceCrossBackupPolicyResponseBody {
	return s.Body
}

func (s *ModifyInstanceCrossBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyInstanceCrossBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponse) SetStatusCode(v int32) *ModifyInstanceCrossBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponse) SetBody(v *ModifyInstanceCrossBackupPolicyResponseBody) *ModifyInstanceCrossBackupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
