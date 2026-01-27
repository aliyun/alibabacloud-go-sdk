// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoRepairPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoRepairPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoRepairPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoRepairPoliciesResponseBody) *ListAutoRepairPoliciesResponse
	GetBody() *ListAutoRepairPoliciesResponseBody
}

type ListAutoRepairPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoRepairPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoRepairPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoRepairPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoRepairPoliciesResponse) GetBody() *ListAutoRepairPoliciesResponseBody {
	return s.Body
}

func (s *ListAutoRepairPoliciesResponse) SetHeaders(v map[string]*string) *ListAutoRepairPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoRepairPoliciesResponse) SetStatusCode(v int32) *ListAutoRepairPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoRepairPoliciesResponse) SetBody(v *ListAutoRepairPoliciesResponseBody) *ListAutoRepairPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListAutoRepairPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
