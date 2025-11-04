// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNotificationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNotificationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNotificationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNotificationConfigurationResponseBody) *ModifyNotificationConfigurationResponse
	GetBody() *ModifyNotificationConfigurationResponseBody
}

type ModifyNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNotificationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNotificationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNotificationConfigurationResponse) GetBody() *ModifyNotificationConfigurationResponseBody {
	return s.Body
}

func (s *ModifyNotificationConfigurationResponse) SetHeaders(v map[string]*string) *ModifyNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyNotificationConfigurationResponse) SetStatusCode(v int32) *ModifyNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNotificationConfigurationResponse) SetBody(v *ModifyNotificationConfigurationResponseBody) *ModifyNotificationConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyNotificationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
