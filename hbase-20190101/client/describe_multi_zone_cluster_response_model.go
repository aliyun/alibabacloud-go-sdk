// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiZoneClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiZoneClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiZoneClusterResponseBody) *DescribeMultiZoneClusterResponse
	GetBody() *DescribeMultiZoneClusterResponseBody
}

type DescribeMultiZoneClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiZoneClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiZoneClusterResponse) GetBody() *DescribeMultiZoneClusterResponseBody {
	return s.Body
}

func (s *DescribeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneClusterResponse) SetStatusCode(v int32) *DescribeMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneClusterResponse) SetBody(v *DescribeMultiZoneClusterResponseBody) *DescribeMultiZoneClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiZoneClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
