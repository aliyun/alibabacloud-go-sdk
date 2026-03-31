// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatConfigurationResponseBody) *CreateChatConfigurationResponse
	GetBody() *CreateChatConfigurationResponseBody
}

type CreateChatConfigurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateChatConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatConfigurationResponse) GetBody() *CreateChatConfigurationResponseBody {
	return s.Body
}

func (s *CreateChatConfigurationResponse) SetHeaders(v map[string]*string) *CreateChatConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateChatConfigurationResponse) SetStatusCode(v int32) *CreateChatConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatConfigurationResponse) SetBody(v *CreateChatConfigurationResponseBody) *CreateChatConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateChatConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
