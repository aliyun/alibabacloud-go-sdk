// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatasetItemListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatasetItemListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatasetItemListResponseBody) *DescribeDatasetItemListResponse
	GetBody() *DescribeDatasetItemListResponseBody
}

type DescribeDatasetItemListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatasetItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatasetItemListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatasetItemListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatasetItemListResponse) GetBody() *DescribeDatasetItemListResponseBody {
	return s.Body
}

func (s *DescribeDatasetItemListResponse) SetHeaders(v map[string]*string) *DescribeDatasetItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatasetItemListResponse) SetStatusCode(v int32) *DescribeDatasetItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatasetItemListResponse) SetBody(v *DescribeDatasetItemListResponseBody) *DescribeDatasetItemListResponse {
	s.Body = v
	return s
}

func (s *DescribeDatasetItemListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
