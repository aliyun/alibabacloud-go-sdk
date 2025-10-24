// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudClustersResponseBody) *DescribeHybridCloudClustersResponse
	GetBody() *DescribeHybridCloudClustersResponseBody
}

type DescribeHybridCloudClustersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudClustersResponse) GetBody() *DescribeHybridCloudClustersResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudClustersResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudClustersResponse) SetStatusCode(v int32) *DescribeHybridCloudClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudClustersResponse) SetBody(v *DescribeHybridCloudClustersResponseBody) *DescribeHybridCloudClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
