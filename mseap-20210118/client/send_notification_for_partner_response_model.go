// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendNotificationForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendNotificationForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *SendNotificationForPartnerResponseBody) *SendNotificationForPartnerResponse
	GetBody() *SendNotificationForPartnerResponseBody
}

type SendNotificationForPartnerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendNotificationForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendNotificationForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationForPartnerResponse) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendNotificationForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendNotificationForPartnerResponse) GetBody() *SendNotificationForPartnerResponseBody {
	return s.Body
}

func (s *SendNotificationForPartnerResponse) SetHeaders(v map[string]*string) *SendNotificationForPartnerResponse {
	s.Headers = v
	return s
}

func (s *SendNotificationForPartnerResponse) SetStatusCode(v int32) *SendNotificationForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SendNotificationForPartnerResponse) SetBody(v *SendNotificationForPartnerResponseBody) *SendNotificationForPartnerResponse {
	s.Body = v
	return s
}

func (s *SendNotificationForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
