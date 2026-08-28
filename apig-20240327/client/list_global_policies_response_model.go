// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGlobalPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGlobalPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListGlobalPoliciesResponseBody) *ListGlobalPoliciesResponse
	GetBody() *ListGlobalPoliciesResponseBody
}

type ListGlobalPoliciesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGlobalPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGlobalPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGlobalPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGlobalPoliciesResponse) GetBody() *ListGlobalPoliciesResponseBody {
	return s.Body
}

func (s *ListGlobalPoliciesResponse) SetHeaders(v map[string]*string) *ListGlobalPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalPoliciesResponse) SetStatusCode(v int32) *ListGlobalPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGlobalPoliciesResponse) SetBody(v *ListGlobalPoliciesResponseBody) *ListGlobalPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListGlobalPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
