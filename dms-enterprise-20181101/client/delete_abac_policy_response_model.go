// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAbacPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAbacPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAbacPolicyResponseBody) *DeleteAbacPolicyResponse
	GetBody() *DeleteAbacPolicyResponseBody
}

type DeleteAbacPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAbacPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAbacPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAbacPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAbacPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAbacPolicyResponse) GetBody() *DeleteAbacPolicyResponseBody {
	return s.Body
}

func (s *DeleteAbacPolicyResponse) SetHeaders(v map[string]*string) *DeleteAbacPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAbacPolicyResponse) SetStatusCode(v int32) *DeleteAbacPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAbacPolicyResponse) SetBody(v *DeleteAbacPolicyResponseBody) *DeleteAbacPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteAbacPolicyResponse) Validate() error {
	return dara.Validate(s)
}
