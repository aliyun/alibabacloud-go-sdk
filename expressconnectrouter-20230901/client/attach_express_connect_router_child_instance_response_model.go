// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachExpressConnectRouterChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachExpressConnectRouterChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AttachExpressConnectRouterChildInstanceResponseBody) *AttachExpressConnectRouterChildInstanceResponse
	GetBody() *AttachExpressConnectRouterChildInstanceResponseBody
}

type AttachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachExpressConnectRouterChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachExpressConnectRouterChildInstanceResponse) GetBody() *AttachExpressConnectRouterChildInstanceResponseBody {
	return s.Body
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *AttachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetBody(v *AttachExpressConnectRouterChildInstanceResponseBody) *AttachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) Validate() error {
	return dara.Validate(s)
}
