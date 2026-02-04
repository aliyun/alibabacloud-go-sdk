// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeregisterTransitRouterMulticastGroupSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeregisterTransitRouterMulticastGroupSourcesResponse
	GetStatusCode() *int32
	SetBody(v *DeregisterTransitRouterMulticastGroupSourcesResponseBody) *DeregisterTransitRouterMulticastGroupSourcesResponse
	GetBody() *DeregisterTransitRouterMulticastGroupSourcesResponseBody
}

type DeregisterTransitRouterMulticastGroupSourcesResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterTransitRouterMulticastGroupSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterTransitRouterMulticastGroupSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupSourcesResponse) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) GetBody() *DeregisterTransitRouterMulticastGroupSourcesResponseBody {
	return s.Body
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) SetHeaders(v map[string]*string) *DeregisterTransitRouterMulticastGroupSourcesResponse {
	s.Headers = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) SetStatusCode(v int32) *DeregisterTransitRouterMulticastGroupSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) SetBody(v *DeregisterTransitRouterMulticastGroupSourcesResponseBody) *DeregisterTransitRouterMulticastGroupSourcesResponse {
	s.Body = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
