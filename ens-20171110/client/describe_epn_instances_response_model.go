// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEpnInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEpnInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEpnInstancesResponseBody) *DescribeEpnInstancesResponse
	GetBody() *DescribeEpnInstancesResponseBody
}

type DescribeEpnInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEpnInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEpnInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEpnInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEpnInstancesResponse) GetBody() *DescribeEpnInstancesResponseBody {
	return s.Body
}

func (s *DescribeEpnInstancesResponse) SetHeaders(v map[string]*string) *DescribeEpnInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnInstancesResponse) SetStatusCode(v int32) *DescribeEpnInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEpnInstancesResponse) SetBody(v *DescribeEpnInstancesResponseBody) *DescribeEpnInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeEpnInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
