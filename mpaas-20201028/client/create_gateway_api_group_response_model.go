// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayApiGroupResponseBody) *CreateGatewayApiGroupResponse
	GetBody() *CreateGatewayApiGroupResponseBody
}

type CreateGatewayApiGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayApiGroupResponse) GetBody() *CreateGatewayApiGroupResponseBody {
	return s.Body
}

func (s *CreateGatewayApiGroupResponse) SetHeaders(v map[string]*string) *CreateGatewayApiGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayApiGroupResponse) SetStatusCode(v int32) *CreateGatewayApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayApiGroupResponse) SetBody(v *CreateGatewayApiGroupResponseBody) *CreateGatewayApiGroupResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayApiGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
