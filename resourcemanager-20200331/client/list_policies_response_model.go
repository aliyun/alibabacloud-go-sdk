// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse
	GetBody() *ListPoliciesResponseBody
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoliciesResponse) GetBody() *ListPoliciesResponseBody {
	return s.Body
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetStatusCode(v int32) *ListPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
