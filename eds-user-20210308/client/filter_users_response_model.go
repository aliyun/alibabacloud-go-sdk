// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FilterUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FilterUsersResponse
	GetStatusCode() *int32
	SetBody(v *FilterUsersResponseBody) *FilterUsersResponse
	GetBody() *FilterUsersResponseBody
}

type FilterUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FilterUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FilterUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersResponse) GoString() string {
	return s.String()
}

func (s *FilterUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FilterUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FilterUsersResponse) GetBody() *FilterUsersResponseBody {
	return s.Body
}

func (s *FilterUsersResponse) SetHeaders(v map[string]*string) *FilterUsersResponse {
	s.Headers = v
	return s
}

func (s *FilterUsersResponse) SetStatusCode(v int32) *FilterUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *FilterUsersResponse) SetBody(v *FilterUsersResponseBody) *FilterUsersResponse {
	s.Body = v
	return s
}

func (s *FilterUsersResponse) Validate() error {
	return dara.Validate(s)
}
