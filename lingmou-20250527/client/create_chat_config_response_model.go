// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatConfigResponseBody) *CreateChatConfigResponse
	GetBody() *CreateChatConfigResponseBody
}

type CreateChatConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateChatConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatConfigResponse) GetBody() *CreateChatConfigResponseBody {
	return s.Body
}

func (s *CreateChatConfigResponse) SetHeaders(v map[string]*string) *CreateChatConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateChatConfigResponse) SetStatusCode(v int32) *CreateChatConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatConfigResponse) SetBody(v *CreateChatConfigResponseBody) *CreateChatConfigResponse {
	s.Body = v
	return s
}

func (s *CreateChatConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
