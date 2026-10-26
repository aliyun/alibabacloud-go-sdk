// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubscriptionResponseBody) *CreateSubscriptionResponse
	GetBody() *CreateSubscriptionResponseBody
}

type CreateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubscriptionResponse) GetBody() *CreateSubscriptionResponseBody {
	return s.Body
}

func (s *CreateSubscriptionResponse) SetHeaders(v map[string]*string) *CreateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionResponse) SetStatusCode(v int32) *CreateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionResponse) SetBody(v *CreateSubscriptionResponseBody) *CreateSubscriptionResponse {
	s.Body = v
	return s
}

func (s *CreateSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
