// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteCORSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteCORSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteCORSResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteCORSResponseBody) *UpdateGatewayRouteCORSResponse
	GetBody() *UpdateGatewayRouteCORSResponseBody
}

type UpdateGatewayRouteCORSResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteCORSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteCORSResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteCORSResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteCORSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteCORSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteCORSResponse) GetBody() *UpdateGatewayRouteCORSResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteCORSResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteCORSResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteCORSResponse) SetStatusCode(v int32) *UpdateGatewayRouteCORSResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteCORSResponse) SetBody(v *UpdateGatewayRouteCORSResponseBody) *UpdateGatewayRouteCORSResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteCORSResponse) Validate() error {
	return dara.Validate(s)
}
