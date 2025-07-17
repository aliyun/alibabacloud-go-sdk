// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotificationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNotificationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNotificationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateNotificationConfigurationResponseBody) *CreateNotificationConfigurationResponse
	GetBody() *CreateNotificationConfigurationResponseBody
}

type CreateNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNotificationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNotificationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNotificationConfigurationResponse) GetBody() *CreateNotificationConfigurationResponseBody {
	return s.Body
}

func (s *CreateNotificationConfigurationResponse) SetHeaders(v map[string]*string) *CreateNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateNotificationConfigurationResponse) SetStatusCode(v int32) *CreateNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNotificationConfigurationResponse) SetBody(v *CreateNotificationConfigurationResponseBody) *CreateNotificationConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateNotificationConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
