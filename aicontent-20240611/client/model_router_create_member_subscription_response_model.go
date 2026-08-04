// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateMemberSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateMemberSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateMemberSubscriptionResponseBody) *ModelRouterCreateMemberSubscriptionResponse
	GetBody() *ModelRouterCreateMemberSubscriptionResponseBody
}

type ModelRouterCreateMemberSubscriptionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateMemberSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateMemberSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateMemberSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateMemberSubscriptionResponse) GetBody() *ModelRouterCreateMemberSubscriptionResponseBody {
	return s.Body
}

func (s *ModelRouterCreateMemberSubscriptionResponse) SetHeaders(v map[string]*string) *ModelRouterCreateMemberSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponse) SetStatusCode(v int32) *ModelRouterCreateMemberSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponse) SetBody(v *ModelRouterCreateMemberSubscriptionResponseBody) *ModelRouterCreateMemberSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
