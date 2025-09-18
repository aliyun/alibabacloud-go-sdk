// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushDeviceNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushDeviceNotificationResponse
	GetStatusCode() *int32
	SetBody(v *PushDeviceNotificationResponseBody) *PushDeviceNotificationResponse
	GetBody() *PushDeviceNotificationResponseBody
}

type PushDeviceNotificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDeviceNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationResponse) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushDeviceNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushDeviceNotificationResponse) GetBody() *PushDeviceNotificationResponseBody {
	return s.Body
}

func (s *PushDeviceNotificationResponse) SetHeaders(v map[string]*string) *PushDeviceNotificationResponse {
	s.Headers = v
	return s
}

func (s *PushDeviceNotificationResponse) SetStatusCode(v int32) *PushDeviceNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDeviceNotificationResponse) SetBody(v *PushDeviceNotificationResponseBody) *PushDeviceNotificationResponse {
	s.Body = v
	return s
}

func (s *PushDeviceNotificationResponse) Validate() error {
	return dara.Validate(s)
}
