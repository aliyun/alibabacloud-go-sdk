// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAppNotificationForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NotifyAppNotificationForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NotifyAppNotificationForAdminResponse
	GetStatusCode() *int32
	SetBody(v *NotifyAppNotificationForAdminResponseBody) *NotifyAppNotificationForAdminResponse
	GetBody() *NotifyAppNotificationForAdminResponseBody
}

type NotifyAppNotificationForAdminResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyAppNotificationForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyAppNotificationForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s NotifyAppNotificationForAdminResponse) GoString() string {
	return s.String()
}

func (s *NotifyAppNotificationForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NotifyAppNotificationForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NotifyAppNotificationForAdminResponse) GetBody() *NotifyAppNotificationForAdminResponseBody {
	return s.Body
}

func (s *NotifyAppNotificationForAdminResponse) SetHeaders(v map[string]*string) *NotifyAppNotificationForAdminResponse {
	s.Headers = v
	return s
}

func (s *NotifyAppNotificationForAdminResponse) SetStatusCode(v int32) *NotifyAppNotificationForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyAppNotificationForAdminResponse) SetBody(v *NotifyAppNotificationForAdminResponseBody) *NotifyAppNotificationForAdminResponse {
	s.Body = v
	return s
}

func (s *NotifyAppNotificationForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
