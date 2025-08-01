// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayAuthDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayAuthDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayAuthDetailResponseBody) *GetGatewayAuthDetailResponse
	GetBody() *GetGatewayAuthDetailResponseBody
}

type GetGatewayAuthDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayAuthDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayAuthDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayAuthDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayAuthDetailResponse) GetBody() *GetGatewayAuthDetailResponseBody {
	return s.Body
}

func (s *GetGatewayAuthDetailResponse) SetHeaders(v map[string]*string) *GetGatewayAuthDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayAuthDetailResponse) SetStatusCode(v int32) *GetGatewayAuthDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayAuthDetailResponse) SetBody(v *GetGatewayAuthDetailResponseBody) *GetGatewayAuthDetailResponse {
	s.Body = v
	return s
}

func (s *GetGatewayAuthDetailResponse) Validate() error {
	return dara.Validate(s)
}
