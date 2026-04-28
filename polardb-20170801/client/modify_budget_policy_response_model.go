// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBudgetPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBudgetPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBudgetPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBudgetPolicyResponseBody) *ModifyBudgetPolicyResponse
	GetBody() *ModifyBudgetPolicyResponseBody
}

type ModifyBudgetPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBudgetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBudgetPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBudgetPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBudgetPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBudgetPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBudgetPolicyResponse) GetBody() *ModifyBudgetPolicyResponseBody {
	return s.Body
}

func (s *ModifyBudgetPolicyResponse) SetHeaders(v map[string]*string) *ModifyBudgetPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBudgetPolicyResponse) SetStatusCode(v int32) *ModifyBudgetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBudgetPolicyResponse) SetBody(v *ModifyBudgetPolicyResponseBody) *ModifyBudgetPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyBudgetPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
