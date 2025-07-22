// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressConnectRouterResponseBody) *CreateExpressConnectRouterResponse
	GetBody() *CreateExpressConnectRouterResponseBody
}

type CreateExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressConnectRouterResponse) GetBody() *CreateExpressConnectRouterResponseBody {
	return s.Body
}

func (s *CreateExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetStatusCode(v int32) *CreateExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetBody(v *CreateExpressConnectRouterResponseBody) *CreateExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *CreateExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
