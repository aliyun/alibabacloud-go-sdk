// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayApiGroupResponseBody) *UpdateGatewayApiGroupResponse
	GetBody() *UpdateGatewayApiGroupResponseBody
}

type UpdateGatewayApiGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayApiGroupResponse) GetBody() *UpdateGatewayApiGroupResponseBody {
	return s.Body
}

func (s *UpdateGatewayApiGroupResponse) SetHeaders(v map[string]*string) *UpdateGatewayApiGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayApiGroupResponse) SetStatusCode(v int32) *UpdateGatewayApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayApiGroupResponse) SetBody(v *UpdateGatewayApiGroupResponseBody) *UpdateGatewayApiGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayApiGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
