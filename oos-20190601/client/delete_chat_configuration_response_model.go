// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatConfigurationResponseBody) *DeleteChatConfigurationResponse
	GetBody() *DeleteChatConfigurationResponseBody
}

type DeleteChatConfigurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatConfigurationResponse) GetBody() *DeleteChatConfigurationResponseBody {
	return s.Body
}

func (s *DeleteChatConfigurationResponse) SetHeaders(v map[string]*string) *DeleteChatConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatConfigurationResponse) SetStatusCode(v int32) *DeleteChatConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatConfigurationResponse) SetBody(v *DeleteChatConfigurationResponseBody) *DeleteChatConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteChatConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
