// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOversizeNonPartitionTableInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOversizeNonPartitionTableInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOversizeNonPartitionTableInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOversizeNonPartitionTableInfosResponseBody) *DescribeOversizeNonPartitionTableInfosResponse
	GetBody() *DescribeOversizeNonPartitionTableInfosResponseBody
}

type DescribeOversizeNonPartitionTableInfosResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOversizeNonPartitionTableInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOversizeNonPartitionTableInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOversizeNonPartitionTableInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) GetBody() *DescribeOversizeNonPartitionTableInfosResponseBody {
	return s.Body
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) SetHeaders(v map[string]*string) *DescribeOversizeNonPartitionTableInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) SetStatusCode(v int32) *DescribeOversizeNonPartitionTableInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) SetBody(v *DescribeOversizeNonPartitionTableInfosResponseBody) *DescribeOversizeNonPartitionTableInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
