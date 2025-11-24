// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateASMGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateASMGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateASMGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateASMGatewayResponseBody) *CreateASMGatewayResponse
	GetBody() *CreateASMGatewayResponseBody
}

type CreateASMGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateASMGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateASMGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateASMGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateASMGatewayResponse) GetBody() *CreateASMGatewayResponseBody {
	return s.Body
}

func (s *CreateASMGatewayResponse) SetHeaders(v map[string]*string) *CreateASMGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateASMGatewayResponse) SetStatusCode(v int32) *CreateASMGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateASMGatewayResponse) SetBody(v *CreateASMGatewayResponseBody) *CreateASMGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateASMGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
