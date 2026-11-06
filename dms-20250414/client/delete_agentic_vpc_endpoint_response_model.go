// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticVpcEndpointResponseBody) *DeleteAgenticVpcEndpointResponse
	GetBody() *DeleteAgenticVpcEndpointResponseBody
}

type DeleteAgenticVpcEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticVpcEndpointResponse) GetBody() *DeleteAgenticVpcEndpointResponseBody {
	return s.Body
}

func (s *DeleteAgenticVpcEndpointResponse) SetHeaders(v map[string]*string) *DeleteAgenticVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticVpcEndpointResponse) SetStatusCode(v int32) *DeleteAgenticVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticVpcEndpointResponse) SetBody(v *DeleteAgenticVpcEndpointResponseBody) *DeleteAgenticVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
