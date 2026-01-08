// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatGroupResponseBody) *DeleteChatGroupResponse
	GetBody() *DeleteChatGroupResponseBody
}

type DeleteChatGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatGroupResponse) GetBody() *DeleteChatGroupResponseBody {
	return s.Body
}

func (s *DeleteChatGroupResponse) SetHeaders(v map[string]*string) *DeleteChatGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatGroupResponse) SetStatusCode(v int32) *DeleteChatGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatGroupResponse) SetBody(v *DeleteChatGroupResponseBody) *DeleteChatGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteChatGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
