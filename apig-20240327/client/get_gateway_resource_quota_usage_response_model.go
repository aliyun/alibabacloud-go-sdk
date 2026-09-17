// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResourceQuotaUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayResourceQuotaUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayResourceQuotaUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayResourceQuotaUsageResponseBody) *GetGatewayResourceQuotaUsageResponse
	GetBody() *GetGatewayResourceQuotaUsageResponseBody
}

type GetGatewayResourceQuotaUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayResourceQuotaUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayResourceQuotaUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayResourceQuotaUsageResponse) GetBody() *GetGatewayResourceQuotaUsageResponseBody {
	return s.Body
}

func (s *GetGatewayResourceQuotaUsageResponse) SetHeaders(v map[string]*string) *GetGatewayResourceQuotaUsageResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponse) SetStatusCode(v int32) *GetGatewayResourceQuotaUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponse) SetBody(v *GetGatewayResourceQuotaUsageResponseBody) *GetGatewayResourceQuotaUsageResponse {
	s.Body = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
