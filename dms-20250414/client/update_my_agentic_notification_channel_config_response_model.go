// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMyAgenticNotificationChannelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMyAgenticNotificationChannelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMyAgenticNotificationChannelConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMyAgenticNotificationChannelConfigResponseBody) *UpdateMyAgenticNotificationChannelConfigResponse
	GetBody() *UpdateMyAgenticNotificationChannelConfigResponseBody
}

type UpdateMyAgenticNotificationChannelConfigResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMyAgenticNotificationChannelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMyAgenticNotificationChannelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMyAgenticNotificationChannelConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) GetBody() *UpdateMyAgenticNotificationChannelConfigResponseBody {
	return s.Body
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) SetHeaders(v map[string]*string) *UpdateMyAgenticNotificationChannelConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) SetStatusCode(v int32) *UpdateMyAgenticNotificationChannelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) SetBody(v *UpdateMyAgenticNotificationChannelConfigResponseBody) *UpdateMyAgenticNotificationChannelConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
