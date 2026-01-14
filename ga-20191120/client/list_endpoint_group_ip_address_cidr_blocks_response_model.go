// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupIpAddressCidrBlocksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEndpointGroupIpAddressCidrBlocksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEndpointGroupIpAddressCidrBlocksResponse
	GetStatusCode() *int32
	SetBody(v *ListEndpointGroupIpAddressCidrBlocksResponseBody) *ListEndpointGroupIpAddressCidrBlocksResponse
	GetBody() *ListEndpointGroupIpAddressCidrBlocksResponseBody
}

type ListEndpointGroupIpAddressCidrBlocksResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEndpointGroupIpAddressCidrBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEndpointGroupIpAddressCidrBlocksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupIpAddressCidrBlocksResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) GetBody() *ListEndpointGroupIpAddressCidrBlocksResponseBody {
	return s.Body
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) SetHeaders(v map[string]*string) *ListEndpointGroupIpAddressCidrBlocksResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) SetStatusCode(v int32) *ListEndpointGroupIpAddressCidrBlocksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) SetBody(v *ListEndpointGroupIpAddressCidrBlocksResponseBody) *ListEndpointGroupIpAddressCidrBlocksResponse {
	s.Body = v
	return s
}

func (s *ListEndpointGroupIpAddressCidrBlocksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
