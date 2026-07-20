// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationPendNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotificationPendNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotificationPendNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetNotificationPendNumberResponseBody) *GetNotificationPendNumberResponse
	GetBody() *GetNotificationPendNumberResponseBody
}

type GetNotificationPendNumberResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotificationPendNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotificationPendNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationPendNumberResponse) GoString() string {
	return s.String()
}

func (s *GetNotificationPendNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotificationPendNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotificationPendNumberResponse) GetBody() *GetNotificationPendNumberResponseBody {
	return s.Body
}

func (s *GetNotificationPendNumberResponse) SetHeaders(v map[string]*string) *GetNotificationPendNumberResponse {
	s.Headers = v
	return s
}

func (s *GetNotificationPendNumberResponse) SetStatusCode(v int32) *GetNotificationPendNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotificationPendNumberResponse) SetBody(v *GetNotificationPendNumberResponseBody) *GetNotificationPendNumberResponse {
	s.Body = v
	return s
}

func (s *GetNotificationPendNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
