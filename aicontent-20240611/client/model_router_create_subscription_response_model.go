// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateSubscriptionResponseBody) *ModelRouterCreateSubscriptionResponse
	GetBody() *ModelRouterCreateSubscriptionResponseBody
}

type ModelRouterCreateSubscriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateSubscriptionResponse) GetBody() *ModelRouterCreateSubscriptionResponseBody {
	return s.Body
}

func (s *ModelRouterCreateSubscriptionResponse) SetHeaders(v map[string]*string) *ModelRouterCreateSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateSubscriptionResponse) SetStatusCode(v int32) *ModelRouterCreateSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateSubscriptionResponse) SetBody(v *ModelRouterCreateSubscriptionResponseBody) *ModelRouterCreateSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
