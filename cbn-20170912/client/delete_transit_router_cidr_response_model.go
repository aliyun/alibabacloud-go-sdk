// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterCidrResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterCidrResponseBody) *DeleteTransitRouterCidrResponse
	GetBody() *DeleteTransitRouterCidrResponseBody
}

type DeleteTransitRouterCidrResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterCidrResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterCidrResponse) GetBody() *DeleteTransitRouterCidrResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterCidrResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterCidrResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterCidrResponse) SetStatusCode(v int32) *DeleteTransitRouterCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterCidrResponse) SetBody(v *DeleteTransitRouterCidrResponseBody) *DeleteTransitRouterCidrResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterCidrResponse) Validate() error {
	return dara.Validate(s)
}
