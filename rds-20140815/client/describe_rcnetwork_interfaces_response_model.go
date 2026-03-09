// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCNetworkInterfacesResponseBody) *DescribeRCNetworkInterfacesResponse
	GetBody() *DescribeRCNetworkInterfacesResponseBody
}

type DescribeRCNetworkInterfacesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCNetworkInterfacesResponse) GetBody() *DescribeRCNetworkInterfacesResponseBody {
	return s.Body
}

func (s *DescribeRCNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DescribeRCNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponse) SetStatusCode(v int32) *DescribeRCNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCNetworkInterfacesResponse) SetBody(v *DescribeRCNetworkInterfacesResponseBody) *DescribeRCNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *DescribeRCNetworkInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
