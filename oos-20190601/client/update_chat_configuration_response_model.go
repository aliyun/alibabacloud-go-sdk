// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatConfigurationResponseBody) *UpdateChatConfigurationResponse
	GetBody() *UpdateChatConfigurationResponseBody
}

type UpdateChatConfigurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatConfigurationResponse) GetBody() *UpdateChatConfigurationResponseBody {
	return s.Body
}

func (s *UpdateChatConfigurationResponse) SetHeaders(v map[string]*string) *UpdateChatConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatConfigurationResponse) SetStatusCode(v int32) *UpdateChatConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatConfigurationResponse) SetBody(v *UpdateChatConfigurationResponseBody) *UpdateChatConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateChatConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
