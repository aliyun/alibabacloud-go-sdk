// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAligenieUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAligenieUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAligenieUserResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAligenieUserResponseBody) *UnbindAligenieUserResponse
	GetBody() *UnbindAligenieUserResponseBody
}

type UnbindAligenieUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAligenieUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAligenieUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAligenieUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAligenieUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAligenieUserResponse) GetBody() *UnbindAligenieUserResponseBody {
	return s.Body
}

func (s *UnbindAligenieUserResponse) SetHeaders(v map[string]*string) *UnbindAligenieUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindAligenieUserResponse) SetStatusCode(v int32) *UnbindAligenieUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAligenieUserResponse) SetBody(v *UnbindAligenieUserResponseBody) *UnbindAligenieUserResponse {
	s.Body = v
	return s
}

func (s *UnbindAligenieUserResponse) Validate() error {
	return dara.Validate(s)
}
