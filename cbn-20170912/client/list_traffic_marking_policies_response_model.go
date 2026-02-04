// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMarkingPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficMarkingPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficMarkingPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficMarkingPoliciesResponseBody) *ListTrafficMarkingPoliciesResponse
	GetBody() *ListTrafficMarkingPoliciesResponseBody
}

type ListTrafficMarkingPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficMarkingPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficMarkingPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMarkingPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficMarkingPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficMarkingPoliciesResponse) GetBody() *ListTrafficMarkingPoliciesResponseBody {
	return s.Body
}

func (s *ListTrafficMarkingPoliciesResponse) SetHeaders(v map[string]*string) *ListTrafficMarkingPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponse) SetStatusCode(v int32) *ListTrafficMarkingPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficMarkingPoliciesResponse) SetBody(v *ListTrafficMarkingPoliciesResponseBody) *ListTrafficMarkingPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListTrafficMarkingPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
