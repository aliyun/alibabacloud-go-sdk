// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAgentUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAgentUsersResponse
	GetStatusCode() *int32
	SetBody(v *RevokeAgentUsersResponseBody) *RevokeAgentUsersResponse
	GetBody() *RevokeAgentUsersResponseBody
}

type RevokeAgentUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeAgentUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeAgentUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAgentUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeAgentUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAgentUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAgentUsersResponse) GetBody() *RevokeAgentUsersResponseBody {
	return s.Body
}

func (s *RevokeAgentUsersResponse) SetHeaders(v map[string]*string) *RevokeAgentUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeAgentUsersResponse) SetStatusCode(v int32) *RevokeAgentUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAgentUsersResponse) SetBody(v *RevokeAgentUsersResponseBody) *RevokeAgentUsersResponse {
	s.Body = v
	return s
}

func (s *RevokeAgentUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
