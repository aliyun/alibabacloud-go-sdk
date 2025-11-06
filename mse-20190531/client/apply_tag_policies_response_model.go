// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTagPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyTagPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyTagPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ApplyTagPoliciesResponseBody) *ApplyTagPoliciesResponse
	GetBody() *ApplyTagPoliciesResponseBody
}

type ApplyTagPoliciesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyTagPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyTagPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyTagPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ApplyTagPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyTagPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyTagPoliciesResponse) GetBody() *ApplyTagPoliciesResponseBody {
	return s.Body
}

func (s *ApplyTagPoliciesResponse) SetHeaders(v map[string]*string) *ApplyTagPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ApplyTagPoliciesResponse) SetStatusCode(v int32) *ApplyTagPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyTagPoliciesResponse) SetBody(v *ApplyTagPoliciesResponseBody) *ApplyTagPoliciesResponse {
	s.Body = v
	return s
}

func (s *ApplyTagPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
