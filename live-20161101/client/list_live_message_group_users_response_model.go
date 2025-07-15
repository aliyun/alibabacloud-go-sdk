// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveMessageGroupUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveMessageGroupUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveMessageGroupUsersResponseBody) *ListLiveMessageGroupUsersResponse
	GetBody() *ListLiveMessageGroupUsersResponseBody
}

type ListLiveMessageGroupUsersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveMessageGroupUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveMessageGroupUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupUsersResponse) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveMessageGroupUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveMessageGroupUsersResponse) GetBody() *ListLiveMessageGroupUsersResponseBody {
	return s.Body
}

func (s *ListLiveMessageGroupUsersResponse) SetHeaders(v map[string]*string) *ListLiveMessageGroupUsersResponse {
	s.Headers = v
	return s
}

func (s *ListLiveMessageGroupUsersResponse) SetStatusCode(v int32) *ListLiveMessageGroupUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveMessageGroupUsersResponse) SetBody(v *ListLiveMessageGroupUsersResponseBody) *ListLiveMessageGroupUsersResponse {
	s.Body = v
	return s
}

func (s *ListLiveMessageGroupUsersResponse) Validate() error {
	return dara.Validate(s)
}
