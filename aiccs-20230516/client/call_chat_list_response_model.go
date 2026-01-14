// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallChatListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CallChatListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CallChatListResponse
	GetStatusCode() *int32
	SetBody(v *CallChatListResponseBody) *CallChatListResponse
	GetBody() *CallChatListResponseBody
}

type CallChatListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallChatListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallChatListResponse) String() string {
	return dara.Prettify(s)
}

func (s CallChatListResponse) GoString() string {
	return s.String()
}

func (s *CallChatListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CallChatListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CallChatListResponse) GetBody() *CallChatListResponseBody {
	return s.Body
}

func (s *CallChatListResponse) SetHeaders(v map[string]*string) *CallChatListResponse {
	s.Headers = v
	return s
}

func (s *CallChatListResponse) SetStatusCode(v int32) *CallChatListResponse {
	s.StatusCode = &v
	return s
}

func (s *CallChatListResponse) SetBody(v *CallChatListResponseBody) *CallChatListResponse {
	s.Body = v
	return s
}

func (s *CallChatListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
