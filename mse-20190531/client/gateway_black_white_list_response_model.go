// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayBlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GatewayBlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GatewayBlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *GatewayBlackWhiteListResponseBody) *GatewayBlackWhiteListResponse
	GetBody() *GatewayBlackWhiteListResponseBody
}

type GatewayBlackWhiteListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GatewayBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GatewayBlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s GatewayBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GatewayBlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GatewayBlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GatewayBlackWhiteListResponse) GetBody() *GatewayBlackWhiteListResponseBody {
	return s.Body
}

func (s *GatewayBlackWhiteListResponse) SetHeaders(v map[string]*string) *GatewayBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GatewayBlackWhiteListResponse) SetStatusCode(v int32) *GatewayBlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GatewayBlackWhiteListResponse) SetBody(v *GatewayBlackWhiteListResponseBody) *GatewayBlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *GatewayBlackWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
