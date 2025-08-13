// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsForPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTargetsForPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTargetsForPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListTargetsForPolicyResponseBody) *ListTargetsForPolicyResponse
	GetBody() *ListTargetsForPolicyResponseBody
}

type ListTargetsForPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetsForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetsForPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsForPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTargetsForPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTargetsForPolicyResponse) GetBody() *ListTargetsForPolicyResponseBody {
	return s.Body
}

func (s *ListTargetsForPolicyResponse) SetHeaders(v map[string]*string) *ListTargetsForPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTargetsForPolicyResponse) SetStatusCode(v int32) *ListTargetsForPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetsForPolicyResponse) SetBody(v *ListTargetsForPolicyResponseBody) *ListTargetsForPolicyResponse {
	s.Body = v
	return s
}

func (s *ListTargetsForPolicyResponse) Validate() error {
	return dara.Validate(s)
}
