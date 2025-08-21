// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteILMPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteILMPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteILMPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteILMPolicyResponseBody) *DeleteILMPolicyResponse
	GetBody() *DeleteILMPolicyResponseBody
}

type DeleteILMPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteILMPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteILMPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteILMPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteILMPolicyResponse) GetBody() *DeleteILMPolicyResponseBody {
	return s.Body
}

func (s *DeleteILMPolicyResponse) SetHeaders(v map[string]*string) *DeleteILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteILMPolicyResponse) SetStatusCode(v int32) *DeleteILMPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteILMPolicyResponse) SetBody(v *DeleteILMPolicyResponseBody) *DeleteILMPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteILMPolicyResponse) Validate() error {
	return dara.Validate(s)
}
