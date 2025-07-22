// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectRouterResponseBody) *DeleteExpressConnectRouterResponse
	GetBody() *DeleteExpressConnectRouterResponseBody
}

type DeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectRouterResponse) GetBody() *DeleteExpressConnectRouterResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetBody(v *DeleteExpressConnectRouterResponseBody) *DeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
