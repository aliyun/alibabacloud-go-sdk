// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataObjectsResponseBody) *DescribeDataObjectsResponse
	GetBody() *DescribeDataObjectsResponseBody
}

type DescribeDataObjectsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataObjectsResponse) GetBody() *DescribeDataObjectsResponseBody {
	return s.Body
}

func (s *DescribeDataObjectsResponse) SetHeaders(v map[string]*string) *DescribeDataObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectsResponse) SetStatusCode(v int32) *DescribeDataObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectsResponse) SetBody(v *DescribeDataObjectsResponseBody) *DescribeDataObjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataObjectsResponse) Validate() error {
	return dara.Validate(s)
}
