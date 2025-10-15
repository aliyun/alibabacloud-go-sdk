// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConditionalAccessPoliciesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListConditionalAccessPoliciesForUserResponseBody) *ListConditionalAccessPoliciesForUserResponse
	GetBody() *ListConditionalAccessPoliciesForUserResponseBody
}

type ListConditionalAccessPoliciesForUserResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConditionalAccessPoliciesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConditionalAccessPoliciesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConditionalAccessPoliciesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConditionalAccessPoliciesForUserResponse) GetBody() *ListConditionalAccessPoliciesForUserResponseBody {
	return s.Body
}

func (s *ListConditionalAccessPoliciesForUserResponse) SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponse) SetStatusCode(v int32) *ListConditionalAccessPoliciesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponse) SetBody(v *ListConditionalAccessPoliciesForUserResponseBody) *ListConditionalAccessPoliciesForUserResponse {
	s.Body = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
