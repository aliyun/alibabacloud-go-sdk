// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProhibitedPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProhibitedPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListProhibitedPoliciesResponseBody) *ListProhibitedPoliciesResponse
	GetBody() *ListProhibitedPoliciesResponseBody
}

type ListProhibitedPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProhibitedPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProhibitedPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProhibitedPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProhibitedPoliciesResponse) GetBody() *ListProhibitedPoliciesResponseBody {
	return s.Body
}

func (s *ListProhibitedPoliciesResponse) SetHeaders(v map[string]*string) *ListProhibitedPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListProhibitedPoliciesResponse) SetStatusCode(v int32) *ListProhibitedPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProhibitedPoliciesResponse) SetBody(v *ListProhibitedPoliciesResponseBody) *ListProhibitedPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListProhibitedPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
