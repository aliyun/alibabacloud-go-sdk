// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayApiResponseBody) *CreateGatewayApiResponse
	GetBody() *CreateGatewayApiResponseBody
}

type CreateGatewayApiResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayApiResponse) GetBody() *CreateGatewayApiResponseBody {
	return s.Body
}

func (s *CreateGatewayApiResponse) SetHeaders(v map[string]*string) *CreateGatewayApiResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayApiResponse) SetStatusCode(v int32) *CreateGatewayApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayApiResponse) SetBody(v *CreateGatewayApiResponseBody) *CreateGatewayApiResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
