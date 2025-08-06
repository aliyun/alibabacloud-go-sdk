// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppPolicyResponseBody) *DeleteAppPolicyResponse
	GetBody() *DeleteAppPolicyResponseBody
}

type DeleteAppPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppPolicyResponse) GetBody() *DeleteAppPolicyResponseBody {
	return s.Body
}

func (s *DeleteAppPolicyResponse) SetHeaders(v map[string]*string) *DeleteAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppPolicyResponse) SetStatusCode(v int32) *DeleteAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppPolicyResponse) SetBody(v *DeleteAppPolicyResponseBody) *DeleteAppPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteAppPolicyResponse) Validate() error {
	return dara.Validate(s)
}
