// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterCidrAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterCidrAllocationResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterCidrAllocationResponseBody) *ListTransitRouterCidrAllocationResponse
	GetBody() *ListTransitRouterCidrAllocationResponseBody
}

type ListTransitRouterCidrAllocationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterCidrAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterCidrAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrAllocationResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterCidrAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterCidrAllocationResponse) GetBody() *ListTransitRouterCidrAllocationResponseBody {
	return s.Body
}

func (s *ListTransitRouterCidrAllocationResponse) SetHeaders(v map[string]*string) *ListTransitRouterCidrAllocationResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterCidrAllocationResponse) SetStatusCode(v int32) *ListTransitRouterCidrAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponse) SetBody(v *ListTransitRouterCidrAllocationResponseBody) *ListTransitRouterCidrAllocationResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterCidrAllocationResponse) Validate() error {
	return dara.Validate(s)
}
