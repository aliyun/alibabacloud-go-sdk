// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayNameResponseBody) *UpdateGatewayNameResponse
	GetBody() *UpdateGatewayNameResponseBody
}

type UpdateGatewayNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayNameResponse) GetBody() *UpdateGatewayNameResponseBody {
	return s.Body
}

func (s *UpdateGatewayNameResponse) SetHeaders(v map[string]*string) *UpdateGatewayNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayNameResponse) SetStatusCode(v int32) *UpdateGatewayNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayNameResponse) SetBody(v *UpdateGatewayNameResponseBody) *UpdateGatewayNameResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayNameResponse) Validate() error {
	return dara.Validate(s)
}
