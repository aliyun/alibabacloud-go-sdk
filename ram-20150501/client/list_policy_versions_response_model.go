// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyVersionsResponseBody) *ListPolicyVersionsResponse
	GetBody() *ListPolicyVersionsResponseBody
}

type ListPolicyVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyVersionsResponse) GetBody() *ListPolicyVersionsResponseBody {
	return s.Body
}

func (s *ListPolicyVersionsResponse) SetHeaders(v map[string]*string) *ListPolicyVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyVersionsResponse) SetStatusCode(v int32) *ListPolicyVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyVersionsResponse) SetBody(v *ListPolicyVersionsResponseBody) *ListPolicyVersionsResponse {
	s.Body = v
	return s
}

func (s *ListPolicyVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
