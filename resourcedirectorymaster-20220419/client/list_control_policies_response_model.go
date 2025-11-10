// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListControlPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListControlPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListControlPoliciesResponseBody) *ListControlPoliciesResponse
	GetBody() *ListControlPoliciesResponseBody
}

type ListControlPoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListControlPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListControlPoliciesResponse) GetBody() *ListControlPoliciesResponseBody {
	return s.Body
}

func (s *ListControlPoliciesResponse) SetHeaders(v map[string]*string) *ListControlPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListControlPoliciesResponse) SetStatusCode(v int32) *ListControlPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPoliciesResponse) SetBody(v *ListControlPoliciesResponseBody) *ListControlPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListControlPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
