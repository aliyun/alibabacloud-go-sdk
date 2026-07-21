// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterStopSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterStopSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterStopSubscriptionResponseBody) *ModelRouterStopSubscriptionResponse
	GetBody() *ModelRouterStopSubscriptionResponseBody
}

type ModelRouterStopSubscriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterStopSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterStopSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterStopSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterStopSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterStopSubscriptionResponse) GetBody() *ModelRouterStopSubscriptionResponseBody {
	return s.Body
}

func (s *ModelRouterStopSubscriptionResponse) SetHeaders(v map[string]*string) *ModelRouterStopSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterStopSubscriptionResponse) SetStatusCode(v int32) *ModelRouterStopSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterStopSubscriptionResponse) SetBody(v *ModelRouterStopSubscriptionResponseBody) *ModelRouterStopSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterStopSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
