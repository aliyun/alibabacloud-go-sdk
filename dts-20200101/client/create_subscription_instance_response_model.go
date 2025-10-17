// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubscriptionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubscriptionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubscriptionInstanceResponseBody) *CreateSubscriptionInstanceResponse
	GetBody() *CreateSubscriptionInstanceResponseBody
}

type CreateSubscriptionInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubscriptionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubscriptionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubscriptionInstanceResponse) GetBody() *CreateSubscriptionInstanceResponseBody {
	return s.Body
}

func (s *CreateSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *CreateSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionInstanceResponse) SetStatusCode(v int32) *CreateSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionInstanceResponse) SetBody(v *CreateSubscriptionInstanceResponseBody) *CreateSubscriptionInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateSubscriptionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
