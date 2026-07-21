// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterListSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterListSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterListSubscriptionsResponseBody) *ModelRouterListSubscriptionsResponse
	GetBody() *ModelRouterListSubscriptionsResponseBody
}

type ModelRouterListSubscriptionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterListSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterListSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterListSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterListSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterListSubscriptionsResponse) GetBody() *ModelRouterListSubscriptionsResponseBody {
	return s.Body
}

func (s *ModelRouterListSubscriptionsResponse) SetHeaders(v map[string]*string) *ModelRouterListSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterListSubscriptionsResponse) SetStatusCode(v int32) *ModelRouterListSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterListSubscriptionsResponse) SetBody(v *ModelRouterListSubscriptionsResponseBody) *ModelRouterListSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterListSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
