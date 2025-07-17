// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNotificationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNotificationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNotificationConfigurationResponseBody) *DeleteNotificationConfigurationResponse
	GetBody() *DeleteNotificationConfigurationResponseBody
}

type DeleteNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNotificationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNotificationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNotificationConfigurationResponse) GetBody() *DeleteNotificationConfigurationResponseBody {
	return s.Body
}

func (s *DeleteNotificationConfigurationResponse) SetHeaders(v map[string]*string) *DeleteNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotificationConfigurationResponse) SetStatusCode(v int32) *DeleteNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotificationConfigurationResponse) SetBody(v *DeleteNotificationConfigurationResponseBody) *DeleteNotificationConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteNotificationConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
