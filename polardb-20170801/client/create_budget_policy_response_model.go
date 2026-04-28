// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBudgetPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBudgetPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateBudgetPolicyResponseBody) *CreateBudgetPolicyResponse
	GetBody() *CreateBudgetPolicyResponseBody
}

type CreateBudgetPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBudgetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBudgetPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateBudgetPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBudgetPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBudgetPolicyResponse) GetBody() *CreateBudgetPolicyResponseBody {
	return s.Body
}

func (s *CreateBudgetPolicyResponse) SetHeaders(v map[string]*string) *CreateBudgetPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateBudgetPolicyResponse) SetStatusCode(v int32) *CreateBudgetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBudgetPolicyResponse) SetBody(v *CreateBudgetPolicyResponseBody) *CreateBudgetPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateBudgetPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
