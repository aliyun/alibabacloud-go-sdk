// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayApiResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayApiResponseBody) *UpdateGatewayApiResponse
	GetBody() *UpdateGatewayApiResponseBody
}

type UpdateGatewayApiResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayApiResponse) GetBody() *UpdateGatewayApiResponseBody {
	return s.Body
}

func (s *UpdateGatewayApiResponse) SetHeaders(v map[string]*string) *UpdateGatewayApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayApiResponse) SetStatusCode(v int32) *UpdateGatewayApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayApiResponse) SetBody(v *UpdateGatewayApiResponseBody) *UpdateGatewayApiResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
