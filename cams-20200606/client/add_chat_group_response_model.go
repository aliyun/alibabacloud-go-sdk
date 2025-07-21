// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddChatGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddChatGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddChatGroupResponseBody) *AddChatGroupResponse
	GetBody() *AddChatGroupResponseBody
}

type AddChatGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddChatGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddChatGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupResponse) GoString() string {
	return s.String()
}

func (s *AddChatGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddChatGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddChatGroupResponse) GetBody() *AddChatGroupResponseBody {
	return s.Body
}

func (s *AddChatGroupResponse) SetHeaders(v map[string]*string) *AddChatGroupResponse {
	s.Headers = v
	return s
}

func (s *AddChatGroupResponse) SetStatusCode(v int32) *AddChatGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddChatGroupResponse) SetBody(v *AddChatGroupResponseBody) *AddChatGroupResponse {
	s.Body = v
	return s
}

func (s *AddChatGroupResponse) Validate() error {
	return dara.Validate(s)
}
