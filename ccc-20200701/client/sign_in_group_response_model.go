// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignInGroupResponse
	GetStatusCode() *int32
	SetBody(v *SignInGroupResponseBody) *SignInGroupResponse
	GetBody() *SignInGroupResponseBody
}

type SignInGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SignInGroupResponse) GoString() string {
	return s.String()
}

func (s *SignInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignInGroupResponse) GetBody() *SignInGroupResponseBody {
	return s.Body
}

func (s *SignInGroupResponse) SetHeaders(v map[string]*string) *SignInGroupResponse {
	s.Headers = v
	return s
}

func (s *SignInGroupResponse) SetStatusCode(v int32) *SignInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SignInGroupResponse) SetBody(v *SignInGroupResponseBody) *SignInGroupResponse {
	s.Body = v
	return s
}

func (s *SignInGroupResponse) Validate() error {
	return dara.Validate(s)
}
