// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceVpcEndpointLinkedVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceVpcEndpointLinkedVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceVpcEndpointLinkedVpcResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceVpcEndpointLinkedVpcResponseBody) *DeleteInstanceVpcEndpointLinkedVpcResponse
	GetBody() *DeleteInstanceVpcEndpointLinkedVpcResponseBody
}

type DeleteInstanceVpcEndpointLinkedVpcResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceVpcEndpointLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) GetBody() *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	return s.Body
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetHeaders(v map[string]*string) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetStatusCode(v int32) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetBody(v *DeleteInstanceVpcEndpointLinkedVpcResponseBody) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) Validate() error {
	return dara.Validate(s)
}
