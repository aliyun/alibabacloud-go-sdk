// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListILMPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListILMPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListILMPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListILMPoliciesResponseBody) *ListILMPoliciesResponse
	GetBody() *ListILMPoliciesResponseBody
}

type ListILMPoliciesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListILMPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListILMPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListILMPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListILMPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListILMPoliciesResponse) GetBody() *ListILMPoliciesResponseBody {
	return s.Body
}

func (s *ListILMPoliciesResponse) SetHeaders(v map[string]*string) *ListILMPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListILMPoliciesResponse) SetStatusCode(v int32) *ListILMPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListILMPoliciesResponse) SetBody(v *ListILMPoliciesResponseBody) *ListILMPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListILMPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
