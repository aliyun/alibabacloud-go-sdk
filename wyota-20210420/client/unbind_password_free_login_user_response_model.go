// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPasswordFreeLoginUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindPasswordFreeLoginUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindPasswordFreeLoginUserResponse
	GetStatusCode() *int32
	SetBody(v *UnbindPasswordFreeLoginUserResponseBody) *UnbindPasswordFreeLoginUserResponse
	GetBody() *UnbindPasswordFreeLoginUserResponseBody
}

type UnbindPasswordFreeLoginUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPasswordFreeLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPasswordFreeLoginUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindPasswordFreeLoginUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindPasswordFreeLoginUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindPasswordFreeLoginUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindPasswordFreeLoginUserResponse) GetBody() *UnbindPasswordFreeLoginUserResponseBody {
	return s.Body
}

func (s *UnbindPasswordFreeLoginUserResponse) SetHeaders(v map[string]*string) *UnbindPasswordFreeLoginUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponse) SetStatusCode(v int32) *UnbindPasswordFreeLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponse) SetBody(v *UnbindPasswordFreeLoginUserResponseBody) *UnbindPasswordFreeLoginUserResponse {
	s.Body = v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponse) Validate() error {
	return dara.Validate(s)
}
