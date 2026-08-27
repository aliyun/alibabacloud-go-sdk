// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWithGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserWithGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserWithGroupsResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserWithGroupsResponseBody) *CreateUserWithGroupsResponse
	GetBody() *CreateUserWithGroupsResponseBody
}

type CreateUserWithGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserWithGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserWithGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWithGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateUserWithGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserWithGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserWithGroupsResponse) GetBody() *CreateUserWithGroupsResponseBody {
	return s.Body
}

func (s *CreateUserWithGroupsResponse) SetHeaders(v map[string]*string) *CreateUserWithGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateUserWithGroupsResponse) SetStatusCode(v int32) *CreateUserWithGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserWithGroupsResponse) SetBody(v *CreateUserWithGroupsResponseBody) *CreateUserWithGroupsResponse {
	s.Body = v
	return s
}

func (s *CreateUserWithGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
