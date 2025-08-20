// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendNotificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendNotificationsResponse
	GetStatusCode() *int32
}

type SendNotificationsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SendNotificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsResponse) GoString() string {
	return s.String()
}

func (s *SendNotificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendNotificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendNotificationsResponse) SetHeaders(v map[string]*string) *SendNotificationsResponse {
	s.Headers = v
	return s
}

func (s *SendNotificationsResponse) SetStatusCode(v int32) *SendNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendNotificationsResponse) Validate() error {
	return dara.Validate(s)
}
