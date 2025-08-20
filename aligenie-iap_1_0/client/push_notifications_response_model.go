// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNotificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushNotificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushNotificationsResponse
	GetStatusCode() *int32
}

type PushNotificationsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PushNotificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsResponse) GoString() string {
	return s.String()
}

func (s *PushNotificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushNotificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushNotificationsResponse) SetHeaders(v map[string]*string) *PushNotificationsResponse {
	s.Headers = v
	return s
}

func (s *PushNotificationsResponse) SetStatusCode(v int32) *PushNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *PushNotificationsResponse) Validate() error {
	return dara.Validate(s)
}
