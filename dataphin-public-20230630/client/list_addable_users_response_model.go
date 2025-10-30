// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddableUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddableUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListAddableUsersResponseBody) *ListAddableUsersResponse
	GetBody() *ListAddableUsersResponseBody
}

type ListAddableUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddableUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddableUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersResponse) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddableUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddableUsersResponse) GetBody() *ListAddableUsersResponseBody {
	return s.Body
}

func (s *ListAddableUsersResponse) SetHeaders(v map[string]*string) *ListAddableUsersResponse {
	s.Headers = v
	return s
}

func (s *ListAddableUsersResponse) SetStatusCode(v int32) *ListAddableUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddableUsersResponse) SetBody(v *ListAddableUsersResponseBody) *ListAddableUsersResponse {
	s.Body = v
	return s
}

func (s *ListAddableUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
