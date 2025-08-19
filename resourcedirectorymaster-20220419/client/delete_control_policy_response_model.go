// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse
	GetBody() *DeleteControlPolicyResponseBody
}

type DeleteControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteControlPolicyResponse) GetBody() *DeleteControlPolicyResponseBody {
	return s.Body
}

func (s *DeleteControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyResponse) SetStatusCode(v int32) *DeleteControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
