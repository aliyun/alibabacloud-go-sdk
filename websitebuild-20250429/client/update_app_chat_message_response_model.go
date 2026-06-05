// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppChatMessageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppChatMessageResponseBody) *UpdateAppChatMessageResponse
	GetBody() *UpdateAppChatMessageResponseBody
}

type UpdateAppChatMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppChatMessageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppChatMessageResponse) GetBody() *UpdateAppChatMessageResponseBody {
	return s.Body
}

func (s *UpdateAppChatMessageResponse) SetHeaders(v map[string]*string) *UpdateAppChatMessageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppChatMessageResponse) SetStatusCode(v int32) *UpdateAppChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppChatMessageResponse) SetBody(v *UpdateAppChatMessageResponseBody) *UpdateAppChatMessageResponse {
	s.Body = v
	return s
}

func (s *UpdateAppChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
