// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertNotificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertNotificationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertNotificationsResponseBody) *ListAlertNotificationsResponse
	GetBody() *ListAlertNotificationsResponseBody
}

type ListAlertNotificationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertNotificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertNotificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertNotificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertNotificationsResponse) GetBody() *ListAlertNotificationsResponseBody {
	return s.Body
}

func (s *ListAlertNotificationsResponse) SetHeaders(v map[string]*string) *ListAlertNotificationsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertNotificationsResponse) SetStatusCode(v int32) *ListAlertNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertNotificationsResponse) SetBody(v *ListAlertNotificationsResponseBody) *ListAlertNotificationsResponse {
	s.Body = v
	return s
}

func (s *ListAlertNotificationsResponse) Validate() error {
	return dara.Validate(s)
}
