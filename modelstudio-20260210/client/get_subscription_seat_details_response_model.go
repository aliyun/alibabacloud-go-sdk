// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionSeatDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubscriptionSeatDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubscriptionSeatDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetSubscriptionSeatDetailsResponseBody) *GetSubscriptionSeatDetailsResponse
	GetBody() *GetSubscriptionSeatDetailsResponseBody
}

type GetSubscriptionSeatDetailsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionSeatDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionSeatDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubscriptionSeatDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubscriptionSeatDetailsResponse) GetBody() *GetSubscriptionSeatDetailsResponseBody {
	return s.Body
}

func (s *GetSubscriptionSeatDetailsResponse) SetHeaders(v map[string]*string) *GetSubscriptionSeatDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionSeatDetailsResponse) SetStatusCode(v int32) *GetSubscriptionSeatDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponse) SetBody(v *GetSubscriptionSeatDetailsResponseBody) *GetSubscriptionSeatDetailsResponse {
	s.Body = v
	return s
}

func (s *GetSubscriptionSeatDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
