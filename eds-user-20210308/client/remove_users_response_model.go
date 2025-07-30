// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUsersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse
	GetBody() *RemoveUsersResponseBody
}

type RemoveUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUsersResponse) GetBody() *RemoveUsersResponseBody {
	return s.Body
}

func (s *RemoveUsersResponse) SetHeaders(v map[string]*string) *RemoveUsersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersResponse) SetStatusCode(v int32) *RemoveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersResponse) SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse {
	s.Body = v
	return s
}

func (s *RemoveUsersResponse) Validate() error {
	return dara.Validate(s)
}
