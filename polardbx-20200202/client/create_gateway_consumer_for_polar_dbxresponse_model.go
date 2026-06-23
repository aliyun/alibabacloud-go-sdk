// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayConsumerForPolarDBXResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayConsumerForPolarDBXResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayConsumerForPolarDBXResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayConsumerForPolarDBXResponseBody) *CreateGatewayConsumerForPolarDBXResponse
	GetBody() *CreateGatewayConsumerForPolarDBXResponseBody
}

type CreateGatewayConsumerForPolarDBXResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayConsumerForPolarDBXResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayConsumerForPolarDBXResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayConsumerForPolarDBXResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayConsumerForPolarDBXResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayConsumerForPolarDBXResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayConsumerForPolarDBXResponse) GetBody() *CreateGatewayConsumerForPolarDBXResponseBody {
	return s.Body
}

func (s *CreateGatewayConsumerForPolarDBXResponse) SetHeaders(v map[string]*string) *CreateGatewayConsumerForPolarDBXResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponse) SetStatusCode(v int32) *CreateGatewayConsumerForPolarDBXResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponse) SetBody(v *CreateGatewayConsumerForPolarDBXResponseBody) *CreateGatewayConsumerForPolarDBXResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
