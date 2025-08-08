// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmNotificationResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmNotificationResponseBody) *ConfirmNotificationResponse
	GetBody() *ConfirmNotificationResponseBody
}

type ConfirmNotificationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotificationResponse) GoString() string {
	return s.String()
}

func (s *ConfirmNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmNotificationResponse) GetBody() *ConfirmNotificationResponseBody {
	return s.Body
}

func (s *ConfirmNotificationResponse) SetHeaders(v map[string]*string) *ConfirmNotificationResponse {
	s.Headers = v
	return s
}

func (s *ConfirmNotificationResponse) SetStatusCode(v int32) *ConfirmNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmNotificationResponse) SetBody(v *ConfirmNotificationResponseBody) *ConfirmNotificationResponse {
	s.Body = v
	return s
}

func (s *ConfirmNotificationResponse) Validate() error {
	return dara.Validate(s)
}
