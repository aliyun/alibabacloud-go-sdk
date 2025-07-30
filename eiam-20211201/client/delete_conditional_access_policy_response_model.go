// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConditionalAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConditionalAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConditionalAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConditionalAccessPolicyResponseBody) *DeleteConditionalAccessPolicyResponse
	GetBody() *DeleteConditionalAccessPolicyResponseBody
}

type DeleteConditionalAccessPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConditionalAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConditionalAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteConditionalAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConditionalAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConditionalAccessPolicyResponse) GetBody() *DeleteConditionalAccessPolicyResponseBody {
	return s.Body
}

func (s *DeleteConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *DeleteConditionalAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteConditionalAccessPolicyResponse) SetStatusCode(v int32) *DeleteConditionalAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConditionalAccessPolicyResponse) SetBody(v *DeleteConditionalAccessPolicyResponseBody) *DeleteConditionalAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteConditionalAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
