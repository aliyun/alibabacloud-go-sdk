// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableVpcEndpointConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableVpcEndpointConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DisableVpcEndpointConnectionResponseBody) *DisableVpcEndpointConnectionResponse
	GetBody() *DisableVpcEndpointConnectionResponseBody
}

type DisableVpcEndpointConnectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVpcEndpointConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointConnectionResponse) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableVpcEndpointConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableVpcEndpointConnectionResponse) GetBody() *DisableVpcEndpointConnectionResponseBody {
	return s.Body
}

func (s *DisableVpcEndpointConnectionResponse) SetHeaders(v map[string]*string) *DisableVpcEndpointConnectionResponse {
	s.Headers = v
	return s
}

func (s *DisableVpcEndpointConnectionResponse) SetStatusCode(v int32) *DisableVpcEndpointConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVpcEndpointConnectionResponse) SetBody(v *DisableVpcEndpointConnectionResponseBody) *DisableVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

func (s *DisableVpcEndpointConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
