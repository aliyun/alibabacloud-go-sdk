// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartGatewayResponse
	GetStatusCode() *int32
	SetBody(v *RestartGatewayResponseBody) *RestartGatewayResponse
	GetBody() *RestartGatewayResponseBody
}

type RestartGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartGatewayResponse) GoString() string {
	return s.String()
}

func (s *RestartGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartGatewayResponse) GetBody() *RestartGatewayResponseBody {
	return s.Body
}

func (s *RestartGatewayResponse) SetHeaders(v map[string]*string) *RestartGatewayResponse {
	s.Headers = v
	return s
}

func (s *RestartGatewayResponse) SetStatusCode(v int32) *RestartGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartGatewayResponse) SetBody(v *RestartGatewayResponseBody) *RestartGatewayResponse {
	s.Body = v
	return s
}

func (s *RestartGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
