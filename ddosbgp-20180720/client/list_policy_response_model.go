// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyResponseBody) *ListPolicyResponse
	GetBody() *ListPolicyResponseBody
}

type ListPolicyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyResponse) GetBody() *ListPolicyResponseBody {
	return s.Body
}

func (s *ListPolicyResponse) SetHeaders(v map[string]*string) *ListPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyResponse) SetStatusCode(v int32) *ListPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyResponse) SetBody(v *ListPolicyResponseBody) *ListPolicyResponse {
	s.Body = v
	return s
}

func (s *ListPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
