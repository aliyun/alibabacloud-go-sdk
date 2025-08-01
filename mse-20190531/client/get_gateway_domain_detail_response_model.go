// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayDomainDetailResponseBody) *GetGatewayDomainDetailResponse
	GetBody() *GetGatewayDomainDetailResponseBody
}

type GetGatewayDomainDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayDomainDetailResponse) GetBody() *GetGatewayDomainDetailResponseBody {
	return s.Body
}

func (s *GetGatewayDomainDetailResponse) SetHeaders(v map[string]*string) *GetGatewayDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayDomainDetailResponse) SetStatusCode(v int32) *GetGatewayDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayDomainDetailResponse) SetBody(v *GetGatewayDomainDetailResponseBody) *GetGatewayDomainDetailResponse {
	s.Body = v
	return s
}

func (s *GetGatewayDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
