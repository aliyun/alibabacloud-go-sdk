// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPolicyRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityPolicyRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityPolicyRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityPolicyRelationsResponseBody) *ListSecurityPolicyRelationsResponse
	GetBody() *ListSecurityPolicyRelationsResponseBody
}

type ListSecurityPolicyRelationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityPolicyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityPolicyRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityPolicyRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityPolicyRelationsResponse) GetBody() *ListSecurityPolicyRelationsResponseBody {
	return s.Body
}

func (s *ListSecurityPolicyRelationsResponse) SetHeaders(v map[string]*string) *ListSecurityPolicyRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) SetStatusCode(v int32) *ListSecurityPolicyRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) SetBody(v *ListSecurityPolicyRelationsResponseBody) *ListSecurityPolicyRelationsResponse {
	s.Body = v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
