// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse
	GetBody() *DeletePolicyResponseBody
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyResponse) GetBody() *DeletePolicyResponseBody {
	return s.Body
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyResponse) Validate() error {
	return dara.Validate(s)
}
