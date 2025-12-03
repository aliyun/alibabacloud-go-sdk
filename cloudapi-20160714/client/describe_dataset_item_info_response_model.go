// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatasetItemInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatasetItemInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatasetItemInfoResponseBody) *DescribeDatasetItemInfoResponse
	GetBody() *DescribeDatasetItemInfoResponseBody
}

type DescribeDatasetItemInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatasetItemInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatasetItemInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatasetItemInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatasetItemInfoResponse) GetBody() *DescribeDatasetItemInfoResponseBody {
	return s.Body
}

func (s *DescribeDatasetItemInfoResponse) SetHeaders(v map[string]*string) *DescribeDatasetItemInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatasetItemInfoResponse) SetStatusCode(v int32) *DescribeDatasetItemInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatasetItemInfoResponse) SetBody(v *DescribeDatasetItemInfoResponseBody) *DescribeDatasetItemInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDatasetItemInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
