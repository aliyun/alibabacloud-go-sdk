// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoliciesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoliciesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListPoliciesForUserResponseBody) *ListPoliciesForUserResponse
	GetBody() *ListPoliciesForUserResponseBody
}

type ListPoliciesForUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoliciesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoliciesForUserResponse) GetBody() *ListPoliciesForUserResponseBody {
	return s.Body
}

func (s *ListPoliciesForUserResponse) SetHeaders(v map[string]*string) *ListPoliciesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForUserResponse) SetStatusCode(v int32) *ListPoliciesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForUserResponse) SetBody(v *ListPoliciesForUserResponseBody) *ListPoliciesForUserResponse {
	s.Body = v
	return s
}

func (s *ListPoliciesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
