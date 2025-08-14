// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectRouterChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectRouterChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectRouterChildInstanceResponseBody) *ModifyExpressConnectRouterChildInstanceResponse
	GetBody() *ModifyExpressConnectRouterChildInstanceResponseBody
}

type ModifyExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) GetBody() *ModifyExpressConnectRouterChildInstanceResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) SetBody(v *ModifyExpressConnectRouterChildInstanceResponseBody) *ModifyExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponse) Validate() error {
	return dara.Validate(s)
}
