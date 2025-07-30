// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeApplicationFromUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeApplicationFromUsersResponse
	GetStatusCode() *int32
	SetBody(v *RevokeApplicationFromUsersResponseBody) *RevokeApplicationFromUsersResponse
	GetBody() *RevokeApplicationFromUsersResponseBody
}

type RevokeApplicationFromUsersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeApplicationFromUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeApplicationFromUsersResponse) GetBody() *RevokeApplicationFromUsersResponseBody {
	return s.Body
}

func (s *RevokeApplicationFromUsersResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromUsersResponse) SetStatusCode(v int32) *RevokeApplicationFromUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromUsersResponse) SetBody(v *RevokeApplicationFromUsersResponseBody) *RevokeApplicationFromUsersResponse {
	s.Body = v
	return s
}

func (s *RevokeApplicationFromUsersResponse) Validate() error {
	return dara.Validate(s)
}
