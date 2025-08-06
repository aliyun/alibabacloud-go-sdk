// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPolicyResponseBody) *ListAppPolicyResponse
	GetBody() *ListAppPolicyResponseBody
}

type ListAppPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListAppPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPolicyResponse) GetBody() *ListAppPolicyResponseBody {
	return s.Body
}

func (s *ListAppPolicyResponse) SetHeaders(v map[string]*string) *ListAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListAppPolicyResponse) SetStatusCode(v int32) *ListAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPolicyResponse) SetBody(v *ListAppPolicyResponseBody) *ListAppPolicyResponse {
	s.Body = v
	return s
}

func (s *ListAppPolicyResponse) Validate() error {
	return dara.Validate(s)
}
