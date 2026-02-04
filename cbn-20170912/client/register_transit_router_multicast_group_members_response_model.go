// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterTransitRouterMulticastGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterTransitRouterMulticastGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *RegisterTransitRouterMulticastGroupMembersResponseBody) *RegisterTransitRouterMulticastGroupMembersResponse
	GetBody() *RegisterTransitRouterMulticastGroupMembersResponseBody
}

type RegisterTransitRouterMulticastGroupMembersResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterTransitRouterMulticastGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterTransitRouterMulticastGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) GetBody() *RegisterTransitRouterMulticastGroupMembersResponseBody {
	return s.Body
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) SetHeaders(v map[string]*string) *RegisterTransitRouterMulticastGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) SetStatusCode(v int32) *RegisterTransitRouterMulticastGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) SetBody(v *RegisterTransitRouterMulticastGroupMembersResponseBody) *RegisterTransitRouterMulticastGroupMembersResponse {
	s.Body = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
