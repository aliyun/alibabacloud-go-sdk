// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayWebsocketUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayWebsocketUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayWebsocketUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayWebsocketUsageResponseBody) *GetGatewayWebsocketUsageResponse
	GetBody() *GetGatewayWebsocketUsageResponseBody
}

type GetGatewayWebsocketUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayWebsocketUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayWebsocketUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayWebsocketUsageResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayWebsocketUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayWebsocketUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayWebsocketUsageResponse) GetBody() *GetGatewayWebsocketUsageResponseBody {
	return s.Body
}

func (s *GetGatewayWebsocketUsageResponse) SetHeaders(v map[string]*string) *GetGatewayWebsocketUsageResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayWebsocketUsageResponse) SetStatusCode(v int32) *GetGatewayWebsocketUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponse) SetBody(v *GetGatewayWebsocketUsageResponseBody) *GetGatewayWebsocketUsageResponse {
	s.Body = v
	return s
}

func (s *GetGatewayWebsocketUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
