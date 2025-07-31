// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAuditPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAuditPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAuditPolicyResponseBody) *ModifyAuditPolicyResponse
	GetBody() *ModifyAuditPolicyResponseBody
}

type ModifyAuditPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAuditPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAuditPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAuditPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAuditPolicyResponse) GetBody() *ModifyAuditPolicyResponseBody {
	return s.Body
}

func (s *ModifyAuditPolicyResponse) SetHeaders(v map[string]*string) *ModifyAuditPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditPolicyResponse) SetStatusCode(v int32) *ModifyAuditPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAuditPolicyResponse) SetBody(v *ModifyAuditPolicyResponseBody) *ModifyAuditPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAuditPolicyResponse) Validate() error {
	return dara.Validate(s)
}
