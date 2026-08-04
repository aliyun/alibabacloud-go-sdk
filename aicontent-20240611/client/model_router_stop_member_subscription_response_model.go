// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterStopMemberSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterStopMemberSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterStopMemberSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterStopMemberSubscriptionResponseBody) *ModelRouterStopMemberSubscriptionResponse
	GetBody() *ModelRouterStopMemberSubscriptionResponseBody
}

type ModelRouterStopMemberSubscriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterStopMemberSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterStopMemberSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterStopMemberSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterStopMemberSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterStopMemberSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterStopMemberSubscriptionResponse) GetBody() *ModelRouterStopMemberSubscriptionResponseBody {
	return s.Body
}

func (s *ModelRouterStopMemberSubscriptionResponse) SetHeaders(v map[string]*string) *ModelRouterStopMemberSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponse) SetStatusCode(v int32) *ModelRouterStopMemberSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponse) SetBody(v *ModelRouterStopMemberSubscriptionResponseBody) *ModelRouterStopMemberSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterStopMemberSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
