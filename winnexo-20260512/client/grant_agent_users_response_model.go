// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAgentUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantAgentUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantAgentUsersResponse
	GetStatusCode() *int32
	SetBody(v *GrantAgentUsersResponseBody) *GrantAgentUsersResponse
	GetBody() *GrantAgentUsersResponseBody
}

type GrantAgentUsersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantAgentUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantAgentUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantAgentUsersResponse) GoString() string {
	return s.String()
}

func (s *GrantAgentUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantAgentUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantAgentUsersResponse) GetBody() *GrantAgentUsersResponseBody {
	return s.Body
}

func (s *GrantAgentUsersResponse) SetHeaders(v map[string]*string) *GrantAgentUsersResponse {
	s.Headers = v
	return s
}

func (s *GrantAgentUsersResponse) SetStatusCode(v int32) *GrantAgentUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantAgentUsersResponse) SetBody(v *GrantAgentUsersResponseBody) *GrantAgentUsersResponse {
	s.Body = v
	return s
}

func (s *GrantAgentUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
