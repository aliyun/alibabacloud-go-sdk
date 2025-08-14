// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterCidrResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterCidrResponseBody) *ListTransitRouterCidrResponse
	GetBody() *ListTransitRouterCidrResponseBody
}

type ListTransitRouterCidrResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterCidrResponse) GetBody() *ListTransitRouterCidrResponseBody {
	return s.Body
}

func (s *ListTransitRouterCidrResponse) SetHeaders(v map[string]*string) *ListTransitRouterCidrResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterCidrResponse) SetStatusCode(v int32) *ListTransitRouterCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterCidrResponse) SetBody(v *ListTransitRouterCidrResponseBody) *ListTransitRouterCidrResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterCidrResponse) Validate() error {
	return dara.Validate(s)
}
