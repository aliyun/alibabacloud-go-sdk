// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantInstanceToExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *GrantInstanceToExpressConnectRouterResponseBody) *GrantInstanceToExpressConnectRouterResponse
	GetBody() *GrantInstanceToExpressConnectRouterResponseBody
}

type GrantInstanceToExpressConnectRouterResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantInstanceToExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantInstanceToExpressConnectRouterResponse) GetBody() *GrantInstanceToExpressConnectRouterResponseBody {
	return s.Body
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *GrantInstanceToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetBody(v *GrantInstanceToExpressConnectRouterResponseBody) *GrantInstanceToExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
