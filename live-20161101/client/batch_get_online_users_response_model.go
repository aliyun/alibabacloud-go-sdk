// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetOnlineUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetOnlineUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetOnlineUsersResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetOnlineUsersResponseBody) *BatchGetOnlineUsersResponse
	GetBody() *BatchGetOnlineUsersResponseBody
}

type BatchGetOnlineUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetOnlineUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetOnlineUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetOnlineUsersResponse) GoString() string {
	return s.String()
}

func (s *BatchGetOnlineUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetOnlineUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetOnlineUsersResponse) GetBody() *BatchGetOnlineUsersResponseBody {
	return s.Body
}

func (s *BatchGetOnlineUsersResponse) SetHeaders(v map[string]*string) *BatchGetOnlineUsersResponse {
	s.Headers = v
	return s
}

func (s *BatchGetOnlineUsersResponse) SetStatusCode(v int32) *BatchGetOnlineUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetOnlineUsersResponse) SetBody(v *BatchGetOnlineUsersResponseBody) *BatchGetOnlineUsersResponse {
	s.Body = v
	return s
}

func (s *BatchGetOnlineUsersResponse) Validate() error {
	return dara.Validate(s)
}
