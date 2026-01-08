// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessengerSubscriptionTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessengerSubscriptionTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessengerSubscriptionTokenResponse
	GetStatusCode() *int32
	SetBody(v *ListMessengerSubscriptionTokenResponseBody) *ListMessengerSubscriptionTokenResponse
	GetBody() *ListMessengerSubscriptionTokenResponseBody
}

type ListMessengerSubscriptionTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessengerSubscriptionTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessengerSubscriptionTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessengerSubscriptionTokenResponse) GoString() string {
	return s.String()
}

func (s *ListMessengerSubscriptionTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessengerSubscriptionTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessengerSubscriptionTokenResponse) GetBody() *ListMessengerSubscriptionTokenResponseBody {
	return s.Body
}

func (s *ListMessengerSubscriptionTokenResponse) SetHeaders(v map[string]*string) *ListMessengerSubscriptionTokenResponse {
	s.Headers = v
	return s
}

func (s *ListMessengerSubscriptionTokenResponse) SetStatusCode(v int32) *ListMessengerSubscriptionTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponse) SetBody(v *ListMessengerSubscriptionTokenResponseBody) *ListMessengerSubscriptionTokenResponse {
	s.Body = v
	return s
}

func (s *ListMessengerSubscriptionTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
