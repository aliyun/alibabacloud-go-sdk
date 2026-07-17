// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPolarClawGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartPolarClawGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartPolarClawGatewayResponse
	GetStatusCode() *int32
	SetBody(v *RestartPolarClawGatewayResponseBody) *RestartPolarClawGatewayResponse
	GetBody() *RestartPolarClawGatewayResponseBody
}

type RestartPolarClawGatewayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartPolarClawGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartPolarClawGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartPolarClawGatewayResponse) GoString() string {
	return s.String()
}

func (s *RestartPolarClawGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartPolarClawGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartPolarClawGatewayResponse) GetBody() *RestartPolarClawGatewayResponseBody {
	return s.Body
}

func (s *RestartPolarClawGatewayResponse) SetHeaders(v map[string]*string) *RestartPolarClawGatewayResponse {
	s.Headers = v
	return s
}

func (s *RestartPolarClawGatewayResponse) SetStatusCode(v int32) *RestartPolarClawGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartPolarClawGatewayResponse) SetBody(v *RestartPolarClawGatewayResponseBody) *RestartPolarClawGatewayResponse {
	s.Body = v
	return s
}

func (s *RestartPolarClawGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
