// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatGroupResponseBody) *UpdateChatGroupResponse
	GetBody() *UpdateChatGroupResponseBody
}

type UpdateChatGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatGroupResponse) GetBody() *UpdateChatGroupResponseBody {
	return s.Body
}

func (s *UpdateChatGroupResponse) SetHeaders(v map[string]*string) *UpdateChatGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatGroupResponse) SetStatusCode(v int32) *UpdateChatGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatGroupResponse) SetBody(v *UpdateChatGroupResponseBody) *UpdateChatGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateChatGroupResponse) Validate() error {
	return dara.Validate(s)
}
