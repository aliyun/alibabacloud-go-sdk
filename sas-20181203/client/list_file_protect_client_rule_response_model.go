// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectClientRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectClientRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectClientRuleResponseBody) *ListFileProtectClientRuleResponse
	GetBody() *ListFileProtectClientRuleResponseBody
}

type ListFileProtectClientRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectClientRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectClientRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectClientRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectClientRuleResponse) GetBody() *ListFileProtectClientRuleResponseBody {
	return s.Body
}

func (s *ListFileProtectClientRuleResponse) SetHeaders(v map[string]*string) *ListFileProtectClientRuleResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectClientRuleResponse) SetStatusCode(v int32) *ListFileProtectClientRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectClientRuleResponse) SetBody(v *ListFileProtectClientRuleResponseBody) *ListFileProtectClientRuleResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectClientRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
