// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSubscriptionResponseBody) *UpdateSubscriptionResponse
	GetBody() *UpdateSubscriptionResponseBody
}

type UpdateSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubscriptionResponse) GetBody() *UpdateSubscriptionResponseBody {
	return s.Body
}

func (s *UpdateSubscriptionResponse) SetHeaders(v map[string]*string) *UpdateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscriptionResponse) SetStatusCode(v int32) *UpdateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponse) SetBody(v *UpdateSubscriptionResponseBody) *UpdateSubscriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
