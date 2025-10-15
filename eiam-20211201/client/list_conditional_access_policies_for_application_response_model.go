// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConditionalAccessPoliciesForApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListConditionalAccessPoliciesForApplicationResponseBody) *ListConditionalAccessPoliciesForApplicationResponse
	GetBody() *ListConditionalAccessPoliciesForApplicationResponseBody
}

type ListConditionalAccessPoliciesForApplicationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConditionalAccessPoliciesForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConditionalAccessPoliciesForApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) GetBody() *ListConditionalAccessPoliciesForApplicationResponseBody {
	return s.Body
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) SetStatusCode(v int32) *ListConditionalAccessPoliciesForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) SetBody(v *ListConditionalAccessPoliciesForApplicationResponseBody) *ListConditionalAccessPoliciesForApplicationResponse {
	s.Body = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
