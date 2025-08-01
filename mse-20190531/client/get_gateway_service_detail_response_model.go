// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayServiceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayServiceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayServiceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayServiceDetailResponseBody) *GetGatewayServiceDetailResponse
	GetBody() *GetGatewayServiceDetailResponseBody
}

type GetGatewayServiceDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayServiceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayServiceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayServiceDetailResponse) GetBody() *GetGatewayServiceDetailResponseBody {
	return s.Body
}

func (s *GetGatewayServiceDetailResponse) SetHeaders(v map[string]*string) *GetGatewayServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayServiceDetailResponse) SetStatusCode(v int32) *GetGatewayServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayServiceDetailResponse) SetBody(v *GetGatewayServiceDetailResponseBody) *GetGatewayServiceDetailResponse {
	s.Body = v
	return s
}

func (s *GetGatewayServiceDetailResponse) Validate() error {
	return dara.Validate(s)
}
