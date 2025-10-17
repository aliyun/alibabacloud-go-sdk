// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUsedUtilizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterUsedUtilizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterUsedUtilizationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterUsedUtilizationResponseBody) *DescribeClusterUsedUtilizationResponse
	GetBody() *DescribeClusterUsedUtilizationResponseBody
}

type DescribeClusterUsedUtilizationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterUsedUtilizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterUsedUtilizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUsedUtilizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterUsedUtilizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterUsedUtilizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterUsedUtilizationResponse) GetBody() *DescribeClusterUsedUtilizationResponseBody {
	return s.Body
}

func (s *DescribeClusterUsedUtilizationResponse) SetHeaders(v map[string]*string) *DescribeClusterUsedUtilizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterUsedUtilizationResponse) SetStatusCode(v int32) *DescribeClusterUsedUtilizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterUsedUtilizationResponse) SetBody(v *DescribeClusterUsedUtilizationResponseBody) *DescribeClusterUsedUtilizationResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterUsedUtilizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
