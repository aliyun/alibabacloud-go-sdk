// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayApiGroupResponseBody) *DeleteGatewayApiGroupResponse
	GetBody() *DeleteGatewayApiGroupResponseBody
}

type DeleteGatewayApiGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayApiGroupResponse) GetBody() *DeleteGatewayApiGroupResponseBody {
	return s.Body
}

func (s *DeleteGatewayApiGroupResponse) SetHeaders(v map[string]*string) *DeleteGatewayApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayApiGroupResponse) SetStatusCode(v int32) *DeleteGatewayApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayApiGroupResponse) SetBody(v *DeleteGatewayApiGroupResponseBody) *DeleteGatewayApiGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayApiGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
