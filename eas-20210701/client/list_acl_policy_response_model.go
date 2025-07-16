// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAclPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAclPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListAclPolicyResponseBody) *ListAclPolicyResponse
	GetBody() *ListAclPolicyResponseBody
}

type ListAclPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAclPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAclPolicyResponse) GetBody() *ListAclPolicyResponseBody {
	return s.Body
}

func (s *ListAclPolicyResponse) SetHeaders(v map[string]*string) *ListAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListAclPolicyResponse) SetStatusCode(v int32) *ListAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclPolicyResponse) SetBody(v *ListAclPolicyResponseBody) *ListAclPolicyResponse {
	s.Body = v
	return s
}

func (s *ListAclPolicyResponse) Validate() error {
	return dara.Validate(s)
}
