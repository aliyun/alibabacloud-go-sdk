// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectRouterResponseBody) *ModifyExpressConnectRouterResponse
	GetBody() *ModifyExpressConnectRouterResponseBody
}

type ModifyExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectRouterResponse) GetBody() *ModifyExpressConnectRouterResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetBody(v *ModifyExpressConnectRouterResponseBody) *ModifyExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
