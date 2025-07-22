// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachExpressConnectRouterChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachExpressConnectRouterChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DetachExpressConnectRouterChildInstanceResponseBody) *DetachExpressConnectRouterChildInstanceResponse
	GetBody() *DetachExpressConnectRouterChildInstanceResponseBody
}

type DetachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachExpressConnectRouterChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachExpressConnectRouterChildInstanceResponse) GetBody() *DetachExpressConnectRouterChildInstanceResponseBody {
	return s.Body
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DetachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetBody(v *DetachExpressConnectRouterChildInstanceResponseBody) *DetachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) Validate() error {
	return dara.Validate(s)
}
