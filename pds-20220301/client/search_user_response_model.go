// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchUserResponse
	GetStatusCode() *int32
	SetBody(v *SearchUserResponseBody) *SearchUserResponse
	GetBody() *SearchUserResponseBody
}

type SearchUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchUserResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchUserResponse) GoString() string {
	return s.String()
}

func (s *SearchUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchUserResponse) GetBody() *SearchUserResponseBody {
	return s.Body
}

func (s *SearchUserResponse) SetHeaders(v map[string]*string) *SearchUserResponse {
	s.Headers = v
	return s
}

func (s *SearchUserResponse) SetStatusCode(v int32) *SearchUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchUserResponse) SetBody(v *SearchUserResponseBody) *SearchUserResponse {
	s.Body = v
	return s
}

func (s *SearchUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
