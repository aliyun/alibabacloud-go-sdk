// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUsersResponse
	GetStatusCode() *int32
	SetBody(v *CreateUsersResponseBody) *CreateUsersResponse
	GetBody() *CreateUsersResponseBody
}

type CreateUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponse) GoString() string {
	return s.String()
}

func (s *CreateUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUsersResponse) GetBody() *CreateUsersResponseBody {
	return s.Body
}

func (s *CreateUsersResponse) SetHeaders(v map[string]*string) *CreateUsersResponse {
	s.Headers = v
	return s
}

func (s *CreateUsersResponse) SetStatusCode(v int32) *CreateUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUsersResponse) SetBody(v *CreateUsersResponseBody) *CreateUsersResponse {
	s.Body = v
	return s
}

func (s *CreateUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
