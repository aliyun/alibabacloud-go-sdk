// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyClassesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyClassesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyClassesResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyClassesResponseBody) *ListPolicyClassesResponse
	GetBody() *ListPolicyClassesResponseBody
}

type ListPolicyClassesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyClassesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyClassesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyClassesResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyClassesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyClassesResponse) GetBody() *ListPolicyClassesResponseBody {
	return s.Body
}

func (s *ListPolicyClassesResponse) SetHeaders(v map[string]*string) *ListPolicyClassesResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyClassesResponse) SetStatusCode(v int32) *ListPolicyClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyClassesResponse) SetBody(v *ListPolicyClassesResponseBody) *ListPolicyClassesResponse {
	s.Body = v
	return s
}

func (s *ListPolicyClassesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
