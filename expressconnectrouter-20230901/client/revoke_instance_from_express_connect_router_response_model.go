// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeInstanceFromExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *RevokeInstanceFromExpressConnectRouterResponseBody) *RevokeInstanceFromExpressConnectRouterResponse
	GetBody() *RevokeInstanceFromExpressConnectRouterResponseBody
}

type RevokeInstanceFromExpressConnectRouterResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) GetBody() *RevokeInstanceFromExpressConnectRouterResponseBody {
	return s.Body
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetBody(v *RevokeInstanceFromExpressConnectRouterResponseBody) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
