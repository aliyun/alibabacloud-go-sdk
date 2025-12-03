// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQosPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQosPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQosPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListQosPoliciesResponseBody) *ListQosPoliciesResponse
	GetBody() *ListQosPoliciesResponseBody
}

type ListQosPoliciesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQosPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQosPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQosPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListQosPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQosPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQosPoliciesResponse) GetBody() *ListQosPoliciesResponseBody {
	return s.Body
}

func (s *ListQosPoliciesResponse) SetHeaders(v map[string]*string) *ListQosPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListQosPoliciesResponse) SetStatusCode(v int32) *ListQosPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQosPoliciesResponse) SetBody(v *ListQosPoliciesResponseBody) *ListQosPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListQosPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
