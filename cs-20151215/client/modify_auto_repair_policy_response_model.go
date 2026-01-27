// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRepairPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoRepairPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoRepairPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoRepairPolicyResponseBody) *ModifyAutoRepairPolicyResponse
	GetBody() *ModifyAutoRepairPolicyResponseBody
}

type ModifyAutoRepairPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoRepairPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoRepairPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoRepairPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoRepairPolicyResponse) GetBody() *ModifyAutoRepairPolicyResponseBody {
	return s.Body
}

func (s *ModifyAutoRepairPolicyResponse) SetHeaders(v map[string]*string) *ModifyAutoRepairPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRepairPolicyResponse) SetStatusCode(v int32) *ModifyAutoRepairPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoRepairPolicyResponse) SetBody(v *ModifyAutoRepairPolicyResponseBody) *ModifyAutoRepairPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoRepairPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
