// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayApiGroupResponseBody) *GetGatewayApiGroupResponse
	GetBody() *GetGatewayApiGroupResponseBody
}

type GetGatewayApiGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayApiGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayApiGroupResponse) GetBody() *GetGatewayApiGroupResponseBody {
	return s.Body
}

func (s *GetGatewayApiGroupResponse) SetHeaders(v map[string]*string) *GetGatewayApiGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayApiGroupResponse) SetStatusCode(v int32) *GetGatewayApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayApiGroupResponse) SetBody(v *GetGatewayApiGroupResponseBody) *GetGatewayApiGroupResponse {
	s.Body = v
	return s
}

func (s *GetGatewayApiGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
