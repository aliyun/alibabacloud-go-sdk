// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterListMemberSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterListMemberSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterListMemberSubscriptionsResponseBody) *ModelRouterListMemberSubscriptionsResponse
	GetBody() *ModelRouterListMemberSubscriptionsResponseBody
}

type ModelRouterListMemberSubscriptionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterListMemberSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterListMemberSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterListMemberSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterListMemberSubscriptionsResponse) GetBody() *ModelRouterListMemberSubscriptionsResponseBody {
	return s.Body
}

func (s *ModelRouterListMemberSubscriptionsResponse) SetHeaders(v map[string]*string) *ModelRouterListMemberSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponse) SetStatusCode(v int32) *ModelRouterListMemberSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponse) SetBody(v *ModelRouterListMemberSubscriptionsResponseBody) *ModelRouterListMemberSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterListMemberSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
