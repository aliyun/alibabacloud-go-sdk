// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEndUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEndUsersResponse
	GetStatusCode() *int32
	SetBody(v *AttachEndUsersResponseBody) *AttachEndUsersResponse
	GetBody() *AttachEndUsersResponseBody
}

type AttachEndUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEndUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUsersResponse) GoString() string {
	return s.String()
}

func (s *AttachEndUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEndUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEndUsersResponse) GetBody() *AttachEndUsersResponseBody {
	return s.Body
}

func (s *AttachEndUsersResponse) SetHeaders(v map[string]*string) *AttachEndUsersResponse {
	s.Headers = v
	return s
}

func (s *AttachEndUsersResponse) SetStatusCode(v int32) *AttachEndUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEndUsersResponse) SetBody(v *AttachEndUsersResponseBody) *AttachEndUsersResponse {
	s.Body = v
	return s
}

func (s *AttachEndUsersResponse) Validate() error {
	return dara.Validate(s)
}
