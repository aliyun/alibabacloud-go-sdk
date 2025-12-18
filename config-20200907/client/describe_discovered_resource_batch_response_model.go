// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiscoveredResourceBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiscoveredResourceBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiscoveredResourceBatchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiscoveredResourceBatchResponseBody) *DescribeDiscoveredResourceBatchResponse
	GetBody() *DescribeDiscoveredResourceBatchResponseBody
}

type DescribeDiscoveredResourceBatchResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiscoveredResourceBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiscoveredResourceBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiscoveredResourceBatchResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiscoveredResourceBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiscoveredResourceBatchResponse) GetBody() *DescribeDiscoveredResourceBatchResponseBody {
	return s.Body
}

func (s *DescribeDiscoveredResourceBatchResponse) SetHeaders(v map[string]*string) *DescribeDiscoveredResourceBatchResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiscoveredResourceBatchResponse) SetStatusCode(v int32) *DescribeDiscoveredResourceBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiscoveredResourceBatchResponse) SetBody(v *DescribeDiscoveredResourceBatchResponseBody) *DescribeDiscoveredResourceBatchResponse {
	s.Body = v
	return s
}

func (s *DescribeDiscoveredResourceBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
