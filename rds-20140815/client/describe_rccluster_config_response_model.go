// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCClusterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCClusterConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCClusterConfigResponseBody) *DescribeRCClusterConfigResponse
	GetBody() *DescribeRCClusterConfigResponseBody
}

type DescribeRCClusterConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCClusterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCClusterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCClusterConfigResponse) GetBody() *DescribeRCClusterConfigResponseBody {
	return s.Body
}

func (s *DescribeRCClusterConfigResponse) SetHeaders(v map[string]*string) *DescribeRCClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCClusterConfigResponse) SetStatusCode(v int32) *DescribeRCClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCClusterConfigResponse) SetBody(v *DescribeRCClusterConfigResponseBody) *DescribeRCClusterConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeRCClusterConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
