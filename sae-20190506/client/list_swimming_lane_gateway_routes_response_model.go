// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGatewayRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSwimmingLaneGatewayRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSwimmingLaneGatewayRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListSwimmingLaneGatewayRoutesResponseBody) *ListSwimmingLaneGatewayRoutesResponse
	GetBody() *ListSwimmingLaneGatewayRoutesResponseBody
}

type ListSwimmingLaneGatewayRoutesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSwimmingLaneGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSwimmingLaneGatewayRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGatewayRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSwimmingLaneGatewayRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSwimmingLaneGatewayRoutesResponse) GetBody() *ListSwimmingLaneGatewayRoutesResponseBody {
	return s.Body
}

func (s *ListSwimmingLaneGatewayRoutesResponse) SetHeaders(v map[string]*string) *ListSwimmingLaneGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponse) SetStatusCode(v int32) *ListSwimmingLaneGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponse) SetBody(v *ListSwimmingLaneGatewayRoutesResponseBody) *ListSwimmingLaneGatewayRoutesResponse {
	s.Body = v
	return s
}

func (s *ListSwimmingLaneGatewayRoutesResponse) Validate() error {
	return dara.Validate(s)
}
