// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceVpcEndpointLinkedVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceVpcEndpointLinkedVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceVpcEndpointLinkedVpcResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceVpcEndpointLinkedVpcResponseBody) *CreateInstanceVpcEndpointLinkedVpcResponse
	GetBody() *CreateInstanceVpcEndpointLinkedVpcResponseBody
}

type CreateInstanceVpcEndpointLinkedVpcResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceVpcEndpointLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) GetBody() *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	return s.Body
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetHeaders(v map[string]*string) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetStatusCode(v int32) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetBody(v *CreateInstanceVpcEndpointLinkedVpcResponseBody) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
