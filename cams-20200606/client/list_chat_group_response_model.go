// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListChatGroupResponseBody) *ListChatGroupResponse
	GetBody() *ListChatGroupResponseBody
}

type ListChatGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponse) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatGroupResponse) GetBody() *ListChatGroupResponseBody {
	return s.Body
}

func (s *ListChatGroupResponse) SetHeaders(v map[string]*string) *ListChatGroupResponse {
	s.Headers = v
	return s
}

func (s *ListChatGroupResponse) SetStatusCode(v int32) *ListChatGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatGroupResponse) SetBody(v *ListChatGroupResponseBody) *ListChatGroupResponse {
	s.Body = v
	return s
}

func (s *ListChatGroupResponse) Validate() error {
	return dara.Validate(s)
}
