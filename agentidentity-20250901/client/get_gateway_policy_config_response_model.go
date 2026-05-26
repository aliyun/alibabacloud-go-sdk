// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayPolicyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayPolicyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayPolicyConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayPolicyConfigResponseBody) *GetGatewayPolicyConfigResponse
	GetBody() *GetGatewayPolicyConfigResponseBody
}

type GetGatewayPolicyConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayPolicyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayPolicyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayPolicyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayPolicyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayPolicyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayPolicyConfigResponse) GetBody() *GetGatewayPolicyConfigResponseBody {
	return s.Body
}

func (s *GetGatewayPolicyConfigResponse) SetHeaders(v map[string]*string) *GetGatewayPolicyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayPolicyConfigResponse) SetStatusCode(v int32) *GetGatewayPolicyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayPolicyConfigResponse) SetBody(v *GetGatewayPolicyConfigResponseBody) *GetGatewayPolicyConfigResponse {
	s.Body = v
	return s
}

func (s *GetGatewayPolicyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
