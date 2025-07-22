// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForceDeleteExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForceDeleteExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForceDeleteExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *ForceDeleteExpressConnectRouterResponseBody) *ForceDeleteExpressConnectRouterResponse
	GetBody() *ForceDeleteExpressConnectRouterResponseBody
}

type ForceDeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceDeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForceDeleteExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForceDeleteExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForceDeleteExpressConnectRouterResponse) GetBody() *ForceDeleteExpressConnectRouterResponseBody {
	return s.Body
}

func (s *ForceDeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ForceDeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetStatusCode(v int32) *ForceDeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetBody(v *ForceDeleteExpressConnectRouterResponseBody) *ForceDeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
