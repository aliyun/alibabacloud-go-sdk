// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockUserResponse
	GetStatusCode() *int32
	SetBody(v *UnlockUserResponseBody) *UnlockUserResponse
	GetBody() *UnlockUserResponseBody
}

type UnlockUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockUserResponse) GoString() string {
	return s.String()
}

func (s *UnlockUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockUserResponse) GetBody() *UnlockUserResponseBody {
	return s.Body
}

func (s *UnlockUserResponse) SetHeaders(v map[string]*string) *UnlockUserResponse {
	s.Headers = v
	return s
}

func (s *UnlockUserResponse) SetStatusCode(v int32) *UnlockUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUserResponse) SetBody(v *UnlockUserResponseBody) *UnlockUserResponse {
	s.Body = v
	return s
}

func (s *UnlockUserResponse) Validate() error {
	return dara.Validate(s)
}
