// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUsersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse
	GetBody() *DeleteUsersResponseBody
}

type DeleteUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUsersResponse) GetBody() *DeleteUsersResponseBody {
	return s.Body
}

func (s *DeleteUsersResponse) SetHeaders(v map[string]*string) *DeleteUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsersResponse) SetStatusCode(v int32) *DeleteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

func (s *DeleteUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
