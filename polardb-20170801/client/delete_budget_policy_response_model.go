// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBudgetPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBudgetPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBudgetPolicyResponseBody) *DeleteBudgetPolicyResponse
	GetBody() *DeleteBudgetPolicyResponseBody
}

type DeleteBudgetPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBudgetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBudgetPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBudgetPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBudgetPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBudgetPolicyResponse) GetBody() *DeleteBudgetPolicyResponseBody {
	return s.Body
}

func (s *DeleteBudgetPolicyResponse) SetHeaders(v map[string]*string) *DeleteBudgetPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBudgetPolicyResponse) SetStatusCode(v int32) *DeleteBudgetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBudgetPolicyResponse) SetBody(v *DeleteBudgetPolicyResponseBody) *DeleteBudgetPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteBudgetPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
