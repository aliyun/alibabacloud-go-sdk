// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCClustersResponseBody) *DescribeRCClustersResponse
	GetBody() *DescribeRCClustersResponseBody
}

type DescribeRCClustersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCClustersResponse) GetBody() *DescribeRCClustersResponseBody {
	return s.Body
}

func (s *DescribeRCClustersResponse) SetHeaders(v map[string]*string) *DescribeRCClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCClustersResponse) SetStatusCode(v int32) *DescribeRCClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCClustersResponse) SetBody(v *DescribeRCClustersResponseBody) *DescribeRCClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeRCClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
