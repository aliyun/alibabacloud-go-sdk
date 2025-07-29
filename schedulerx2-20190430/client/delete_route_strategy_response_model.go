// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteStrategyResponseBody) *DeleteRouteStrategyResponse
	GetBody() *DeleteRouteStrategyResponseBody
}

type DeleteRouteStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteStrategyResponse) GetBody() *DeleteRouteStrategyResponseBody {
	return s.Body
}

func (s *DeleteRouteStrategyResponse) SetHeaders(v map[string]*string) *DeleteRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteStrategyResponse) SetStatusCode(v int32) *DeleteRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteStrategyResponse) SetBody(v *DeleteRouteStrategyResponseBody) *DeleteRouteStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteStrategyResponse) Validate() error {
	return dara.Validate(s)
}
