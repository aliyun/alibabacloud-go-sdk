// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayServiceVersionResponseBody) *UpdateGatewayServiceVersionResponse
	GetBody() *UpdateGatewayServiceVersionResponseBody
}

type UpdateGatewayServiceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayServiceVersionResponse) GetBody() *UpdateGatewayServiceVersionResponseBody {
	return s.Body
}

func (s *UpdateGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceVersionResponse) SetStatusCode(v int32) *UpdateGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceVersionResponse) SetBody(v *UpdateGatewayServiceVersionResponseBody) *UpdateGatewayServiceVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
