// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterMulticastGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterMulticastGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterMulticastGroupsResponseBody) *ListTransitRouterMulticastGroupsResponse
	GetBody() *ListTransitRouterMulticastGroupsResponseBody
}

type ListTransitRouterMulticastGroupsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterMulticastGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterMulticastGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterMulticastGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterMulticastGroupsResponse) GetBody() *ListTransitRouterMulticastGroupsResponseBody {
	return s.Body
}

func (s *ListTransitRouterMulticastGroupsResponse) SetHeaders(v map[string]*string) *ListTransitRouterMulticastGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponse) SetStatusCode(v int32) *ListTransitRouterMulticastGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponse) SetBody(v *ListTransitRouterMulticastGroupsResponseBody) *ListTransitRouterMulticastGroupsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponse) Validate() error {
	return dara.Validate(s)
}
