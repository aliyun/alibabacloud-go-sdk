// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterTransitRouterMulticastGroupSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterTransitRouterMulticastGroupSourcesResponse
	GetStatusCode() *int32
	SetBody(v *RegisterTransitRouterMulticastGroupSourcesResponseBody) *RegisterTransitRouterMulticastGroupSourcesResponse
	GetBody() *RegisterTransitRouterMulticastGroupSourcesResponseBody
}

type RegisterTransitRouterMulticastGroupSourcesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterTransitRouterMulticastGroupSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterTransitRouterMulticastGroupSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupSourcesResponse) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) GetBody() *RegisterTransitRouterMulticastGroupSourcesResponseBody {
	return s.Body
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) SetHeaders(v map[string]*string) *RegisterTransitRouterMulticastGroupSourcesResponse {
	s.Headers = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) SetStatusCode(v int32) *RegisterTransitRouterMulticastGroupSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) SetBody(v *RegisterTransitRouterMulticastGroupSourcesResponseBody) *RegisterTransitRouterMulticastGroupSourcesResponse {
	s.Body = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
