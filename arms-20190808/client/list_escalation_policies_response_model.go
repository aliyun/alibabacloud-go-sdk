// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEscalationPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEscalationPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEscalationPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListEscalationPoliciesResponseBody) *ListEscalationPoliciesResponse
	GetBody() *ListEscalationPoliciesResponseBody
}

type ListEscalationPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEscalationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEscalationPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEscalationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEscalationPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEscalationPoliciesResponse) GetBody() *ListEscalationPoliciesResponseBody {
	return s.Body
}

func (s *ListEscalationPoliciesResponse) SetHeaders(v map[string]*string) *ListEscalationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListEscalationPoliciesResponse) SetStatusCode(v int32) *ListEscalationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEscalationPoliciesResponse) SetBody(v *ListEscalationPoliciesResponseBody) *ListEscalationPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListEscalationPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
