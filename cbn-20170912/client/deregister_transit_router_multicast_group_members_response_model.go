// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeregisterTransitRouterMulticastGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeregisterTransitRouterMulticastGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeregisterTransitRouterMulticastGroupMembersResponseBody) *DeregisterTransitRouterMulticastGroupMembersResponse
	GetBody() *DeregisterTransitRouterMulticastGroupMembersResponseBody
}

type DeregisterTransitRouterMulticastGroupMembersResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterTransitRouterMulticastGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterTransitRouterMulticastGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) GetBody() *DeregisterTransitRouterMulticastGroupMembersResponseBody {
	return s.Body
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) SetHeaders(v map[string]*string) *DeregisterTransitRouterMulticastGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) SetStatusCode(v int32) *DeregisterTransitRouterMulticastGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) SetBody(v *DeregisterTransitRouterMulticastGroupMembersResponseBody) *DeregisterTransitRouterMulticastGroupMembersResponse {
	s.Body = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
