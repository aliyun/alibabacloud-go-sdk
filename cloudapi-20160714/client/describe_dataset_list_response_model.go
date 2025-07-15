// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatasetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatasetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatasetListResponseBody) *DescribeDatasetListResponse
	GetBody() *DescribeDatasetListResponseBody
}

type DescribeDatasetListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatasetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatasetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatasetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatasetListResponse) GetBody() *DescribeDatasetListResponseBody {
	return s.Body
}

func (s *DescribeDatasetListResponse) SetHeaders(v map[string]*string) *DescribeDatasetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatasetListResponse) SetStatusCode(v int32) *DescribeDatasetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatasetListResponse) SetBody(v *DescribeDatasetListResponseBody) *DescribeDatasetListResponse {
	s.Body = v
	return s
}

func (s *DescribeDatasetListResponse) Validate() error {
	return dara.Validate(s)
}
