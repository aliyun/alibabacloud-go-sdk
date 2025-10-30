// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagsForPrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagsForPrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListTagsForPrivateAccessPolicyResponseBody) *ListTagsForPrivateAccessPolicyResponse
	GetBody() *ListTagsForPrivateAccessPolicyResponseBody
}

type ListTagsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagsForPrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagsForPrivateAccessPolicyResponse) GetBody() *ListTagsForPrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListTagsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListTagsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetBody(v *ListTagsForPrivateAccessPolicyResponseBody) *ListTagsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
