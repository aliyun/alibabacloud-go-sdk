// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterResponseBody) *DeleteTransitRouterResponse
	GetBody() *DeleteTransitRouterResponseBody
}

type DeleteTransitRouterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterResponse) GetBody() *DeleteTransitRouterResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterResponse) SetStatusCode(v int32) *DeleteTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterResponse) SetBody(v *DeleteTransitRouterResponseBody) *DeleteTransitRouterResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterResponse) Validate() error {
	return dara.Validate(s)
}
