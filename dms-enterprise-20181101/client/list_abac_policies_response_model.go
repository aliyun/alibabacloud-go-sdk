// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAbacPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAbacPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListAbacPoliciesResponseBody) *ListAbacPoliciesResponse
	GetBody() *ListAbacPoliciesResponseBody
}

type ListAbacPoliciesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAbacPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAbacPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAbacPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListAbacPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAbacPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAbacPoliciesResponse) GetBody() *ListAbacPoliciesResponseBody {
	return s.Body
}

func (s *ListAbacPoliciesResponse) SetHeaders(v map[string]*string) *ListAbacPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListAbacPoliciesResponse) SetStatusCode(v int32) *ListAbacPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAbacPoliciesResponse) SetBody(v *ListAbacPoliciesResponseBody) *ListAbacPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListAbacPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
