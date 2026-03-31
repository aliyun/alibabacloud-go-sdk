// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetChatConfigurationResponseBody) *GetChatConfigurationResponse
	GetBody() *GetChatConfigurationResponseBody
}

type GetChatConfigurationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetChatConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatConfigurationResponse) GetBody() *GetChatConfigurationResponseBody {
	return s.Body
}

func (s *GetChatConfigurationResponse) SetHeaders(v map[string]*string) *GetChatConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetChatConfigurationResponse) SetStatusCode(v int32) *GetChatConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatConfigurationResponse) SetBody(v *GetChatConfigurationResponseBody) *GetChatConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetChatConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
