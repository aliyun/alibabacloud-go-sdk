// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSilencePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSilencePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSilencePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListSilencePoliciesResponseBody) *ListSilencePoliciesResponse
	GetBody() *ListSilencePoliciesResponseBody
}

type ListSilencePoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSilencePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSilencePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSilencePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSilencePoliciesResponse) GetBody() *ListSilencePoliciesResponseBody {
	return s.Body
}

func (s *ListSilencePoliciesResponse) SetHeaders(v map[string]*string) *ListSilencePoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSilencePoliciesResponse) SetStatusCode(v int32) *ListSilencePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSilencePoliciesResponse) SetBody(v *ListSilencePoliciesResponseBody) *ListSilencePoliciesResponse {
	s.Body = v
	return s
}

func (s *ListSilencePoliciesResponse) Validate() error {
	return dara.Validate(s)
}
