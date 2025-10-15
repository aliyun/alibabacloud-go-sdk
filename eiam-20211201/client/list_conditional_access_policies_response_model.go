// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConditionalAccessPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListConditionalAccessPoliciesResponseBody) *ListConditionalAccessPoliciesResponse
	GetBody() *ListConditionalAccessPoliciesResponseBody
}

type ListConditionalAccessPoliciesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConditionalAccessPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConditionalAccessPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConditionalAccessPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConditionalAccessPoliciesResponse) GetBody() *ListConditionalAccessPoliciesResponseBody {
	return s.Body
}

func (s *ListConditionalAccessPoliciesResponse) SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListConditionalAccessPoliciesResponse) SetStatusCode(v int32) *ListConditionalAccessPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponse) SetBody(v *ListConditionalAccessPoliciesResponseBody) *ListConditionalAccessPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListConditionalAccessPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
